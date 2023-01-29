package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	CreateBook(bookRequest BookRequest) (Book, error)
	UpdateBook(ID int, updateBookRequest UpdateBookRequest) (Book, error)
	DeleteBook(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.Get()
	return books, err
	// cara singkat spt dibawah, itu karena return value dari repository.Get() itu sama" mengembalikan book dan error
	// return s.repository.Get()
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.GetById(ID)
	return book, err
}

func (s *service) CreateBook(bookRequest BookRequest) (Book, error) {

	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       bookRequest.Price,
		Rating:      bookRequest.Rating,
		Discount:    bookRequest.Discount,
	}

	newBook, err := s.repository.Save(book)
	return newBook, err
}

func (s *service) UpdateBook(ID int, updateBookRequest UpdateBookRequest) (Book, error) {
	book, err := s.repository.GetById(ID)

	book.Title = updateBookRequest.Title
	book.Description = updateBookRequest.Description
	book.Price = updateBookRequest.Price
	book.Rating = updateBookRequest.Rating
	book.Discount = updateBookRequest.Discount

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) DeleteBook(ID int) (Book, error) {
	book, err := s.repository.GetById(ID)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}