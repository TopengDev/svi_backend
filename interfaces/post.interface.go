package interfaces

type ICreatePostDTO struct {
	Title string
	Content string
	Category string
	Status string
}

type IUpdatePostDTO struct {
	Id int
	Title string
	Content string
	Category string
	Status string
}
