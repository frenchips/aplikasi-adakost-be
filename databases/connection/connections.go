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

	_ = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	viper.Set("db_host", os.Getenv("PGHOST"))
	viper.Set("db_port", os.Getenv("PGPORT"))
	viper.Set("db_user", os.Getenv("PGUSER"))
	viper.Set("db_password", os.Getenv("PGPASSWORD"))
	viper.Set("db_name", os.Getenv("PGDATABASE"))

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
