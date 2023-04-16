package main

import (
	"github.com/Lauri92/go-cats-api/pkg/cats"
	"github.com/Lauri92/go-cats-api/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	cats.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}
