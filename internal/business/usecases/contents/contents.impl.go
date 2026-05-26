package contents

import (
	repointerface "github.com/snykk/go-rest-boilerplate/internal/datasources/repositories/interface"
)

type usecase struct {
	repo repointerface.ContentRepository
}

func NewUsecase(repo repointerface.ContentRepository) Usecase {
	return &usecase{repo: repo}
}
