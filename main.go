package main
import (
	"golang.org/x/xerrors"
	"github.com/gin-gonic/gin"
	cont "github.com/say-ya-sigma/hello_task/backend/controller"
)

func main() {
	router := gin.Default()
	hello := cont.CreateHello()
	router.GET("/", func(ctx *gin.Context) {
		body, err := hello.Get(ctx)
		if err != nil {
			detectErr(ctx, err)
			return
		}
		ctx.JSON(200, body)
	})
	router.Run()
}

func detectErr(ctx *gin.Context, err error) {
	if xerrors.Is(err, cont.ErrHTTPNotFound) {
		ctx.JSON(404, gin.H{
			"message": cont.ErrHTTPNotFound,
		})
	} else if xerrors.Is(err, cont.ErrHTTPInternalServerError) {
		ctx.JSON(500, gin.H{
			"message": cont.ErrHTTPInternalServerError,
		})
	} else if xerrors.Is(err, cont.ErrHTTPUnauthorized) {
		ctx.JSON(401, gin.H{
			"message": cont.ErrHTTPUnauthorized,
		})
	}
}

