package controllers

import (
	"github.com/revel/revel"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"crypto/rand"
	"net/http"
	uuid "github.com/hashicorp/go-uuid"

)

var Options map[string]models.Option

type OptionController struct {
	*revel.Controller
}

func (c OptionController) Create(option models.Option) revel.Result {
	id, _ := uuid.GenerateUUID()
	Options[id] = option
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(option)
}


func (c OptionController) Read(optionID string) revel.Result {
	return c.RenderJSON(Options[optionID])
}

func (c OptionController) Update(optionID string, option models.Option) revel.Result {
	Options[optionID] = option
	return c.RenderJSON(option)
}

func (c OptionController) Delete(optionID string) revel.Result {
	delete(Options, optionID)
	return c.RenderJSON("")
}

func (c OptionController) List() revel.Result {
	return c.RenderJSON(Options)
}

func (c OptionController) ShowList() revel.Result {
	return c.Render(Options)
}
