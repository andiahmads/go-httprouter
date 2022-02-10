package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/* inti dari library httprouter adalah StructRouter
router ini merupakan implementasi dari http.Handler, sehingga kita bisa dengan mudah menambahkan kedalam http.server.
untuk membuat router, kita bisa menggunakan function httprouter.New() yg akan mengembalikan Router pointer.

berbeda dengan servemux, pada httprouter kita menggunakan type httprouter.Handle.
*/

func main() {
	router := httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "hello http router")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
