package controller

import "net/http"

type userController struct{}

func (uc userContoller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{"Hello from the User Controller"})
}
