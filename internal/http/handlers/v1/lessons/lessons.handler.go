package lessons

import (
	"net/http"

	"github.com/gin-gonic/gin"
	lessonsuc "github.com/snykk/go-rest-boilerplate/internal/business/usecases/lessons"
	v1 "github.com/snykk/go-rest-boilerplate/internal/http/handlers/v1"
	"github.com/snykk/go-rest-boilerplate/pkg/logger"
)

type Handler struct {
	usecase lessonsuc.Usecase
}

func NewHandler(usecase lessonsuc.Usecase) Handler {
	return Handler{usecase: usecase}
}

// ListLessons godoc
// @Summary      影片列表 / Get lesson list
// @Description  Returns all published video lessons for the home/list page.
// @Tags         lessons
// @Produce      json
// @Success      200  {object}  v1.BaseResponse{data=map[string][]responses.LessonListItemResponse}
// @Router       /lessons [get]
func (h Handler) ListLessons(ctx *gin.Context) {
	lessons, err := h.usecase.List(ctx.Request.Context())
	if err != nil {
		logger.ErrorWithContext(ctx.Request.Context(), "ListLessons failed", logger.Fields{
			"error": err.Error(),
		})
		v1.RespondWithError(ctx, err)
		return
	}

	v1.NewSuccessResponse(ctx, http.StatusOK, "lessons fetched successfully", map[string]interface{}{
		"lessons": lessons,
	})
}

// GetLessonDetail godoc
// @Summary      影片詳情 / Get lesson detail
// @Description  Returns full lesson data: lesson info, playback config, captions, and vocabulary.
// @Tags         lessons
// @Produce      json
// @Param        lessonId  path      string  true  "Lesson ID (e.g. starbucks_001)"
// @Success      200       {object}  v1.BaseResponse{data=responses.LessonDetailResponse}
// @Failure      404       {object}  v1.BaseResponse
// @Router       /lessons/{lessonId} [get]
func (h Handler) GetLessonDetail(ctx *gin.Context) {
	id := ctx.Param("lessonId")
	detail, err := h.usecase.GetDetail(ctx.Request.Context(), id)
	if err != nil {
		logger.ErrorWithContext(ctx.Request.Context(), "GetLessonDetail failed", logger.Fields{
			"error": err.Error(),
			"id":    id,
		})
		v1.RespondWithError(ctx, err)
		return
	}
	v1.NewSuccessResponse(ctx, http.StatusOK, "lesson detail fetched successfully", detail)
}
