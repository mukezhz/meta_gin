package main

func NewUserDTOHandler() *DTOHandler[User, UserRequestDTO, UserResponseDTO] {
	return &DTOHandler[User, UserRequestDTO, UserResponseDTO]{
		ToModel: func(dto UserRequestDTO) User {
			return dto.ToModel()
		},
		FromModel: func(model User) UserResponseDTO {
			return FromModel(model)
		},
	}
}
