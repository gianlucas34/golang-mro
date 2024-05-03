package internal

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle() {
	array := make([]string, 800)
	usecase := NewUsecase()

	usecase.Execute(UsecaseInput{Array: array})
}
