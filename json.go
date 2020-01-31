package main
import (
	"net/http"
	"encoding/json"
)

type Curso struct {
	Title string
	NumeroVideos int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request) {
		curso := Curso{"Curso de go", 30}
		json.NewEncoder(w).Encode(curso)
	})
	http.ListenAndServe(":8000", nil)
}