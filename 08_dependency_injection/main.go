package main

import (
	"fmt"
	"lgwts/08_dependency_injection/greet"
	"net/http"
)

func MyGreetrHandler(w http.ResponseWriter, r *http.Request) {
	greet.Greet(w, "world")
}

func main() {
	err := http.ListenAndServe("localhost:5555",
		http.HandlerFunc(MyGreetrHandler))
	if err != nil {
		fmt.Println(err)
	}
}
