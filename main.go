package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := httprouter.New()
	closure := NewClosureController(GetSession())
	router.GET("/", Index)
	router.GET("/clinics", AllClinics)
	router.GET("/clinics/:id", FindByIDClinic)
	router.GET("/closures/:id", closure.GetClosure)
	router.POST("/closures", closure.CreateClosure)
	router.DELETE("/closures/:id", closure.RemoveClosure)
	router.GET("/closures", closure.SearchClosures)
	router.PUT("/closures/:id", closure.UpdateClosure)
	router.GET("/kinds/receipt", KindsReceipt)
	router.GET("/kinds/costs", KindsCosts)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
