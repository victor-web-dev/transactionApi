package http

import (
	"encoding/json"
	"finance-planner/adapter/http/transaction"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Init starts the api service
func Init() {

	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	http.HandleFunc("GET /health", healthCheckHandler)
	http.HandleFunc("GET /time", timeHandler)
	http.HandleFunc("GET /transactions", transaction.GetTransactions)
	http.HandleFunc("POST /transactions/create", transaction.CreateTransactions)

	//Metrics
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	fmt.Println("Listening on port 3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, time.Now().String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type healthResponse struct {
	Status string `json:"status"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response = healthResponse{Status: "ok"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
