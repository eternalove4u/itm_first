package service

type producer interface {
	Produce() ([]string, error)
}

type presenter interface {
	Present([]string) error
}
