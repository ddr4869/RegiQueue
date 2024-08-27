package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/ddr4869/RegiQueue/service"
)

type LoadTestRequest struct {
	Duration string `json:"duration"`
	Users    int    `json:"users"`
}

func RunLoadTest(w http.ResponseWriter, r *http.Request) {
	var req LoadTestRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	service.RestoreCourseEnrollment()
	// Prepare k6 command
	cmd := exec.Command("k6", "run", "--vus", fmt.Sprintf("%d", req.Users), "--duration", req.Duration, "loadtest.js")

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to run load test: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the output
	w.Write(output)
}
