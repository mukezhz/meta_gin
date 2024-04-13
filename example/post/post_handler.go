package post

import "github.com/mukezhz/meta_gin/meta_gin"

func NewPostDTOHandler() *meta_gin.DTOHandler[Post, PostRequestDTO, PostResponseDTO] {
	return &meta_gin.DTOHandler[Post, PostRequestDTO, PostResponseDTO]{
		ToModel: func(dto PostRequestDTO) Post {
			return dto.ToModel()
		},
		FromModel: func(model Post) PostResponseDTO {
			return FromModel(model)
		},
	}
}
