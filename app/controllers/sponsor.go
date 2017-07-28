package controllers

import (
	"github.com/revel/revel"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"net/http"
	uuid "github.com/hashicorp/go-uuid"

)

var Sponsors map[string]models.Sponsor

type SponsorController struct {
	*revel.Controller
}

func (c SponsorController) Create(sponsor models.Sponsor) revel.Result {
	id, _ := uuid.GenerateUUID()
	sponsor.ID = id
	Sponsors[id] = sponsor
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(sponsor)
}


func (c SponsorController) Read(id string) revel.Result {
	return c.RenderJSON(Sponsors[id])
}

func (c SponsorController) Update(id string, sponsor models.Sponsor) revel.Result {
	Sponsors[id] = sponsor
	return c.RenderJSON(sponsor)
}

func (c SponsorController) Delete(id string) revel.Result {
	delete(Sponsors, id)
	return c.RenderJSON("")
}

func (c SponsorController) List() revel.Result {
	return c.RenderJSON(Sponsors)
}

func (c SponsorController) ShowList() revel.Result {
	return c.Render(Sponsors)
}
