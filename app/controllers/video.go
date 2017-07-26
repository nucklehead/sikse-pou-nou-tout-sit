package controllers

import (
	"github.com/revel/revel"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"crypto/rand"
	"net/http"
	uuid "github.com/hashicorp/go-uuid"

)

var Videos map[string]models.Video

type VideoController struct {
	*revel.Controller
}

func (c VideoController) Create(video models.Video) revel.Result {
	id, _ := uuid.GenerateUUID()
	Videos[id] = video
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(video)
}


func (c VideoController) Read(videoID string) revel.Result {
	return c.RenderJSON(Videos[videoID])
}

func (c VideoController) Update(video models.Video) revel.Result {
	Videos[video.ID] = video
	return c.RenderJSON(video)
}

func (c VideoController) Delete(videoID string) revel.Result {
	delete(Videos, videoID)
	return c.RenderJSON("")
}

func (c VideoController) List() revel.Result {
	return c.RenderJSON(Videos)
}

func (c VideoController) ShowList() revel.Result {
	return c.Render(Videos)
}
