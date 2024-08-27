package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ddr4869/RegiQueue/kafka"
	"github.com/ddr4869/RegiQueue/redis"
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
func Register(w http.ResponseWriter, r *http.Request) {
	var req RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// marshal the request
	reqBytes, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to marshal request", http.StatusInternalServerError)
		return
	}
	// Kafka에 메시지 전송
	err = kafka.SendMessage("registration_topic", reqBytes)
	if err != nil {
		http.Error(w, "Failed to send message to Kafka", http.StatusInternalServerError)
		return
	}

	// Redis에서 대기열에 추가하고, 현재 대기열의 위치를 확인
	queuePosition, err := redis.IncrementQueue(ctx, req.UserID)
	if err != nil {
		http.Error(w, "Failed to add to queue", http.StatusInternalServerError)
		return
	}

	// 사용자에게 대기열 위치를 반환
	response := fmt.Sprintf("Your registration is in progress. There are %d users ahead of you.", queuePosition-1)
	w.Write([]byte(response))
}

// 사용자의 대기열 위치 확인 핸들러
func GetQueuePosition(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// Redis에서 대기열 위치 확인
	position, err := redis.GetQueuePosition(ctx, userId)
	if err != nil {
		http.Error(w, "Failed to get queue position", http.StatusInternalServerError)
		return
	}

	if position == 0 {
		// 사용자의 순번이 오면, 수강신청 페이지로 이동하도록 상태 코드 200과 함께 리다이렉트 URL 반환
		http.Redirect(w, r, "http://localhost:5500/regist", http.StatusSeeOther)
		return
	}

	// 대기열 위치 반환
	response := QueuePositionResponse{Position: position}
	json.NewEncoder(w).Encode(response)
}
