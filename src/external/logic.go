package external

type ILogic interface {
	Init() error
	Run()
}
