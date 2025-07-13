package handler

import (
	"net/http"
	"strconv"

	"github.com/vmamchur/vacancy-board/internal/model"
	"github.com/vmamchur/vacancy-board/internal/service"
	"github.com/vmamchur/vacancy-board/pkg/httputil"
)

type VacancyHandler struct {
	vacancyService *service.VacancyService
}

func NewVacancyHandler(vacancyService *service.VacancyService) *VacancyHandler {
	return &VacancyHandler{vacancyService: vacancyService}
}

func (h *VacancyHandler) Get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	search := query.Get("search")
	limitRaw := query.Get("limit")
	offsetRaw := query.Get("offset")

	limit, err := strconv.Atoi(limitRaw)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetRaw)
	if err != nil {
		offset = 0
	}

	vacancies, err := h.vacancyService.GetAll(r.Context(), model.GetVacanciesFilter{
		Search: search,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		httputil.RespondWithError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	httputil.RespondWithJSON(w, http.StatusOK, vacancies)
}
