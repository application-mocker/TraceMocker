package handler

import (
	"TraceMocker/internal/web/model"
	"TraceMocker/utils"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
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

func PingHandler(ctx *gin.Context) {
	utils.Logger.Info("Ping -> Pong")
	w := ctx.Writer
	_, err := w.WriteString("pong")
	if err != nil {
		ErrorCtx(ctx, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// SimpleHandler only GET target next-route, and return target response.
func SimpleHandler(ctx *gin.Context) {
	reqBody := &model.SimpleRequestBody{}
	if err := ctx.BindJSON(reqBody); err != nil {
		ErrorCtx(ctx, http.StatusBadRequest, err)
		return
	}

	var nextBody io.Reader
	if reqBody.NextBody != nil {
		nextBodyBytes, err := json.Marshal(reqBody.NextBody)
		if err != nil {
			ErrorCtx(ctx, http.StatusBadRequest, err)
			return
		}
		nextBody = bytes.NewReader(nextBodyBytes)
	}

	request, err := http.NewRequest(http.MethodGet, reqBody.NextRoute, nextBody)
	if err != nil {
		ErrorCtx(ctx, http.StatusBadRequest, err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		ErrorCtx(ctx, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ErrorCtx(ctx, http.StatusBadRequest, err)
		return
	}

	_, _ = ctx.Writer.Write(body)
	ctx.Writer.WriteHeader(http.StatusOK)
}
