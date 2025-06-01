package route

import (
	"net/http"

	"github.com/vmamchur/vacancy-board/internal/handler"
)

func NewRouter(appSecret string, authHandler *handler.AuthHandler, vacancyHandler *handler.VacancyHandler) http.Handler {
	mux := http.NewServeMux()

	RegisterAuthRoutes(mux, authHandler, appSecret)
	RegisterVacancyRoutes(mux, vacancyHandler)

	return mux
}
