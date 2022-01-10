package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SingleResponseSizeHandler return specify size of response.
func SingleResponseSizeHandler(ctx *gin.Context){
	size := ctx.Query("size")

	sizeInt, err := strconv.ParseInt(size, 10, 0)
	if err != nil{
		ErrorCtx(ctx, http.StatusBadRequest, err)
		return
	}

	bf := bytes.Buffer{}
	for sizeInt > 0 {
		bf.WriteByte('A')
	}

	w := ctx.Writer
	_, _ = w.Write(bf.Bytes())
	w.WriteHeader(http.StatusOK)
}

// PingHandler return pong of response body
func PingHandler(ctx *gin.Context) {
	w := ctx.Writer
	_, err := w.WriteString("pong")
	if err != nil {
		ErrorCtx(ctx, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
