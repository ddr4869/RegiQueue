package internal

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/ddr4869/RegiQueue/internal/dto"
	"github.com/ddr4869/RegiQueue/internal/service"
	"github.com/gin-gonic/gin"
)

type LoadTestRequest struct {
	Duration string `json:"duration"`
	Users    int    `json:"users"`
}

func (s *Server) RunLoadTest(c *gin.Context) {

	var req LoadTestRequest
	err := c.BindJSON(&req)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Invalid request")
		return
	}
	service.RestoreCourseEnrollment()
	// Prepare k6 command
	cmd := exec.Command("k6", "run", "--vus", fmt.Sprintf("%d", req.Users), "--duration", req.Duration, "loadtest.js")

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to run load test")
		return
	}
	// Return the output
	dto.NewSuccessResponse(c, string(output))
}
