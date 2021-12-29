package task

import (
	"TraceMocker/config"
	"TraceMocker/utils"
	"bytes"
	"github.com/robfig/cron/v3"
	"net/http"
	"sync"
)

var ProcessorInstance Processor

func InitProcessor() {
	ProcessorInstance = Processor{
		ids:          map[string]cron.EntryID{},
		infos:        map[string]*Info{},
		cronInstance: cron.New(),
		lock:         &sync.RWMutex{},
	}
	ProcessorInstance.cronInstance.Start()
}

type Processor struct {
	cronInstance *cron.Cron
	ids          map[string]cron.EntryID
	infos        map[string]*Info

	lock *sync.RWMutex
}

func (p *Processor) RegisterTask(task *Info) bool {
	p.lock.Lock()
	defer p.lock.Unlock()
	if _, ok := p.ids[task.Name]; ok {
		return false
	}

	id, err := p.cronInstance.AddFunc(task.Cron, GeneratorTaskFunc(task))
	if err != nil {
		return false
	}

	p.ids[task.Name] = id
	p.infos[task.Name] = task
	return true
}

func (p *Processor) Remove(task string) bool {
	p.lock.Lock()
	defer p.lock.Unlock()
	if _, ok := p.ids[task]; !ok {
		return false
	} else {
		p.cronInstance.Remove(p.ids[task])
		delete(p.ids, task)
		delete(p.infos, task)
		return true
	}
}

func (p *Processor) Sync() {
	p.lock.Lock()
	LocalTasks := map[string]bool{}
	for name, _ := range p.infos {
		LocalTasks[name] = true
	}
	p.lock.Unlock()

	allTasks, err := ListTask()
	if err != nil {
		utils.Logger.Errorf("Sync task error: %v", allTasks)
	}

	// sync task in object-mocker
	for _, taskItem := range allTasks {
		if taskItem.Holder == config.NodeId {
			p.RegisterTask(taskItem)
			LocalTasks[taskItem.Name] = false
		}
	}

	// remove task which is not exits
	for taskName, exits := range LocalTasks {
		if exits {
			p.Remove(taskName)
		}
	}
}

func (p *Processor) ListAllTask() []Info {
	p.lock.RLock()
	defer p.lock.RUnlock()

	res := []Info{}

	for _, value := range p.infos {
		res = append(res, *value)
	}

	return res
}

func GeneratorTaskFunc(info *Info) func() {
	if info.Tasks == nil {
		return func() {
			utils.Logger.Infof("Empty task of %s", info.Name)
		}
	}
	return func() {
		for index, item := range info.Tasks {
			body := bytes.NewReader([]byte(item.TaskBody))
			request, err := http.NewRequest(item.TaskMethod, item.TaskUrl, body)
			if err != nil {
				utils.Logger.Errorf("Init task error of task: %s.%d, with err: %v", info.Name, index, err)
				continue
			}
			utils.Logger.Infof("Do task: %s.%d, URL: %s, METHOD: %s", info.Name, index, item.TaskUrl, item.TaskMethod)

			if info.SyncAble {
				go http.DefaultClient.Do(request)
			} else {
				_, _ = http.DefaultClient.Do(request)
			}
		}
	}
}
