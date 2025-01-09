package main

import (
	"fmt"
	"io"
	"net/http"
)

const webPort = "80"

func getIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from Get request")
}

func main() {
	http.HandleFunc("/", getIndex)

	err := http.ListenAndServe(fmt.Sprintf(":%s", webPort), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on port:", webPort)
}
