package models

type Product struct {
	ID          int    `json:"id"`
	KodeProduct uint   `json:"Kode"`
	NameProduct string `json:"name"`
	QtyProduct  uint   `json:"qty"`
}