package handlers

import (
	"a2sv_sample/Golang_Library/MicroServices/models"
	"log"
	"net/http"
)




type Products struct{
	l *log.Logger
}

func NewProducts(logger *log.Logger) *Products{
	return &Products{logger}

}



func (products *Products) ServeHTTP(w http.ResponseWriter, r *http.Request){
   if r.Method == http.MethodGet{
	products.getProducts(w, r)
	return

   }


   w.WriteHeader(http.StatusMethodNotAllowed)

}



func (products *Products) getProducts(w http.ResponseWriter, _ *http.Request){
	productList := models.GetProducts()

	err := productList.ToJSON(w)
    if err != nil{
		http.Error(w, "Error In Fetching ALl Product List", http.StatusInternalServerError)
		return
	}

}