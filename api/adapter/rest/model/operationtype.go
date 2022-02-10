package model

type OperatoinTypeModel struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Operation   byte   `json:"operation"`
}
