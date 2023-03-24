package workflow

type Workflow interface {
	Run() (interface{}, error)
}
