package repo

type PongRepo struct{}

func NewPongRepo() *PongRepo {
	return &PongRepo{}
}

func (pr *PongRepo) GetPong() string {
	return "bom wabg"
}
