package domain

import (
	apiModels "dev11/internal/api/models"
	dModels "dev11/internal/domain/models"
	"dev11/internal/repository/storage"

	"github.com/sirupsen/logrus"
)

type Domain struct {
	log     *logrus.Logger
	storage *storage.Storage
}

func NewDomain(storage *storage.Storage, log *logrus.Logger) *Domain {
	return &Domain{
		log:     log,
		storage: storage,
	}
}

func (d Domain) AddNewEvent(params apiModels.Params) error {

	event := dModels.NewEvent(params.Date, params.Description, params.Title)

	err := d.storage.Add(params.UserId, *event)
	if err != nil {
		d.log.WithError(err).Errorf("failed to add new event. ID: %v, Date: %v, title: %v, desc: %v", params.UserId, params.Date, params.Title, params.Description)
		return err
	}
	d.log.Infof("add new event. ID: %v, Date: %v, title: %v, desc: %v", params.UserId, params.Date, params.Title, params.Description)
	return nil
}

func (d Domain) SetEvent(params apiModels.Params) error {

	event := dModels.NewEvent(params.Date, params.Description, params.Title)

	err := d.storage.Set(params.UserId, *event)
	if err != nil {
		d.log.WithError(err).Errorf("failed to set event. ID: %v Date: %v title: %v desc: %v", params.UserId, params.Date, params.Title, params.Description)
		return err
	}
	d.log.Infof("set event. ID: %v, Date: %v, title: %v, desc: %v", params.UserId, params.Date, params.Title, params.Description)
	return nil
}

func (d Domain) DeleteEvent(params apiModels.Params) error {

	err := d.storage.Delete(params.UserId, params.Date)
	if err != nil {
		d.log.WithError(err).Errorf("failed to delete event. ID: %v Date: %v", params.UserId, params.Date)
		return err
	}
	d.log.Infof("delete event. ID: %v, Date: %v", params.UserId, params.Date)
	return nil
}

func (d Domain) GetEventsForDay(params apiModels.Params) ([]dModels.Event, error) {

	events, err := d.storage.GetEventsForDay(params.UserId, params.Date)
	if err != nil {
		d.log.WithError(err).Errorf("failed to get events for given day. ID: %v Date: %v", params.UserId, params.Date)
		return nil, err
	}
	d.log.Infof("get event for given day. ID: %v, Date: %v", params.UserId, params.Date)

	return events, nil
}

func (d Domain) GetEventsForWeek(params apiModels.Params) ([]dModels.Event, error) {

	events, err := d.storage.GetEventsForWeek(params.UserId, params.Date)
	if err != nil {
		d.log.WithError(err).Errorf("failed to get events for given week. ID: %v Date: %v", params.UserId, params.Date)
		return nil, err
	}
	d.log.Infof("get event for given weem. ID: %v, Date: %v", params.UserId, params.Date)

	return events, nil
}

func (d Domain) GetEventsForMonth(params apiModels.Params) ([]dModels.Event, error) {

	events, err := d.storage.GetEventsForMonth(params.UserId, params.Date)
	if err != nil {
		d.log.WithError(err).Errorf("failed to get events for given month. ID: %v Date: %v", params.UserId, params.Date)
		return nil, err
	}
	d.log.Infof("get event for given month. ID: %v, Date: %v", params.UserId, params.Date)

	return events, nil
}
