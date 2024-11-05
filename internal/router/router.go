package router

import (
	// "errors"
	// "os"

	"github.com/gin-gonic/gin"
	//"github.com/vinicius457/BancoDeDados2/config"
)

var (
	port = "8080"
)

func Inicialize() {
	//logger := config.GetLogger(("router"))
	router := gin.Default()

	initRoutes(router)
	
	// port := os.Getenv("APP_PORT")
	// if port == "" {
	// 	errDotEnv :=  errors.New("Missing environment variables for app initialization")
    //     logger.Warnf("app initialization warning: %v", errDotEnv)
    //     logger.Info("Initializing default app connection...")
		//port = "8080"
	//}
	
	router.Run(":"+port)
}
