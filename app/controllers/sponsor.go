package controllers

import (
	"github.com/revel/revel"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"crypto/rand"
	"net/http"
	uuid "github.com/hashicorp/go-uuid"

)

var Sponsors map[string]models.Sponsor

type SponsorController struct {
	*revel.Controller
}

func (c SponsorController) Create(sponsor models.Sponsor) revel.Result {
	id, _ := uuid.GenerateUUID()
	Sponsors[id] = sponsor
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(sponsor)
}


func (c SponsorController) Read(sponsorID string) revel.Result {
	return c.RenderJSON(Sponsors[sponsorID])
}

func (c SponsorController) Update(sponsor models.Sponsor) revel.Result {
	Sponsors[sponsor.ID] = sponsor
	return c.RenderJSON(sponsor)
}

func (c SponsorController) Delete(sponsorID string) revel.Result {
	delete(Sponsors, sponsorID)
	return c.RenderJSON("")
}

func (c SponsorController) List() revel.Result {
	return c.RenderJSON(Sponsors)
}

func (c SponsorController) ShowList() revel.Result {
	return c.Render(Sponsors)
}
