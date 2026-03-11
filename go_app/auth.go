package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var oauthCfg *oauth2.Config

func InitOAuth() {
	oauthCfg = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

func randomString(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func HandleLogin(ctx *gin.Context) {
	state, err := randomString(16)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Redirect(http.StatusFound, oauthCfg.AuthCodeURL(state, oauth2.AccessTypeOffline))
}

func HandleOAuthCallback(ctx *gin.Context) {
	token, err := oauthCfg.Exchange(ctx, ctx.Query("code"))
	if err != nil {
		log.Printf("OAuth token exchange error: %s", err)
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	resp, err := oauthCfg.Client(ctx, token).Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Failed to get user info: %s", err)
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var userInfo map[string]any
	json.Unmarshal(body, &userInfo)

	email, _ := userInfo["email"].(string)
	if email == "" {
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	sessionID, err := randomString(32)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	cigarDB.Create(&Session{ID: sessionID, Email: email, CreatedAt: time.Now()})
	ctx.SetCookie("cigarland_session", sessionID, 86400*30, "/", "", false, true)
	ctx.Redirect(http.StatusSeeOther, "/")
}

func HandleLogout(ctx *gin.Context) {
	if cookie, err := ctx.Cookie("cigarland_session"); err == nil {
		cigarDB.Delete(&Session{ID: cookie})
	}
	ctx.SetCookie("cigarland_session", "", -1, "/", "", false, true)
	ctx.Redirect(http.StatusSeeOther, "/login")
}

func GetCurrentUser(ctx *gin.Context) (string, bool) {
	sessionID, err := ctx.Cookie("cigarland_session")
	if err != nil {
		return "", false
	}
	var session Session
	if cigarDB.Where("id = ?", sessionID).First(&session).Error != nil {
		return "", false
	}
	return session.Email, true
}

func GetUserPermission(email string) *UserPermission {
	var perm UserPermission
	if cigarDB.Where("email = ?", email).First(&perm).Error != nil {
		return nil
	}
	return &perm
}

func HandleMe(ctx *gin.Context) {
	email, ok := GetCurrentUser(ctx)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return
	}
	perm := GetUserPermission(email)
	if perm == nil {
		perm = &UserPermission{Email: email}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"email":      email,
		"can_add":    perm.CanAdd,
		"can_delete": perm.CanDelete,
		"can_admin":  perm.CanAdmin,
	})
}

func WithPermission(check func(*UserPermission) bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, ok := GetCurrentUser(ctx)
		if !ok {
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Abort()
			return
		}
		perm := GetUserPermission(email)
		if perm == nil || !check(perm) {
			ctx.Status(http.StatusForbidden)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func WithAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := GetCurrentUser(ctx); !ok {
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
