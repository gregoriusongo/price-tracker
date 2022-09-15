package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gregoriusongo/price-tracker/pkg/util"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

var (
	dbpool *pgxpool.Pool
)

func init() {
	Connect()
}

func Connect() {
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connect to db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.DB.Host, config.DB.User, config.DB.Pass, config.DB.Name, config.DB.Port)

	dbpoolConn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("connected to database")
	}

	dbpool = dbpoolConn
}

func GetDB() *pgxpool.Pool {
	return dbpool
}