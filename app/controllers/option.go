package controllers

import (
	"net/http"

	uuid "github.com/hashicorp/go-uuid"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"github.com/revel/revel"
)

var Options = map[string]models.Option{}

type OptionController struct {
	*revel.Controller
}

func (c OptionController) Create(option models.Option) revel.Result {
	id, _ := uuid.GenerateUUID()
	option.ID = id
	Options[id] = option
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(option)
}

func (c OptionController) Read(id string) revel.Result {
	return c.RenderJSON(Options[id])
}

func (c OptionController) Update(id string, option models.Option) revel.Result {
	Options[id] = option
	return c.RenderJSON(option)
}

func (c OptionController) Delete(id string) revel.Result {
	delete(Options, id)
	return c.RenderJSON("")
}

func (c OptionController) List() revel.Result {
	return c.RenderJSON(Options)
}

func (c OptionController) ShowList() revel.Result {
	return c.Render(Options)
}
