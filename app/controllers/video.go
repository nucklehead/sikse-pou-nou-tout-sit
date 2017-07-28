package controllers

import (
	"github.com/revel/revel"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"net/http"
	uuid "github.com/hashicorp/go-uuid"

)

var Videos map[string]models.Video

type VideoController struct {
	*revel.Controller
}

func (c VideoController) Create(video models.Video) revel.Result {
	id, _ := uuid.GenerateUUID()
	video.ID = id
	Videos[id] = video
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(video)
}


func (c VideoController) Read(id string) revel.Result {
	return c.RenderJSON(Videos[id])
}

func (c VideoController) Update(id string, video models.Video) revel.Result {
	Videos[id] = video
	return c.RenderJSON(video)
}

func (c VideoController) Delete(id string) revel.Result {
	delete(Videos, id)
	return c.RenderJSON("")
}

func (c VideoController) List() revel.Result {
	return c.RenderJSON(Videos)
}

func (c VideoController) ShowList() revel.Result {
	return c.Render(Videos)
}
