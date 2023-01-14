package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
}

type service struct{
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       int(price),
		Rating: bookRequest.Rating,
	}
	return s.repository.Create(book)
}