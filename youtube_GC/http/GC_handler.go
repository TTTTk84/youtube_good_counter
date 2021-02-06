package http

import (
	"github.com/TTTTk84/youtube_good_counter/domain"
	"github.com/gin-gonic/gin"
)


type GCHandler struct {
	gcusecase domain.GCUsecase
}

func NewGC_handler(c *gin.Engine, gu domain.GCUsecase) {
	handler := &GCHandler{
		gcusecase: gu,
	}

	c.POST("/GoodCounte", handler.AddGC)
	c.POST("/GCOnceday", handler.GCOnceday)
	c.POST("/GCOnceWeek", handler.GCOnceWeek)
}

func (h *GCHandler) AddGC(c *gin.Context) {

}

func (h *GCHandler) GCOnceday(c *gin.Context) {

}

func (h *GCHandler) GCOnceWeek(c *gin.Context) {
}
