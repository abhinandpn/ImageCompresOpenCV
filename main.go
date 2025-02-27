package main

import (
	"log"
	"net/http"

	"github.com/abhinandpn/ImageCompresOpenCV/handler"
	"github.com/gorilla/mux"
)

func main() {
	// INITIALIZE THE ROUTER USING MUX.NEWROUTER()
	router := mux.NewRouter()
	router.HandleFunc("/convert", handler.ConvertHandler).Methods("POST")
	router.HandleFunc("/resize", handler.ResizeHandler).Methods("POST")
	router.HandleFunc("/compress", handler.CompressHandler).Methods("POST")

	// SPECIFY THE SERVER ADDRESS
	serverAddr := ":8080"

	// PRINT A MESSAGE TO THE CONSOLE THAT THE SERVER IS STARTING AND LISTENING ON A SPECIFIC ADDRESS
	log.Printf("Server is starting and listening on %s...", serverAddr)

	// USE HTTP.LISTENANDSERVE() TO START THE SERVER WITH THE PREVIOUSLY DEFINED ROUTER
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
