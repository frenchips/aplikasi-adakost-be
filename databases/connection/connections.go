package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

var (
	DBConnections *sql.DB
	err           error
)

func DbConnection() {

	err = godotenv.Load("config/db.env")
	if err != nil {
		panic("Error loading .env file")
	}

	viper.Set("db_host", os.Getenv("DB_HOST"))
	viper.Set("db_port", os.Getenv("DB_PORT"))
	viper.Set("db_user", os.Getenv("DB_USER"))
	viper.Set("db_password", os.Getenv("DB_PASSWORD"))
	viper.Set("db_name", os.Getenv("DB_NAME"))

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("db_host"),
		viper.GetInt("db_port"),
		viper.GetString("db_user"),
		viper.GetString("db_password"),
		viper.GetString("db_name"),
	)

	DBConnections, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	// check connection
	err = DBConnections.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
}
