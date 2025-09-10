package router

import (
	"net/http"
	"fmt"
)

var APIHandler *http.ServeMux = http.NewServeMux()

func init(){
	APIHandler.HandleFunc("/v1", func (w http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Method)
		if req.Method == http.MethodGet {
			fmt.Fprintf(w, "You reached v1 of API")
		}
	})
}