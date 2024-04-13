package post

type PostRequestDTO struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func (dto PostRequestDTO) ToModel() Post {
	return Post{
		Title:  dto.Title,
		Author: dto.Author,
	}
}

type PostResponseDTO struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FromModel(model Post) PostResponseDTO {
	return PostResponseDTO{
		ID:     model.ID,
		Title:  model.Title,
		Author: model.Author,
	}
}
