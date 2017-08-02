package controllers

import (
	"net/http"

	uuid "github.com/hashicorp/go-uuid"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"github.com/revel/revel"
)

var Presenters map[string]models.Presenter

type PresenterController struct {
	*revel.Controller
}

func (c PresenterController) Create(presenter models.Presenter) revel.Result {
	id, _ := uuid.GenerateUUID()
	presenter.ID = id
	Presenters[id] = presenter
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(presenter)
}

func (c PresenterController) Read(id string) revel.Result {
	return c.RenderJSON(Presenters[id])
}

func (c PresenterController) Update(id string, presenter models.Presenter) revel.Result {
	Presenters[id] = presenter
	return c.RenderJSON(presenter)
}

func (c PresenterController) Delete(id string) revel.Result {
	delete(Presenters, id)
	return c.RenderJSON("")
}

func (c PresenterController) List() revel.Result {
	return c.RenderJSON(Presenters)
}

func (c PresenterController) ShowList() revel.Result {
	return c.Render(Presenters)
}
