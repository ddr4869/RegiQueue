package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ddr4869/RegiQueue/internal/dto"
	"github.com/ddr4869/RegiQueue/internal/service"
	"github.com/ddr4869/RegiQueue/kafka"
	"github.com/ddr4869/RegiQueue/redis"
	"github.com/gin-gonic/gin"
)

type QueuePositionResponse struct {
	Position int64 `json:"position"`
}

type RegistrationRequest struct {
	UserID      string `json:"user_id"`
	CourseName  string `json:"course_name"`
	StudentName string `json:"student_name"`
}

// Register 요청 처리 핸들러
func (s *Server) Register(c *gin.Context) {
	var req RegistrationRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if !service.CanEnroll(req.CourseName) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course is full"})
		return
	}
	ctx := context.Background()

	// marshal the request
	reqBytes, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request"})
		return
	}
	// Kafka에 메시지 전송
	err = kafka.SendMessage(os.Getenv("KAFKA_TOPIC"), reqBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message to Kafka"})
		return
	}

	// Redis에서 대기열에 추가하고, 현재 대기열의 위치를 확인
	queuePosition, err := redis.IncrementQueue(ctx, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to increment queue"})
		return
	}

	// 사용자에게 대기열 위치를 반환
	dto.NewSuccessResponse(c, fmt.Sprintf("Your registration is in progress. There are %d users ahead of you.", queuePosition-1))
}

// 사용자의 대기열 위치 확인 핸들러 GetQueuePosition
func (s *Server) GetQueuePosition(c *gin.Context) {
	userId := c.Query("user_id")
	if userId == "" {
		dto.NewErrorResponse(c, http.StatusBadRequest, fmt.Errorf("user_id is required"), "user_id is required")
		return
	}

	ctx := context.Background()

	// Redis에서 대기열 위치 확인
	position, err := redis.GetQueuePosition(ctx, userId)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to get queue position")
		return
	}

	if position == 0 {
		c.JSON(http.StatusOK, gin.H{"redirect_url": fmt.Sprintf("/register?user_id=%s", userId)})
		return
	}

	// 대기열 위치 반환
	response := QueuePositionResponse{Position: position}
	dto.NewSuccessResponse(c, response)
}

func (s *Server) CourseInfo(c *gin.Context) {
	dto.NewSuccessResponse(c, service.GetAllCourseInfo())
}
