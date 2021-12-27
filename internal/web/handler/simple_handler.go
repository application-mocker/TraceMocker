package handler

import (
	"TraceMocker/internal/web/model"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
)

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
