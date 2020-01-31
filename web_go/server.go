package main
import (
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("public"))
/* 	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request) {
		http.ServeFile(w,r,r.URL.Path[1:])
	}) */
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))
	http.ListenAndServe(":8000", nil)

}