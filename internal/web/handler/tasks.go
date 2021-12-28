package handler

import (
	"TraceMocker/internal/task"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListTask(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, task.ProcessorInstance.ListAllTask())
}

func RegisterTask(ctx *gin.Context) {
	taskObj := &task.Info{}
	err := ctx.BindJSON(taskObj)
	if err != nil {
		ErrorCtx(ctx, http.StatusBadRequest, err)
		return
	}

	res := task.CreateTask(*taskObj)
	if res != nil {
		ErrorCtx(ctx, http.StatusBadRequest, res)
		return
	}

	return
}