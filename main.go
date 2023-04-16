package main

import (
	"github.com/Lauri92/go-cats-api/pkg/cats"
	"github.com/Lauri92/go-cats-api/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	cats.RegisterCatRoutes(r, h)

	err := r.Run(port)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
