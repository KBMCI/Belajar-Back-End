package book

// respon json dikirim dengan format huruf kecil dan jika terdiri 2 kata dipisah dengan _
type BookResponse struct{
	ID          int `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int `json:"price"`
	Rating      int `json:"rating"`
}