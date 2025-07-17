package main

import (
	"api"
	"api/pkg/handlers"
	"api/pkg/repository"
	"api/pkg/service"
	"api/utils/logger"
	"fmt"
	"os"
)

func main() {
	/*if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}*/
	fmt.Println("Initializing...")
	fmt.Println("\n██╗   ██╗███████╗██████╗ ██████╗ ██╗███████╗██╗   ██╗\n██║   ██║██╔════╝██╔══██╗██╔══██╗██║██╔════╝╚██╗ ██╔╝\n██║   ██║█████╗  ██████╔╝██████╔╝██║█████╗   ╚████╔╝ \n╚██╗ ██╔╝██╔══╝  ██╔══██╗██╔══██╗██║██╔══╝    ╚██╔╝  \n ╚████╔╝ ███████╗██║  ██║██████╔╝██║██║        ██║   \n  ╚═══╝  ╚══════╝╚═╝  ╚═╝╚═════╝ ╚═╝╚═╝        ╚═╝   \n                                                     \n")
	fmt.Println("Version: 1.0.0")
	db, err := repository.NewDB(repository.DB{
		Hostname: os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Dbname:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL"),
	})
	if err != nil {
		logger.Log.Fatalf("Error connecting to database: %v", err)
	}
	logger.Log.Println("Connecting to database")
	NewRepository := repository.NewRepository(db)
	NewService := service.NewService(NewRepository)
	NewHandler := handlers.NewHandler(NewService)
	server := new(api.Server)
	logger.Log.Println("Running server")
	if err = server.Start(os.Getenv("SERVER_PORT"), NewHandler.InitRoutes()); err != nil {
		logger.Log.Fatalf("Fatal Error: %v", err)
	}
}
