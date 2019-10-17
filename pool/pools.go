package pool

type ConnPool interface {
	Get() ConnPool
}
