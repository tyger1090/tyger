package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("Website\\Tyger_site"))
	http.Handle("/Tyger_site/", http.StripPrefix("/Tyger_site/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Tyger_site/tyger.html")
	})

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}