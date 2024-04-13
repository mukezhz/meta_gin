package user

import "github.com/mukezhz/meta_gin/meta_gin"

func NewUserDTOHandler() *meta_gin.DTOHandler[User, UserRequestDTO, UserResponseDTO] {
	return &meta_gin.DTOHandler[User, UserRequestDTO, UserResponseDTO]{
		ToModel: func(dto UserRequestDTO) User {
			return dto.ToModel()
		},
		FromModel: func(model User) UserResponseDTO {
			return FromModel(model)
		},
	}
}
