package post

type CreateProductRequestDTO struct {
	Name        string  `json:"name" binding:"required" validate:"required"`
	Description string  `json:"description" binding:"required" validate:"required"`
	Price       float64 `json:"price" binding:"required" validate:"required"`
	Category    string  `json:"category" binding:"required" validate:"required"`
}

type CreateProductResponseDTO struct {
	ID          uint    `json:"id" binding:"required" validate:"required"`
	Name        string  `json:"name" binding:"required" validate:"required"`
	Description string  `json:"description" binding:"required" validate:"required"`
	Price       float64 `json:"price" binding:"required" validate:"required"`
	Category    string  `json:"category" binding:"required" validate:"required"`
}

type CreatePostHandler struct{}

func NewCreateHandler() *CreatePostHandler {
	return &CreatePostHandler{}
}

func (h CreatePostHandler) ToModel(dto CreateProductRequestDTO) Post {
	return Post{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Category:    dto.Category,
	}
}

func (h CreatePostHandler) FromModel(model Post) CreateProductResponseDTO {
	return CreateProductResponseDTO{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		Category:    model.Category,
	}
}

type ReadProductRequestDTO struct{}
type ReadProductResponseDTO struct {
	ID          uint    `json:"id" binding:"required" validate:"required"`
	Name        string  `json:"name" binding:"required" validate:"required"`
	Description string  `json:"description" binding:"required" validate:"required"`
	Price       float64 `json:"price" binding:"required" validate:"required"`
	Category    string  `json:"category" binding:"required" validate:"required"`
}
type ReadProductHandler struct{}

func NewReadHandler() *ReadProductHandler {
	return &ReadProductHandler{}
}
func (h ReadProductHandler) ToModel(dto ReadProductRequestDTO) Post {
	return Post{}
}

func (h ReadProductHandler) FromModel(model Post) ReadProductResponseDTO {
	return ReadProductResponseDTO{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		Category:    model.Category,
	}
}

type UpdateProductRequestDTO struct {
	Name        string  `json:"name" binding:"required" validate:"required"`
	Description string  `json:"description" binding:"required" validate:"required"`
	Price       float64 `json:"price" binding:"required" validate:"required"`
	Category    string  `json:"category" binding:"required" validate:"required"`
}

type UpdateProductResponseDTO struct {
	ID          uint    `json:"id" binding:"required" validate:"required"`
	Name        string  `json:"name" binding:"required" validate:"required"`
	Description string  `json:"description" binding:"required" validate:"required"`
	Price       float64 `json:"price" binding:"required" validate:"required"`
	Category    string  `json:"category" binding:"required" validate:"required"`
}

type UpdatePostHandler struct{}

func NewUpdateHandler() *UpdatePostHandler {
	return &UpdatePostHandler{}
}

func (h UpdatePostHandler) ToModel(dto UpdateProductRequestDTO) Post {
	return Post{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Category:    dto.Category,
	}
}

func (h UpdatePostHandler) FromModel(model Post) UpdateProductResponseDTO {
	return UpdateProductResponseDTO{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		Category:    model.Category,
	}
}

type DeleteProductRequestDTO struct{}
type DeleteProductResponseDTO struct {
	ID uint `json:"id" binding:"required" validate:"required"`
}

type DeletePostHandler struct{}

func NewDeleteHandler() *DeletePostHandler {
	return &DeletePostHandler{}
}

func (h DeletePostHandler) ToModel(dto DeleteProductRequestDTO) Post {
	return Post{}
}

func (h DeletePostHandler) FromModel(model Post) DeleteProductResponseDTO {
	return DeleteProductResponseDTO{
		ID: model.ID,
	}
}
