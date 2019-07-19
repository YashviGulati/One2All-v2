package basicauthmiddleware

import (
	"net/http"
)

func BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pw, flag := r.BasicAuth()
		if !flag || !checkCred(user, pw) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Enter your credentials"`)
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized.\n"))
			return
		}
		handler(w, r)
	}
}

func checkCred(username, password string) bool {
	return username == "admin" && password == "admin"
}

/*
Documentation

func BasicAuth():
	returns the username and password provided in the request's Authorization header

func Set():
	used to set keys and values. Also known as HTTP Trailers.
 */
