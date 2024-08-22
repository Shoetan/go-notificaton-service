package database

import (
	"fmt"
	"log"

	"github.com/Shoetan/utils"
	"github.com/jmoiron/sqlx"
)

func Database() ( *sqlx.DB, error)  {


	envVariables := utils.GetEnvVariables("POSTGRES_PASSWORD", "POSTGRES_USER", "POSTGRES_DB")

	db_password := envVariables["POSTGRES_PASSWORD"]
	db_user := envVariables["POSTGRES_USER"]
	db_name := envVariables["POSTGRES_DB"]
	
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", "localhost", db_user, db_password, db_name, "disable" )

	db, err := sqlx.Connect("postgres", connectionString )

	if err != nil {
		log.Fatal("Could not connect to the database", err.Error())
	}

	if err := db.Ping(); err !=nil {
		log.Fatal(err.Error())
	} else {
		log.Println("Connection to database successful   ")
	}

	return db, err
	
}