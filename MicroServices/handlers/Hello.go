package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//  Create Some Struct Here That Used To Imlplement The Interface of Serve Http Like Meaning Ou



type Hello struct{
	l *log.Logger
}


func NewHello(logger *log.Logger) *Hello{
	return &Hello{logger}

}




//  We Pass Our Hello Struct To ServeHttp That Mean Out Hello Struct Is Able To Implement The 
func (logger *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request){
	logger.l.Println("Hello World From My Custome Logger")

	data, err := io.ReadAll((r.Body))

	if err != nil{
		http.Error(w, "There Is Error In The Hello Handler", http.StatusInternalServerError)
		return

	}
	fmt.Fprintln(w, "This Is Response Message With Data Microservices : ", string(data))

}