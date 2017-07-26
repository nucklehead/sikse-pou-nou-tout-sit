package controllers

import (
	"github.com/revel/revel"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"crypto/rand"
	"net/http"
)

var Options []models.Option

type OptionController struct {
	*revel.Controller
}

func (c OptionController) Create(option models.Option) revel.Result {
	Options = append(Options, option)
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(option)
}


func (c OptionController) Read(optionID string) revel.Result {
	option := models.Option{}
	return c.RenderJSON(option)
}

func (c OptionController) Update(option models.Option) revel.Result {
	return c.RenderJSON(option)
}

func (c OptionController) Delete(optionID string) revel.Result {
	return c.RenderJSON("")
}

func (c OptionController) List() revel.Result {
	return c.RenderJSON(Options)
}

func (c OptionController) ShowList() revel.Result {
	return c.Render(Options)
}
