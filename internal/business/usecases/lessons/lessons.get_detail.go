package lessons

import (
	"context"
	"fmt"

	"github.com/snykk/go-rest-boilerplate/internal/apperror"
	"github.com/snykk/go-rest-boilerplate/internal/http/datatransfers/responses"
)

func (uc *usecase) GetDetail(ctx context.Context, id string) (responses.LessonDetailResponse, error) {
	row, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		// NotFound is passed directly, causing the handler to return a 404 error.
		return responses.LessonDetailResponse{}, err
	}
	detail, err := toDetail(row)
	if err != nil {
		return responses.LessonDetailResponse{}, apperror.InternalCause(
			fmt.Errorf("lessons.GetDetail marshal %s: %w", id, err),
		)
	}
	return detail, nil
}
