package routes

import (
	"github.com/gin-gonic/gin"
	lessonsuc "github.com/snykk/go-rest-boilerplate/internal/business/usecases/lessons"
	lessonshandler "github.com/snykk/go-rest-boilerplate/internal/http/handlers/v1/lessons"
)

type lessonsRoute struct {
	handler lessonshandler.Handler
	router  *gin.RouterGroup
}

func NewLessonsRoute(router *gin.RouterGroup, lessonsUC lessonsuc.Usecase) *lessonsRoute {
	return &lessonsRoute{
		handler: lessonshandler.NewHandler(lessonsUC),
		router:  router,
	}
}

func (r *lessonsRoute) Routes() {
	v1 := r.router.Group("/v1")
	lessonsGrp := v1.Group("/lessons")
	// GET /api/v1/lessons
	lessonsGrp.GET("", r.handler.ListLessons)
	// GET /api/v1/lessons/:lessonId
	lessonsGrp.GET("/:lessonId", r.handler.GetLessonDetail)
}
