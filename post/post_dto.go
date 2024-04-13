package post

type PostRequestDTO struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (dto PostRequestDTO) ToModel() Post {
	return Post{
		Name:  dto.Name,
		Email: dto.Email,
	}
}

type PostResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromModel(model Post) PostResponseDTO {
	return PostResponseDTO{
		ID:    model.ID,
		Name:  model.Name,
		Email: model.Email,
	}
}
