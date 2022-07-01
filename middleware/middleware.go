package middleware

import (
	"julo-test/pkg"
	"julo-test/presenter"
	"julo-test/service/account"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type authHeader struct {
	Token string `header:"Authorization"`
}

func AuthUser(ias *account.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		// bind Authorization Header to h and check for validation errors
		if err := c.ShouldBindHeader(&h); err != nil {
			if _, ok := err.(validator.ValidationErrors); ok {
				// we used this type in bind_data to extract desired fields from errs
				// you might consider extracting it
				result := presenter.Response{
					Status: pkg.HTTP_STATUS_FAIL,
					Data: presenter.ErrorResponseMessage{
						Error: pkg.ErrHeaderInvalid.Error(),
					},
				}
				c.JSON(http.StatusBadRequest, result)
				c.Abort()
				return
			}

			result := presenter.Response{
				Status:  pkg.HTTP_STATUS_ERROR,
				Message: pkg.ErrHeaderInvalid.Error(),
			}
			c.JSON(http.StatusInternalServerError, result)
			c.Abort()
			return
		}

		idTokenHeader := strings.Split(h.Token, "Token ")
		if len(idTokenHeader) < 2 {
			result := presenter.Response{
				Status:  pkg.HTTP_STATUS_ERROR,
				Message: pkg.ErrAuthorizationBearer.Error(),
			}
			c.JSON(http.StatusInternalServerError, result)
			c.Abort()
			return
		}

		//check user
		token := idTokenHeader[1]
		checkUser, errCheckUser := ias.GetUserService(token)
		if errCheckUser != nil {
			result := presenter.Response{
				Status: pkg.HTTP_STATUS_FAIL,
				Data: presenter.ErrorResponseMessage{
					Error: pkg.ErrForbiddenAccess.Error(),
				},
			}
			c.JSON(http.StatusInternalServerError, result)
			c.Abort()
			return
		}

		if checkUser.CustomerXID == "" {
			result := presenter.Response{
				Status: pkg.HTTP_STATUS_FAIL,
				Data: presenter.ErrorResponseMessage{
					Error: pkg.ErrForbiddenAccess.Error(),
				},
			}
			c.JSON(http.StatusInternalServerError, result)
			c.Abort()
			return
		}
		c.Set("customer_xid", checkUser.CustomerXID)
		c.Next()

	}
}
