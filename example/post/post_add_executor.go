package post

import (
	"context"
	"log"
)

type PostAddExecutor struct {
}

func (e *PostAddExecutor) Execute(context context.Context, model *Post) (*Post, error) {
	log.Println("PostAddExecutor.Execute")
	return model, nil
}

func NewPostAddExecutor() *PostAddExecutor {
	return &PostAddExecutor{}
}
