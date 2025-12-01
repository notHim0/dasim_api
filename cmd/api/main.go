package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/notHim0/dasim_api/pkg/database"
	"github.com/notHim0/dasim_api/pkg/supabase"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	if _,err:= database.ConnectDB(); err != nil {
		log.Fatalf("Could not connect to Database: %v", err)
	}

	defer database.Close()

	if err := supabase.Init(); err != nil {
		log.Fatalf("Could not initialise Supabase Client: :%v", err)
	}

	log.Print("Server is running...\n")

}