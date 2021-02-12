package main

import (
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/savoiringfaire/overseerr-prom-exporter/lib/overseerr"
	log "github.com/sirupsen/logrus"
)

var (
	pendingRequests = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "overseerr_pending_requests",
		Help: "The number of requests pending in overseerr",
	})

	approvedRequests = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "overseerr_approved_requests",
		Help: "The number of requests approved in overseerr",
	})
)

func refreshRequestCount(client *overseerr.Overseerr) {
	requests, err := client.RequestCount()
	if err != nil {
		log.Errorf("Error getting request counts: %s", err)
		return
	}

	log.Infof("Approved requests: %d", requests.Approved)
	log.Infof("Pending requests: %d", requests.Pending)

	approvedRequests.Set(float64(requests.Approved))
	pendingRequests.Set(float64(requests.Pending))
}

func recordMetrics(client *overseerr.Overseerr) {
	go func() {
		for {
			refreshRequestCount(client)
			time.Sleep(20 * time.Second)
		}
	}()
}

func main() {
	baseURL := os.Getenv("OVERSEERR_API_BASE_URL")
	if baseURL == "" {
		log.Fatalf("No base URL for overseer API set.")
	}

	apiKey := os.Getenv("OVERSEERR_API_KEY")
	if apiKey == "" {
		log.Fatalf("No API key for overseer API set.")
	}

	client := overseerr.New(baseURL, apiKey)

	recordMetrics(&client)

	promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   "overseerr_exporter",
		Subsystem:   "internal",
		Name:        "build_info",
		Help:        "Build information.",
		ConstLabels: prometheus.Labels{"branch": "master"},
	})

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
