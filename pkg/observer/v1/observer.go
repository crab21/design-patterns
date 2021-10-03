package v1

type Observer interface {
	update(string)
	getID() string
}
