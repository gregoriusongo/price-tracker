package main

import (
	"fmt"
	"log"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	// dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	// dbConn, err := sql.Open(`mysql`, dsn)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dbConn)
	// err = dbConn.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer func() {
	// 	err := dbConn.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// e := echo.New()
	// middL := _articleHttpDeliveryMiddleware.InitMiddleware()
	// e.Use(middL.CORS)
	// authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	// ar := _articleRepo.NewMysqlArticleRepository(dbConn)

	// timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	// au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	// _articleHttpDelivery.NewArticleHandler(e, au)

	// log.Fatal(e.Start(viper.GetString("server.address")))
}
