package handler

import (
	"net/http"

	"github.com/vmamchur/vacancy-board/internal/service"
	"github.com/vmamchur/vacancy-board/pkg/httputil"
)

type VacancyHandler struct {
	vacancyService *service.VacancyService
}

func NewVacancyHandler(vacancyService *service.VacancyService) *VacancyHandler {
	return &VacancyHandler{vacancyService: vacancyService}
}

func (h *VacancyHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	vacancies, err := h.vacancyService.GetAll(r.Context())
	if err != nil {
		httputil.RespondWithError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	httputil.RespondWithJSON(w, http.StatusOK, vacancies)
}
