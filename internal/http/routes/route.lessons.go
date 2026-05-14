package routes

import (
	"github.com/gin-gonic/gin"
	lessonshandler "github.com/snykk/go-rest-boilerplate/internal/http/handlers/v1/lessons"
)

// lessonsRoute wires the /lessons/* group.
// No auth middleware required — lessons are public endpoints for MVP.
type lessonsRoute struct {
	handler lessonshandler.Handler
	router  *gin.RouterGroup
}

// NewLessonsRoute builds the route module.
func NewLessonsRoute(router *gin.RouterGroup) *lessonsRoute {
	return &lessonsRoute{
		handler: lessonshandler.NewHandler(),
		router:  router,
	}
}

// Routes mounts the /lessons group and its endpoints.
func (r *lessonsRoute) Routes() {
	v1 := r.router.Group("/v1")
	lessonsGrp := v1.Group("/lessons")

	// GET /api/v1/lessons → Movie List
	lessonsGrp.GET("", r.handler.ListLessons)

	// GET /api/v1/lessons/:lessonId → Video details (including subtitles and single words)
	lessonsGrp.GET("/:lessonId", r.handler.GetLessonDetail)
}
