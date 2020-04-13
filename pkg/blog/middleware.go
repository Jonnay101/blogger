package blog

import "net/http"

func adminOnly(originalHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !testUserIsAdmin(true) {
			http.NotFound(w, r)
			return
		}
		originalHandler(w, r)
	}
}

func testUserIsAdmin(b bool) bool {
	user := &User{admin: b}
	return user.IsAdmin()
}
