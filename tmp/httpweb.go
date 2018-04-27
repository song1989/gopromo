package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["bookId"])
	for k, v := range r.Form {
		fmt.Printf("key: %s, value: %s \n", k, strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Respone message:Server get bookId successed....")
}

func main() {
	http.HandleFunc("/", HandleRequest)
	err := http.ListenAndServe(":6666", nil)
	if err != nil {
		log.Fatal("ListenError:", err)
	}
}
