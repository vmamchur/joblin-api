package route

import (
	"net/http"

	"github.com/vmamchur/vacancy-board/internal/handler"
)

func RegisterVacancyRoutes(mux *http.ServeMux, vacancyHandler *handler.VacancyHandler) {
	mux.Handle("GET /vacancies", http.HandlerFunc(vacancyHandler.Get))
}
