package api

import (
	"log"
	"net/http"

	"github.com/faxa0-0/billy/plan_service/internal/handler"
)

const prefix = "/api/v1"

type Api struct {
	srv *http.Server
}

func NewApi(handler handler.PlanHandler, address string) *Api {
	mux := http.NewServeMux()

	mux.HandleFunc(prefix+"/plans", handler.PlansHandler)
	mux.HandleFunc(prefix+"/plans/", handler.SinglePlanHandler)

	return &Api{&http.Server{
		Addr:    address,
		Handler: mux,
	}}

}
func (api *Api) Run() {
	log.Printf("Starting server at http://%s%s/plans", api.srv.Addr, prefix)
	if err := api.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
