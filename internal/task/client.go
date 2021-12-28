package task

import (
	"TraceMocker/config"
	"TraceMocker/utils"
	"encoding/json"
	oc "github.com/application-mocker/object-client"
)

var Client *oc.ObjectClient

var taskClient *oc.ObjectClient

func Init() {
	if config.Config.Application.ObjectClientConfig.Enable {
		var err error
		utils.Logger.Info("Init object client")
		Client, err = oc.NewObjectClient(config.Config.Application.ObjectClientConfig.Host, "__trace_mocker__")
		if err != nil {
			panic(err)
		}
		taskClient, err = Client.SubClient("task")
		if err != nil {
			panic(err)
		}
	}
}

func CreateTask(taskInfo Info) error {
	if taskClient == nil {
		return nil
	}

	_, err := taskClient.InsertOne(taskInfo)
	return err
}

func ListTask() ([]*Info, error) {
	res := make([]*Info, 0)
	if taskClient == nil {
		return res, nil
	}

	objs, err := taskClient.ListAllValue()
	if err != nil {
		return res, err
	}

	if objs == nil {
		return res, err
	}

	res = make([]*Info, len(objs))

	for index, item := range objs {
		jsonObj, err := json.Marshal(item.DataValue)
		if err != nil {
			return make([]*Info, 0), err
		}

		resItem := &Info{}

		if json.Unmarshal(jsonObj, resItem) != nil {
			return make([]*Info, 0), err
		}
		res[index] = resItem
	}

	return res, nil
}
