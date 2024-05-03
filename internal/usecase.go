package internal

type Usecase struct{}

type UsecaseInput struct {
	Array []string
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (u *Usecase) Execute(input UsecaseInput) {
	jobs := make([]Job, 0)

	for i := 0; i < len(input.Array); i++ {
		jobs = append(jobs, Job{Adapter: *NewAdapter(), Index: i})
	}

	wp := NewWorkerPool(jobs, 20)

	wp.Run()
}
