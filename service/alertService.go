package service

import (
	"log"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type AlertService interface {
	SetAlert(c echo.Context, message string) error
	GetAlert(c echo.Context) map[string]interface{}
}

type alertService struct {
	store *sessions.CookieStore
}

func NewAlertService(store *sessions.CookieStore) AlertService {
	return &alertService{
		store: store,
	}
}

func (s *alertService) SetAlert(c echo.Context, message string) error {
	session, err := s.store.Get(c.Request(), "alert-st")
	if err != nil {
		return err
	}

	session.AddFlash(message)

	return sessions.Save(c.Request(), c.Response())
}

func (s *alertService) GetAlert(c echo.Context) map[string]interface{} {
	session, err := s.store.Get(c.Request(), "alert-st")
	if err != nil {
		log.Println(err)
		return nil
	}

	data := make(map[string]interface{})
	flashes := session.Flashes()

	if len(flashes) > 0 {
		data["is_alert"] = true
		data["message"] = flashes[0]
	} else {
		data["is_alert"] = false
		data["message"] = nil
	}

	session.Save(c.Request(), c.Response())

	return data
}
