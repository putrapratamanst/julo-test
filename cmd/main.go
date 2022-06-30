package main

import (
	"context"
	"fmt"
	"julo-test/infrastructure"
	"julo-test/pkg"
	"julo-test/presenter"
	repository "julo-test/repository/account"
	"julo-test/router"
	service "julo-test/service/account"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	godotenv.Load()
	ctx := context.Background()
	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT_PRODUCT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println("Listening Product API ...", serverString)

	r := gin.Default()
	client, err := infrastructure.CreateClient(ctx)
	if err != nil {
		panic("Cannot create client firestore: " + err.Error())
	}

	defer client.Close()

	repoAccount := repository.NewRepository(client)
	serviceAccount := service.NewService(repoAccount)

	v1 := r.Group("/api/v1")
	{
		router.RouteAccount(v1, serviceAccount)
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
