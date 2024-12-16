package api

import (
	"log"
	"net/http"

	"github.com/faxa0-0/billy/plan_service/internal/handler"
)

type Api struct {
	srv *http.Server
}

func NewApi(handler handler.PlanHandler, address string) *Api {
	mux := http.NewServeMux()

	mux.HandleFunc("/plans", handler.PlansHandler)
	mux.HandleFunc("/plans/", handler.SinglePlanHandler)

	return &Api{&http.Server{
		Addr:    address,
		Handler: mux,
	}}

}
func (api *Api) Run() {
	log.Printf("Starting server at http://%s/plans", api.srv.Addr)
	if err := api.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
