#!/bin/bash
# Usage: ./test_api.sh [session_cookie] [base_url]
# Session cookie is optional — without it only unauthenticated tests run.
# Get cookie from browser DevTools → Application → Cookies → cigarland_session

SESSION="${1:-}"
BASE="${2:-https://cigarland.reneau.me}"

PASS=0
FAIL=0

green='\033[0;32m'
red='\033[0;31m'
yellow='\033[1;33m'
cyan='\033[0;36m'
reset='\033[0m'

section() { echo -e "\n${cyan}=== $1 ===${reset}"; }
pass()    { echo -e "  ${green}PASS${reset} $1"; ((PASS++)); }
fail()    { echo -e "  ${red}FAIL${reset} $1 — expected $2, got $3"; ((FAIL++)); }

# curl returning just status code, no redirect following
status() {
  local method="$1" url="$2" cookie="$3" data="$4" ctype="$5"
  local args=(-s -o /dev/null -w "%{http_code}" -X "$method" "$url" --max-redirs 0)
  [[ -n "$cookie" ]]  && args+=(-H "Cookie: cigarland_session=$cookie")
  [[ -n "$ctype" ]]   && args+=(-H "Content-Type: $ctype")
  [[ -n "$data" ]]    && args+=(--data "$data")
  curl "${args[@]}"
}

body() {
  local method="$1" url="$2" cookie="$3" data="$4" ctype="$5"
  local args=(-s -X "$method" "$url" --max-redirs 0)
  [[ -n "$cookie" ]]  && args+=(-H "Cookie: cigarland_session=$cookie")
  [[ -n "$ctype" ]]   && args+=(-H "Content-Type: $ctype")
  [[ -n "$data" ]]    && args+=(--data "$data")
  curl "${args[@]}"
}

expect() {
  local label="$1" expected="$2" got="$3"
  if [[ "$got" == "$expected" ]]; then pass "$label"
  else fail "$label" "$expected" "$got"; fi
}

expect_not() {
  local label="$1" unexpected="$2" got="$3"
  if [[ "$got" != "$unexpected" ]]; then pass "$label"
  else fail "$label" "!= $unexpected" "$got"; fi
}

# ─────────────────────────────────────────────
section "UNAUTHENTICATED — public endpoints"
# ─────────────────────────────────────────────

s=$(status GET "$BASE/me")
expect "GET /me without session → 401" "401" "$s"

s=$(status GET "$BASE/login")
expect "GET /login → redirect to Google (302)" "302" "$s"

# ─────────────────────────────────────────────
section "UNAUTHENTICATED — protected endpoints (expect redirect 303)"
# ─────────────────────────────────────────────

s=$(status GET "$BASE/api/test")
expect "GET /api/test (no auth) → 303" "303" "$s"

s=$(status GET "$BASE/api/cigars")
expect "GET /api/cigars (no auth) → 303" "303" "$s"

s=$(status PUT "$BASE/api/cigars" "" '{"brand":"X","name":"Y"}' "application/json")
expect "PUT /api/cigars (no auth) → 303" "303" "$s"

s=$(status POST "$BASE/api/test" "" '{"cigar_list":[]}' "application/json")
expect "POST /api/test (no auth) → 303" "303" "$s"

s=$(status POST "$BASE/api/where" "" '{"table":"Cigars","filters":[]}' "application/json")
expect "POST /api/where (no auth) → 303" "303" "$s"

s=$(status DELETE "$BASE/api/cigars?brand=Fake&name=Cigar")
expect "DELETE /api/cigars (no auth) → 303" "303" "$s"

s=$(status GET "$BASE/api/admin/users")
expect "GET /api/admin/users (no auth) → 303" "303" "$s"

s=$(status POST "$BASE/api/admin/users" "" '{"email":"x@x.com"}' "application/json")
expect "POST /api/admin/users (no auth) → 303" "303" "$s"

# ─────────────────────────────────────────────
if [[ -z "$SESSION" ]]; then
  echo -e "\n${yellow}No session cookie provided — skipping authenticated tests.${reset}"
  echo -e "${yellow}Usage: $0 <session_cookie>${reset}"
else
# ─────────────────────────────────────────────

section "AUTHENTICATED — /me and permissions"
# ─────────────────────────────────────────────

me_body=$(body GET "$BASE/me" "$SESSION")
me_status=$(status GET "$BASE/me" "$SESSION")
expect "GET /me with session → 200" "200" "$me_status"

email=$(echo "$me_body" | grep -o '"email":"[^"]*"' | cut -d'"' -f4)
can_add=$(echo "$me_body" | grep -o '"can_add":[a-z]*' | cut -d: -f2)
can_edit=$(echo "$me_body" | grep -o '"can_edit":[a-z]*' | cut -d: -f2)
can_delete=$(echo "$me_body" | grep -o '"can_delete":[a-z]*' | cut -d: -f2)
can_admin=$(echo "$me_body" | grep -o '"can_admin":[a-z]*' | cut -d: -f2)

echo "    email=$email  can_add=$can_add  can_edit=$can_edit  can_delete=$can_delete  can_admin=$can_admin"

# ─────────────────────────────────────────────
section "AUTHENTICATED — general access"
# ─────────────────────────────────────────────

s=$(status GET "$BASE/api/test" "$SESSION")
expect "GET /api/test (authed) → 200" "200" "$s"

s=$(status GET "$BASE/api/cigars" "$SESSION")
expect "GET /api/cigars (authed) → 200" "200" "$s"

