package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	duration := time.Until(d)
	return int64(duration.Hours()) / 24
}
