package models

type Book struct {
	ID                int    `json:"id"`
	ISBN              int64  `json:"isbn"`
	Title             string `json:"book_title"`
	Author            string `json:"book_author"`
	YearOfPublication string `json:"year_of_publication"`
	Publisher         string `json:"publisher"`
	ImageURLS         string `json:"image_url_s"`
	ImageURLM         string `json:"image_url_m"`
	ImageURLL         string `json:"image_url_l"`
	Status            int    `json:"book_status"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}
