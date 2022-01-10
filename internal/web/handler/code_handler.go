package handler

import (
	"TraceMocker/internal/rander"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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

var randomCodeHandlerRandomInstance = rander.NewAdvanceRander(time.Now().Unix())

func RandomCode(ctx *gin.Context) {
	codeMap := ctx.QueryMap("code")

	codeMapInt := map[interface{}]int{}

	for key, value := range codeMap {
		var err error
		var code, weight int64
		if weight, err = strconv.ParseInt(value, 10, 0); err != nil {
			continue
		}
		if code, err = strconv.ParseInt(key, 10, 0); err != nil {
			continue
		}
		codeMapInt[int(code)] = int(weight)
	}

	code := randomCodeHandlerRandomInstance.RandInterface(codeMapInt)
	if code == nil {
		ctx.JSON(http.StatusBadRequest, code)
	}
	if value, ok := code.(int); ok {
		ctx.JSON(value, codeMap)
	} else {
		ctx.JSON(http.StatusBadRequest, codeMap)
	}
}
