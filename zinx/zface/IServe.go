package zface

type Iserver interface {
	Start()
	Stop()
	Serve()
}
