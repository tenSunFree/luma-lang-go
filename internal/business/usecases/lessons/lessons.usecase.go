package lessons

import (
	"context"

	"github.com/snykk/go-rest-boilerplate/internal/http/datatransfers/responses"
)

// Usecase is the input boundary for lessons.
// Directly return the response DTO (instead of the domain struct) because lessons currently
// do not have complex business logic; adding another layer of domain mapping would only increase boilerplate code.
// We will improve this in the future if business logic (payment verification, progress tracking) emerges.
type Usecase interface {
	List(ctx context.Context) ([]responses.LessonListItemResponse, error)
	GetDetail(ctx context.Context, id string) (responses.LessonDetailResponse, error)
}
