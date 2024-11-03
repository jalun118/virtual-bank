package main

import (
	"os"
	"strings"
	"virtual-bank/internal/app"
	"virtual-bank/internal/helpers"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setUpEngine(log *helpers.Logger, viper *viper.Viper) *gin.Engine {
	gin.SetMode(viper.GetString("GIN_MODE"))

	log.Info("engine mode: [%s]", strings.ToUpper(gin.Mode()))

	r := gin.New()
	log.Info("setup router...")

	r.MaxMultipartMemory = 32 << 20
	log.Info("set max multipart memory...")

	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	return r
}

func main() {
	viper := app.NewViper()

	logger := helpers.NewLogger(os.Stdout)

	engine := setUpEngine(logger, viper)

	db := app.NewDatabase(viper)

	app.AppBootstrap(&app.BootstrapConfig{
		App: engine,
		DB:  db,
		Log: logger,
	})

	PORT := viper.GetString("PORT")
	engine.Run(":" + PORT)
}
