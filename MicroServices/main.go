package main

import (
	"a2sv_sample/Golang_Library/MicroServices/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){
	logger := log.New(os.Stdout, "MICROSERVICES_PRODUCT_API : ", log.LstdFlags )


	serverMux := http.NewServeMux()
	FetchProductList := handlers.NewProducts(logger)
	serverMux.Handle("/products", FetchProductList)






	server := &http.Server{
		Addr: ":8080",
		Handler: serverMux,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout: time.Second * 100,
		
	}
	go func ()  {

	err := server.ListenAndServe()
	if err != nil{
		logger.Fatal(err)
	}
		
	}()


	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	sig := <- signalChan
	logger.Printf("Received termination signal: %v. Initiating graceful shutdown...\n", sig)



   tc, cancel:= context.WithTimeout(context.Background(), time.Second * 40)
   defer cancel()
	
   if err := server.Shutdown(tc); err != nil {
    logger.Printf("Error during server shutdown: %v\n", err)
}
}