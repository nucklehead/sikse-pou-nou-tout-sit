package controllers

import (
	"encoding/json"
	"errors"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
)

type EventController struct {
	*revel.Controller
}

func (c EventController) Index() revel.Result {
	var (
		events []models.Event
		err    error
	)
	events, err = models.GetEvents()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(events)
}

func (c EventController) Show(id string) revel.Result {
	var (
		event   models.Event
		err     error
		eventID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid event id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	eventID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid event id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	event, err = models.GetEvent(eventID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(event)
}

func (c EventController) Create() revel.Result {
	var (
		event models.Event
		err   error
	)

	err = json.NewDecoder(c.Request.Body).Decode(&event)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	event, err = models.AddEvent(event)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(event)
}

func (c EventController) Update() revel.Result {
	var (
		event models.Event
		err   error
	)
	err = json.NewDecoder(c.Request.Body).Decode(&event)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = event.UpdateEvent()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(event)
}

func (c EventController) Delete(id string) revel.Result {
	var (
		err     error
		event   models.Event
		eventID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid event id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	eventID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid event id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	event, err = models.GetEvent(eventID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = event.DeleteEvent()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}

func (c EventController) ShowList() revel.Result {
	events, err := models.GetEvents()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.Render(events)
}