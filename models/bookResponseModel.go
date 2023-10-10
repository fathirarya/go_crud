package models

import "time"

type BookResponse struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Publisher string    `json:"publisher"`
}

func (updateBook Book) ResponseConvertBook() BookResponse {
	Response := BookResponse{}
	Response.ID = updateBook.ID
	Response.Title = updateBook.Title
	Response.Author = updateBook.Author
	Response.Publisher = updateBook.Publisher

	return Response
}
