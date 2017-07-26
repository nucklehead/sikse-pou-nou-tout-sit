package controllers

import (
	"github.com/revel/revel"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"crypto/rand"
	"net/http"
)

var Videos []models.Video

type VideoController struct {
	*revel.Controller
}

func (c VideoController) Create(video models.Video) revel.Result {
	Videos = append(Videos, video)
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(video)
}


func (c VideoController) Read(videoID string) revel.Result {
	video := models.Video{}
	return c.RenderJSON(video)
}

func (c VideoController) Update(video models.Video) revel.Result {
	return c.RenderJSON(video)
}

func (c VideoController) Delete(videoID string) revel.Result {
	return c.RenderJSON("")
}

func (c VideoController) List() revel.Result {
	return c.RenderJSON(Videos)
}

func (c VideoController) ShowList() revel.Result {
	return c.Render(Videos)
}
