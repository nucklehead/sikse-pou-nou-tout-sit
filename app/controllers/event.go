package controllers

import (
	"github.com/revel/revel"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"crypto/rand"
	"net/http"
	uuid "github.com/hashicorp/go-uuid"

)

var Events map[string]models.Event

type EventController struct {
	*revel.Controller
}

func (c EventController) Create(event models.Event) revel.Result {
	id, _ := uuid.GenerateUUID()
	Events[id] = event
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(event)
}


func (c EventController) Read(eventID string) revel.Result {
	return c.RenderJSON(Events[eventID])
}

func (c EventController) Update(eventID string, event models.Event) revel.Result {
	Events[eventID] = event
	return c.RenderJSON(event)
}

func (c EventController) Delete(eventID string) revel.Result {
	delete(Events, eventID)
	return c.RenderJSON("")
}

func (c EventController) List() revel.Result {
	return c.RenderJSON(Events)
}

func (c EventController) ShowList() revel.Result {
	return c.Render(Events)
}
