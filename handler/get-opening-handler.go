package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jdgabriel/go-learning/schema"
)

func GetOpeningHandler(ctx *gin.Context) {
	openings := []schema.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
