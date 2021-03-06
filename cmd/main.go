package main

import (
	"fmt"
	"julo-test/infrastructure"
	"julo-test/pkg"
	"julo-test/presenter"
	repositoryAccount "julo-test/repository/account"
	repositoryWallet "julo-test/repository/wallet"
	"julo-test/router"
	serviceAccount "julo-test/service/account"
	serviceWallet "julo-test/service/wallet"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	godotenv.Load()
	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT_PRODUCT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println("Listening MINI WALLET API ...", serverString)

	r := gin.Default()
	// client, err := infrastructure.CreateClient(ctx)
	// if err != nil {
	// 	panic("Cannot create client firestore: " + err.Error())
	// }
	// defer client.Close()

	redisConfig := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	redisCache, errRedis := infrastructure.NewDatabase(redisConfig, os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_DB"))
	if errRedis != nil {
		panic(errRedis.Error())
	}

	repoAccount := repositoryAccount.NewRepository(redisCache)
	serviceAccount := serviceAccount.NewService(repoAccount)

	repoWallet := repositoryWallet.NewRepository(redisCache)
	serviceWallet := serviceWallet.NewService(repoWallet)

	v1 := r.Group("/api/v1")
	{
		router.RouteAccount(v1, serviceAccount)
		router.RouteWallet(v1, serviceWallet, serviceAccount)
	}

	errorHandler(r)
	r.Run(serverString)

}

//handle error method and not found endpoint
func errorHandler(r *gin.Engine) {
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, presenter.Response{
			Status:  pkg.HTTP_STATUS_ERROR,
			Message: pkg.ErrMethodNotAllow.Error(),
		})
		c.Abort()
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, presenter.Response{
			Code:    http.StatusNotFound,
			Message: pkg.ErrInvalidURL.Error(),
		})
		c.Abort()
	})
}
