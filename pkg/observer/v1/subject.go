package v1

type subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}
