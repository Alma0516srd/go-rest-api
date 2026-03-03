package apiserver

type Apiserver struct {
}

func New() *Apiserver {
	return &Apiserver{}
}

func (s *Apiserver) Start() error {
	return nil
}
