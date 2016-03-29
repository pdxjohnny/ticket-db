package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	// Set viper defaults
	viper.SetDefault("user", "user")
	viper.SetDefault("pass", "pass")
	viper.SetDefault("host", "tcp(localhost:3306)")
	viper.SetDefault("name", "tickets")

	// Load viper config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	// Load env
	viper.SetEnvPrefix("tickets_db")
	viper.AutomaticEnv()

	// Format the db connection string
	dbConnectString := fmt.Sprintf(
		"%s:%s@%s/%s",
		viper.GetString("user"),
		viper.GetString("pass"),
		viper.GetString("host"),
		viper.GetString("name"),
	)

	// Open the connection to the database
	db, err := sql.Open("mysql", dbConnectString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
