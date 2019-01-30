package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/luojilab/json2graphqlschema/inspect"
)

func Run(port string) {
	router := chi.NewRouter()
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
