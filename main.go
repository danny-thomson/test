package main

import "net/http"

func main() {
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8090", nil)
}

func home(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("testing jenkins"))
}
