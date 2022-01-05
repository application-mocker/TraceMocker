package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SingleCodeHandler(ctx *gin.Context) {
	code := ctx.Query("code")

	codeInt, err := strconv.ParseInt(code, 10, 0)
	if err != nil {
		ErrorCtx(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(int(codeInt), map[string]string{"code": code})
}

func RandomCode(ctx *gin.Context) {

}