cigar_count=$(body GET "$BASE/api/cigars" "$SESSION" | grep -o '"brand"' | wc -l | tr -d ' ')
echo "    cigars in DB: $cigar_count"

# ─────────────────────────────────────────────
section "AUTHENTICATED — can_delete permission (DELETE /api/cigars)"
# ─────────────────────────────────────────────

s=$(status DELETE "$BASE/api/cigars?brand=Cohiba&name=Siglo%20VI" "$SESSION")
if [[ "$can_delete" == "true" ]]; then
  expect "DELETE Cohiba Siglo VI (can_delete=true) → 200" "200" "$s"
  cigars_after=$(body GET "$BASE/api/cigars" "$SESSION")
  if echo "$cigars_after" | grep -q '"Siglo VI"'; then
    fail "Cohiba Siglo VI removed from DB" "not found" "still present"
  else
    pass "Cohiba Siglo VI confirmed removed from DB"
  fi
else
  expect "DELETE Cohiba Siglo VI (can_delete=false) → 403" "403" "$s"
fi

# ─────────────────────────────────────────────
section "AUTHENTICATED — can_add permission (POST /api/test)"
# ─────────────────────────────────────────────

cohiba_payload=$(cat <<'EOF'
{"cigar_list":[{"brand":"Cohiba","name":"Siglo VI","wrapper":"Cuban","profile":"Medium-Full","binder":"Cuban","pressed":false,"tasty_tip":false,"spicy":4,"rating":9,"length":150,"ring":52,"review":"The flagship of Cuban cigars. Creaminess and complexity in perfect balance — cedar, honey, and subtle spice. Burns slow and even with a tight white ash.","kyle_rating":0,"kyle_review":"","john_rating":0,"john_review":"","image_ref":"","authentic_human_review":""}]}
EOF
)

s=$(status POST "$BASE/api/test" "$SESSION" "$cohiba_payload" "application/json")
if [[ "$can_add" == "true" ]]; then
  expect "POST /api/test re-adds Cohiba Siglo VI (can_add=true) → 200" "200" "$s"
  cigars_after=$(body GET "$BASE/api/cigars" "$SESSION")
  if echo "$cigars_after" | grep -q '"Siglo VI"'; then
    pass "Cohiba Siglo VI confirmed back in DB"
  else
    fail "Cohiba Siglo VI back in DB" "found" "not found"
  fi
else
  expect "POST /api/test (can_add=false) → 403" "403" "$s"
fi

# ─────────────────────────────────────────────
section "AUTHENTICATED — can_edit permission (PUT /api/cigars)"
# ─────────────────────────────────────────────

edit_payload=$(cat <<'EOF'
{"brand":"Cohiba","name":"Siglo VI","wrapper":"Maduro","profile":"Medium-Full","binder":"Cuban","pressed":false,"tasty_tip":false,"spicy":4,"rating":9,"length":150,"ring":52,"review":"Edited by test script.","kyle_rating":0,"kyle_review":"","john_rating":0,"john_review":"","image_ref":"","authentic_human_review":""}
EOF
)

s=$(status PUT "$BASE/api/cigars" "$SESSION" "$edit_payload" "application/json")
if [[ "$can_edit" == "true" ]]; then
  expect "PUT /api/cigars edits Cohiba Siglo VI (can_edit=true) → 200" "200" "$s"
  cigars_body=$(body GET "$BASE/api/cigars" "$SESSION")
  if echo "$cigars_body" | grep -q '"Edited by test script."'; then
    pass "Edit change confirmed in DB"
  else
    fail "Edit change confirmed in DB" "review=Edited by test script." "not found"
  fi
else
  expect "PUT /api/cigars (can_edit=false) → 403" "403" "$s"
fi

# ─────────────────────────────────────────────
section "AUTHENTICATED — can_admin permission (GET/POST /api/admin/users)"
# ─────────────────────────────────────────────

s=$(status GET "$BASE/api/admin/users" "$SESSION")
if [[ "$can_admin" == "true" ]]; then
  expect "GET /api/admin/users (can_admin=true) → 200" "200" "$s"
  user_count=$(body GET "$BASE/api/admin/users" "$SESSION" | grep -o '"email"' | wc -l | tr -d ' ')
  echo "    users in DB: $user_count"
else
  expect "GET /api/admin/users (can_admin=false) → 403" "403" "$s"
fi

s=$(status POST "$BASE/api/admin/users" "$SESSION" '{"email":"noone@example.com","can_add":false,"can_edit":false,"can_delete":false,"can_admin":false}' "application/json")
if [[ "$can_admin" == "true" ]]; then
  expect "POST /api/admin/users (can_admin=true) → 200" "200" "$s"
else
  expect "POST /api/admin/users (can_admin=false) → 403" "403" "$s"
fi

# ─────────────────────────────────────────────
section "AUTHENTICATED — wrong method / bad input"
# ─────────────────────────────────────────────

s=$(status DELETE "$BASE/api/cigars" "$SESSION")
if [[ "$can_delete" == "true" ]]; then
  expect "DELETE /api/cigars with no params → 400" "400" "$s"
else
  expect "DELETE /api/cigars (no perms, no params) → 403" "403" "$s"
fi

fi  # end SESSION block

# ─────────────────────────────────────────────
echo -e "\n${cyan}Results: ${green}$PASS passed${reset}, ${red}$FAIL failed${reset}"
[[ $FAIL -eq 0 ]] && exit 0 || exit 1
