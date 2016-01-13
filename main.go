package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func main() {
	var store = sessions.NewCookieStore([]byte("secret"))
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		fmt.Println("%#v", session)
	})

	http.ListenAndServe(":4321", nil)
}
