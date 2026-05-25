package lessons

import (
	"context"
	"fmt"

	"github.com/snykk/go-rest-boilerplate/internal/apperror"
	"github.com/snykk/go-rest-boilerplate/internal/http/datatransfers/responses"
)

func (uc *usecase) List(ctx context.Context) ([]responses.LessonListItemResponse, error) {
	rows, err := uc.repo.List(ctx)
	if err != nil {
		return nil, apperror.InternalCause(fmt.Errorf("lessons.List: %w", err))
	}

	result := make([]responses.LessonListItemResponse, 0, len(rows))
	for _, r := range rows {
		result = append(result, toListItem(r))
	}
	return result, nil
}
