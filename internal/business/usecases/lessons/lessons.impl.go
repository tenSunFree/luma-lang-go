package lessons

import (
	repointerface "github.com/snykk/go-rest-boilerplate/internal/datasources/repositories/interface"
)

type usecase struct {
	repo repointerface.LessonRepository
}

func NewUsecase(repo repointerface.LessonRepository) Usecase {
	return &usecase{repo: repo}
}
