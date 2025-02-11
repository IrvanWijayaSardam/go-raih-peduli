package dtos

import "mime/multipart"

type InputFundraise struct {
	Title 		string 			`json:"title" form:"title" validate:"required"`
	Description string 			`json:"description" form:"description" validate:"required"`
	Photo	  	multipart.File  `json:"photo" form:"photo"`
	Target    	int32  			`json:"target" form:"target" validate:"required"`
	StartDate 	string 			`json:"start_date" form:"start_date" validate:"required"`
	EndDate   	string 			`json:"end_date" form:"end_date" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}