package handler

import (
	"TraceMocker/internal/web/model"
	"TraceMocker/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

// SimpleTransHandler only GET target next-route, and return target response.
func SimpleTransHandler(ctx *gin.Context) {
	reqBody := &model.SimpleRequestBody{}
	if err := ctx.BindJSON(reqBody); err != nil {
		ErrorCtx(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := utils.DoBlockRequestWithJson(http.MethodGet, reqBody.NextRoute, nil, nil, reqBody.NextBody)
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

// TraceHandler will invoke specify service with route: /mock/http/trace-mock
func TraceHandler(ctx *gin.Context) {
	traceRequestBody := make([]model.TraceRequestBody, 0, 0)
	if err := ctx.BindJSON(&traceRequestBody); err != nil {
		ErrorCtx(ctx, http.StatusBadRequest, err)
	}
	if len(traceRequestBody) > 0 {
		nextRouteInfo := traceRequestBody[0]
		newTraceRequestBody := traceRequestBody[1:]

		if len(nextRouteInfo.NextMethod) == 0 {
			nextRouteInfo.NextMethod = http.MethodPost
		}
		nextRouteInfo.NextMethod = strings.ToUpper(nextRouteInfo.NextMethod)

		resp, err := utils.DoBlockRequestWithJson(
			nextRouteInfo.NextMethod,
			fmt.Sprintf("http://%s/mock/http/trace-mock", nextRouteInfo.NextService),
			nil,
			nil,
			newTraceRequestBody)
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
	} else {
		ctx.JSON(http.StatusOK, nil)
	}

}
