package internal

import (
	"log"
	"net/http"
	"os"
)

func SetUp(s *Server) {

	r := s.router
	api := r.Group("/api")
	api.POST("/register", s.Register)
	api.GET("/queue_position", s.GetQueuePosition)
	api.POST("/run_load_test", s.RunLoadTest)
	api.GET("/course_info", s.CourseInfo)
}

func (s *Server) Start() error {

	srv := &http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: s.router,
	}

	log.Printf("Listening and serving HTTP on %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("listen: %s\n", err)
		return err
	}

	return nil
}
