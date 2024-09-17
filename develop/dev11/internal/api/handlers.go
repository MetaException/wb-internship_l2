package api

import (
	"dev11/internal/api/models"
	"dev11/internal/domain"
	"dev11/pkg/utils"
	"net/http"

	"github.com/sirupsen/logrus"
)

type calendarHandler struct {
	log *logrus.Logger
	d   *domain.Domain
}

func NewCalendarHandler(log *logrus.Logger, domain *domain.Domain) *calendarHandler {
	return &calendarHandler{
		log: log,
		d:   domain,
	}
}

func (ch calendarHandler) GetEventsForDay(w http.ResponseWriter, req *http.Request) {

	userId := req.URL.Query().Get("user_id")
	date := req.URL.Query().Get("date")

	params, err := models.NewParams(userId, date, "", "")
	if err != nil {
		utils.Send(w, err, http.StatusBadRequest)
		ch.log.Error(err)
		return
	}

	events, err := ch.d.GetEventsForDay(*params)
	if err != nil {
		utils.Send(w, err, http.StatusServiceUnavailable)
		return
	}

	utils.Send(w, events, http.StatusOK)
}

func (ch calendarHandler) GetEventsForWeek(w http.ResponseWriter, req *http.Request) {

	userId := req.URL.Query().Get("user_id")
	date := req.URL.Query().Get("date")

	params, err := models.NewParams(userId, date, "", "")
	if err != nil {
		utils.Send(w, err, http.StatusBadRequest)
		ch.log.Error(err)
		return
	}

	events, err := ch.d.GetEventsForWeek(*params)
	if err != nil {
		utils.Send(w, err, http.StatusServiceUnavailable)
		return
	}

	utils.Send(w, events, http.StatusOK)
}

func (ch calendarHandler) GetEventsForMonth(w http.ResponseWriter, req *http.Request) {
	userId := req.URL.Query().Get("user_id")
	date := req.URL.Query().Get("date")

	params, err := models.NewParams(userId, date, "", "")
	if err != nil {
		utils.Send(w, err, http.StatusBadRequest)
		ch.log.Error(err)
		return
	}

	events, err := ch.d.GetEventsForMonth(*params)
	if err != nil {
		utils.Send(w, err, http.StatusServiceUnavailable)
		return
	}

	utils.Send(w, events, http.StatusOK)
}

func (ch calendarHandler) CreateEvent(w http.ResponseWriter, req *http.Request) {

	date := req.FormValue("date")
	userId := req.FormValue("user_id")
	description := req.FormValue("description")
	title := req.FormValue("title")

	params, err := models.NewParams(userId, date, description, title)
	if err != nil {
		utils.Send(w, err, http.StatusBadRequest)
		ch.log.Error(err)
		return
	}

	err = ch.d.AddNewEvent(*params)
	if err != nil {
		utils.Send(w, err, http.StatusServiceUnavailable)
		return
	}

	utils.Send(w, "OK", http.StatusOK)
}

func (ch calendarHandler) UpdateEvent(w http.ResponseWriter, req *http.Request) {

	date := req.FormValue("date")
	userId := req.FormValue("user_id")
	description := req.FormValue("description")
	title := req.FormValue("title")

	params, err := models.NewParams(userId, date, description, title)
	if err != nil {
		utils.Send(w, err, http.StatusBadRequest)
		ch.log.Error(err)
		return
	}

	err = ch.d.SetEvent(*params)
	if err != nil {
		utils.Send(w, err, http.StatusServiceUnavailable)
		return
	}

	utils.Send(w, "OK", http.StatusOK)
}

func (ch calendarHandler) DeleteEvent(w http.ResponseWriter, req *http.Request) {

	date := req.FormValue("date")
	userId := req.FormValue("user_id")
	description := req.FormValue("description")
	title := req.FormValue("title")

	params, err := models.NewParams(userId, date, description, title)
	if err != nil {
		utils.Send(w, err, http.StatusBadRequest)
		ch.log.Error(err)
		return
	}

	err = ch.d.DeleteEvent(*params)
	if err != nil {
		utils.Send(w, err, http.StatusServiceUnavailable)
		return
	}

	utils.Send(w, "OK", http.StatusOK)
}
