package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/syaifuddin.teddy/test-case-majoo/api/controllers"
	"gitlab.com/syaifuddin.teddy/test-case-majoo/api/seed"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	if os.Getenv("NoDB") == "Y" {
		server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	}

	if os.Getenv("SeedDb") == "Y" {
		seed.Load(server.DB) // enable this feature to activate auto-migrate DB
	}

	os.Mkdir("download", 0777) // create download folder

	server.Run(os.Getenv("APP_PORT"))

}
