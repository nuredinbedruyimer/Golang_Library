package handlers

import (
	"log"
	"net/http"
)




type GoodBye struct{
	l *log.Logger 
}

func NewGoodBye(logger *log.Logger) *GoodBye{
	return &GoodBye{logger}
}

func (g *GoodBye) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Good Bye In MicroServices"))
}