package entity

type Persion struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Age       int8   `json:"age" binding:"gte=1, lte=130"`
}

type Video struct {
	Title       string  `json:"title" validate:"is-cool"`
	Description string  `json:"description"`
	URL         string  `json:"url" binding:"required"`
	Author      Persion `json:"author" binding:"required"`
}
