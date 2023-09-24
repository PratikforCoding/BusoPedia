package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	controller "github.com/PratikforCoding/BusoPedia.git/controllers"
	"github.com/PratikforCoding/BusoPedia.git/database"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongouri := os.Getenv("connectlink")

	busCol, usrCol, err := database.CreateDB(mongouri)
	if err != nil {
		log.Fatal("Didn't create connection to mongodb")
	}
	defer database.CloseDB()

	apicfg := controller.NewAPIConfig(busCol, usrCol)

	fmt.Println("MongoDB API")
	router := chi.NewRouter()
	apiRouter := chi.NewRouter()
	userRouter := chi.NewRouter()

	fsHandler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	router.Handle("/app/", fsHandler)
	router.Handle("/app/*", fsHandler)

	apiRouter.Get("/getbuses", apicfg.HandlerGetBuses)
	apiRouter.Get("/getbusbyname", apicfg.HandlerGetBusByName)
	apiRouter.Get("/getstopages", apicfg.HandlerGetBusStopagesByName)
	apiRouter.Post("/addbus", apicfg.HandlerAddBuses)

	userRouter.Post("/createaccount", apicfg.HandlerCreateAccount)
	userRouter.Post("/login", apicfg.HandlerLogin)

	router.Mount("/api", apiRouter)
	router.Mount("/usr", userRouter)

	corsMux := middlewareCors(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: corsMux,
	}

	log.Println("Server is getting started at port: 8080 ....")
	log.Fatal(server.ListenAndServe())
	log.Println("Server is runnig at port: 8080 ....")
}
