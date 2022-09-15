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
	// // load config
	// viper.AddConfigPath(".")
	// viper.SetConfigFile(`pkg/config/config.json`)

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(err)
	// }

	Connect()
}

func Connect() {
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// fmt.Println("host: " + config.DBHost)
	// dbHost := viper.GetString(`database.host`)
	// dbPort := viper.GetString(`database.port`)
	// dbUser := viper.GetString(`database.user`)
	// dbPass := viper.GetString(`database.pass`)
	// dbName := viper.GetString(`database.name`)

	// connect to db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.DB.Host, config.DB.User, config.DB.Pass, config.DB.Name, config.DB.Port)

	// dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dbpoolConn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer dbpoolConn.Close()

	if err != nil {
		log.Fatal(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("connected to database")
	}

	dbpool = dbpoolConn
}

func GetDB() *pgxpool.Pool {
	return dbpool
}

// func Connect() {
// 	dbHost := viper.GetString(`database.host`)
// 	dbPort := viper.GetString(`database.port`)
// 	dbUser := viper.GetString(`database.user`)
// 	dbPass := viper.GetString(`database.pass`)
// 	dbName := viper.GetString(`database.name`)

// 	// connect to db
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)

// 	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if viper.GetBool(`debug`) {
// 		fmt.Println("connected to database")
// 	}

// 	db = dbConn
// }

// func GetDB() *gorm.DB {
// 	return db
// }
