package handler

import (
	"TraceMocker/config"
	"TraceMocker/internal/task"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListTask(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, task.ProcessorInstance.ListAllTask())
}

// DeleteTask will remove specify task from object-mocker
func DeleteTask(ctx *gin.Context) {
	//taskObj := &task.Info{}
	//err := ctx.BindJSON(taskObj)

}

// RegisterTask will create new task. If task.Info.Holder is nil, set to current instance's holder.
func RegisterTask(ctx *gin.Context) {
	taskObj := &task.Info{}
	if err := ctx.BindJSON(taskObj); err != nil{
		ErrorCtx(ctx, http.StatusBadRequest, err)
		return
	}
	if taskObj.Holder == "" {
		// default do task by itself
		taskObj.Holder = config.Config.Application.NodeId
	}

	// todo: check same name hand holder task

	res := task.CreateTask(*taskObj)
	if res != nil {
		ErrorCtx(ctx, http.StatusBadRequest, res)
		return
	}

	return
}
