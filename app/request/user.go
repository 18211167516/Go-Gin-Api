package request

type User struct {
	Name      string `json:"name" xml:"name" form:"name" binding:"required"`
	CreatedBy string `json:"created_by" xml:"created_by" form:"created_by" binding:"lowercase"`
}

type UserId struct {
	ID int `uri:"id" binding:"required"`
}
