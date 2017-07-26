package controllers

import (
	"github.com/revel/revel"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"crypto/rand"
	"net/http"
)

var Sponsors []models.Sponsor

type SponsorController struct {
	*revel.Controller
}

func (c SponsorController) Create(sponsor models.Sponsor) revel.Result {
	Sponsors = append(Sponsors, sponsor)
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(sponsor)
}


func (c SponsorController) Read(sponsorID string) revel.Result {
	sponsor := models.Sponsor{}
	return c.RenderJSON(sponsor)
}

func (c SponsorController) Update(sponsor models.Sponsor) revel.Result {
	return c.RenderJSON(sponsor)
}

func (c SponsorController) Delete(sponsorID string) revel.Result {
	return c.RenderJSON("")
}

func (c SponsorController) List() revel.Result {
	return c.RenderJSON(Sponsors)
}

func (c SponsorController) ShowList() revel.Result {
	return c.Render(Sponsors)
}
