package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAdminListUsers(ctx *gin.Context) {
	var perms []UserPermission
	cigarDB.Find(&perms)
	ctx.JSON(http.StatusOK, perms)
}

func HandleAdminUpdateUser(ctx *gin.Context) {
	var perm UserPermission
	if err := ctx.ShouldBindJSON(&perm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cigarDB.Save(&perm)
	ctx.JSON(http.StatusOK, perm)
}

func HandleDeleteCigar(ctx *gin.Context) {
	brand := ctx.Query("brand")
	name := ctx.Query("name")
	if brand == "" || name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "brand and name required"})
		return
	}
	result := cigarDB.Delete(&Cigar{}, "\"Brand\" = ? AND \"Name\" = ?", brand, name)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"rows_affected": result.RowsAffected})
}
