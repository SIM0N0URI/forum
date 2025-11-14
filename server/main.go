package main

import "net/http"

func main() {
	http.HandleFunc("/", functions.loginSide)
	http.HandleFunc("/", functions.registerSide)
	http.HandleFunc("/", functions.logoutSide)
	http.HandleFunc("/", functions.createPost)
	http.HandleFunc("/", functions.commentPost)
	http.HandleFunc("/", functions.reactPost)
}
