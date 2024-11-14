package zface

type IServer interface {
	Start()
	Stop()
	Serve()
}
