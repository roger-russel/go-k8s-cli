package api

import (
	"encoding/json"
	"net/http"

	log "github.com/google/logger"
	"github.com/roger-russel/go-k8s-cli/internal/k8s"
	"k8s.io/client-go/rest"
)

type podsResponse struct {
	Name     string `json:"name"`
	PodCount int    `json:"podCount"`
}

type nodesResponse struct {
	Name      string `json:"name"`
	NodeCount int    `json:"nodeCount"`
}

func podsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	config, err := rest.InClusterConfig()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("fail"))
		log.Errorf("failed to get cluster config: %v", err)
		return
	}

	cli := k8s.NewClient(config)

	total, err := cli.CountPodsNumber()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("fail"))
		log.Errorf("failed to count pods: %v", err)
		return
	}

	resp := podsResponse{
		Name:     "TBD",
		PodCount: total,
	}

	respByte, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("fail"))
		log.Errorf("error marshalling response: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respByte))
}

func nodesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	config, err := rest.InClusterConfig()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("fail"))
		log.Errorf("failed to get cluster config")
		return
	}

	cli := k8s.NewClient(config)
	total, err := cli.CountNodesNumber()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("fail"))
		log.Errorf("failed to count pods: %v", err)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("fail"))
		log.Errorf("failed to get cluster config")
		return
	}

	resp := nodesResponse{
		Name:      "TBD",
		NodeCount: total,
	}

	respByte, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("fail"))
		log.Errorf("error marshalling response: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respByte))

	w.WriteHeader(http.StatusOK)
}
