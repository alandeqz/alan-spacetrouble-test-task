package models

type Paging struct {
	Limit  *int `form:"limit"`
	Offset *int `form:"offset"`
}
