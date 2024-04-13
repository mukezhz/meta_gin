package user

type UserRequestDTO struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (dto UserRequestDTO) ToModel() User {
	return User{
		Name:  dto.Name,
		Email: dto.Email,
	}
}

type UserResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromModel(model User) UserResponseDTO {
	return UserResponseDTO{
		ID:    model.ID,
		Name:  model.Name,
		Email: model.Email,
	}
}
