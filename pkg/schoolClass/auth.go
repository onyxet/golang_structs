package schoolClass

import "net/http"

type User struct {
	Username string
	Password string
}

var authorizedUsers = map[string]string{
	"teacher": "password",
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || authorizedUsers[username] != password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
