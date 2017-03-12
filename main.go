package main

import (
	"os"
	"syscall"

	"github.com/chyeh/viper"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"

	commonGin "github.com/Cepave/open-falcon-backend/common/gin"
	log "github.com/Cepave/open-falcon-backend/common/logruslog"
	commonOs "github.com/Cepave/open-falcon-backend/common/os"
	"github.com/Cepave/open-falcon-backend/common/vipercfg"
)

var logger = log.NewDefaultLogger("INFO")

func main() {
	/**
	 * Initialize loader of configurations
	 */
	confLoader := vipercfg.NewOwlConfigLoader()
	confLoader.FlagDefiner = pflagDefine

	confLoader.ProcessTrueValueCallbacks()
	// :~)

	config := confLoader.MustLoadConfigFile()

	InitAPIConfig(toAPIConfig(config))
	InitGin(toGinConfig(config))

	commonOs.HoldingAndWaitSignal(exitApp, syscall.SIGINT, syscall.SIGTERM)
}

func exitApp(signal os.Signal) {}

func toAPIConfig(config *viper.Viper) *APIAuthConfig {
	return &APIAuthConfig{
		googlePlaceAPIKey:     config.GetString("api_auth.google_place"),
		bookingDotComUsername: config.GetString("api_auth.booking_dot_com.username"),
		bookingDotComPassword: config.GetString("api_auth.booking_dot_com.password"),
	}
}

func toGinConfig(config *viper.Viper) *commonGin.GinConfig {
	return &commonGin.GinConfig{
		Mode: gin.ReleaseMode,
		Host: config.GetString("restful.listen.host"),
		Port: uint16(config.GetInt("restful.listen.port")),
	}
}

func pflagDefine() {
	pflag.StringP("config", "c", "cfg.json", "configuration file")
	pflag.BoolP("help", "h", false, "usage")
}
