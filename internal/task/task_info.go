package task

// Info save all value of one task.
type Info struct {

	// The task name.
	Name string `json:"name"`

	// The holder define the runner in which TM(Trace-mocker). The value is the NodeId in config.Application.NodeId.
	Holder string `json:"holder"`

	// The time of task to run.
	Cron string `json:"cron"`

	// Task values =======
	Tasks []*Obj `json:"tasks"`
}

type Obj struct {
	TaskUrl    string `json:"task_url"`
	TaskMethod string `json:"task_method"`
	TaskBody   string `json:"task_body"`
}


