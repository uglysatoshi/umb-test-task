package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Endpoint struct {
	s Service
}

type Service interface {
	DaysLeft() int64
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.DaysLeft()
	s := fmt.Sprintf("Days left: %d", d)
	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}
