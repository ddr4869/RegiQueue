package service

import (
	"context"
	"sync"
	"time"
)

type CourseInfo struct {
	Name     string
	Capacity int
	Enrolled int
	mu       sync.Mutex
}

var courses = map[string]*CourseInfo{
	"Math101": {Name: "Math101", Capacity: 20000, Enrolled: 0},
	"Eng201":  {Name: "Eng201", Capacity: 30000, Enrolled: 0},
	// Add more courses as needed
}

type RegistrationRequest struct {
	UserID      string `json:"user_id"`
	CourseName  string `json:"course_name"`
	StudentName string `json:"student_name"`
}

func ProcessRegistration(ctx context.Context, req RegistrationRequest) bool {
	course, exists := courses[req.CourseName]
	if !exists {
		return false
	}

	course.mu.Lock()
	defer course.mu.Unlock()

	if course.Enrolled < course.Capacity {
		// Simulate DB insertion
		//time.Sleep(time.Duration(rand.Intn(91)+10) * time.Millisecond)
		time.Sleep(time.Duration(15 * time.Millisecond))
		course.Enrolled++
		return true
	}
	return false
}

func RestoreCourseEnrollment() {
	for _, course := range courses {
		course.Enrolled = 0
	}
}

func GetAllCourseInfo() map[string]*CourseInfo {
	return courses
}

func CanEnroll(courseName string) bool {
	course, exists := courses[courseName]
	if !exists {
		return false
	}
	return course.Enrolled < course.Capacity
}
