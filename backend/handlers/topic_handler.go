package handlers

import (
	"net/http"
	"sustainability-platform/services"
	"sustainability-platform/utils"
)

type TopicHandler struct {
	Service services.TopicService
}

func (h *TopicHandler) GetTopics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	topics, err := h.Service.FetchTopics(ctx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch topics")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"topics": topics})
}
