package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/luojilab/json2graphql/inspect"
)

func Run(port string) {
	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(cors.Handler)

	router.Post("/inspect/", Inspect)
	fmt.Printf("start listening on port is %s\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func Inspect(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	output, err := inspect.InspectWithBytes(body)
	fmt.Printf("body is\n %v\n", string(body))
	fmt.Printf("return is \n%v\n", string(output))
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(fmt.Sprintf("%v", err)))
	} else {
		w.Write(output)
	}
}
