package book

// struct untuk menerima data dari body POST books postman
type BookRequest struct {
	Title 		string	`json:"title" binding:"required"`
	Description string	`json:"description" binding:"required"`
	// json.Number untuk validasi error jika yg diinputkan bukan number/untuk menghindari internal server error
	Price 		int		`json:"price" binding:"required"`
	Rating 		int		`json:"rating" binding:"required"`	
	Discount	int		`json:"discount" binding:"required"`

}
type UpdateBookRequest struct {
	Title 		string	`json:"title"`
	Description string	`json:"description"`
	Price 		int		`json:"price"`
	Rating 		int		`json:"rating"`	
	Discount	int		`json:"discount"`
}