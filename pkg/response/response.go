package response

import (
	"julo-test/pkg"
	"julo-test/presenter"
	"sort"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, res *presenter.Response) {
	if res.Code == 0 && res.Message == "" {
		res.Code = 204
		res.Message = "No Content"
	}

	res.Status = validateDefaultSuccess(res.Code)
	ctx.JSON(res.Code, res)
}

func validateDefaultSuccess(value int) string {
	defaultSuccess := []int{200, 201, 202, 203, 204, 206}
	i := sort.Search(len(defaultSuccess), func(i int) bool { return value <= defaultSuccess[i] })
	if i < len(defaultSuccess) && defaultSuccess[i] == value {
		return pkg.HTTP_STATUS_SUCCESS
	} else {
		return pkg.HTTP_STATUS_ERROR
	}
}
