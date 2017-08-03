package controllers

import (
	"net/http"

	uuid "github.com/hashicorp/go-uuid"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"github.com/revel/revel"
)

var Events = map[string]models.Event{}

type EventController struct {
	*revel.Controller
}

func (c EventController) Create(event models.Event) revel.Result {
	id, _ := uuid.GenerateUUID()
	event.ID = id
	Events[id] = event
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(event)
}

func (c EventController) Read(id string) revel.Result {
	return c.RenderJSON(Events[id])
}

func (c EventController) Update(id string, event models.Event) revel.Result {
	Events[id] = event
	return c.RenderJSON(event)
}

func (c EventController) Delete(id string) revel.Result {
	delete(Events, id)
	return c.RenderJSON("")
}

func (c EventController) List() revel.Result {
	return c.RenderJSON(Events)
}

func (c EventController) ShowList() revel.Result {
	return c.Render(Events)
}
