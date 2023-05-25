package handlers

import (
	"log"
	"net/http"

	"github.com/ArshpreetS/Golang_microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// lp => list of products
	// lp := data.GetProducts()

	// how to return it to users
	// the way we do that is by converting lp to JSON and we can do that using package encoding/json

	// Method 1: Marshalling!
	// d, err := json.Marshal(lp)
	// if err != nil {
	// 	http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	// }
	// rw.Write(d)

	// Method 2: Using Encoder (it saves buffer memory since we are not creating
	//  		 the object in our server)
	// err := lp.ToJSON(rw)
	// if err != nil {
	// 	http.Error(rw, "Unable to encode Json", http.StatusInternalServerError)
	// }

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode Json", http.StatusInternalServerError)
	}
}
