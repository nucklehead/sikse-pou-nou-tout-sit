package controllers

import (
	"encoding/json"
	"errors"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
)

type SponsorController struct {
	*revel.Controller
}

func (c SponsorController) Index() revel.Result {
	var (
		sponsors []models.Sponsor
		err      error
	)
	sponsors, err = models.GetSponsors()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(sponsors)
}

func (c SponsorController) Show(id string) revel.Result {
	var (
		sponsor   models.Sponsor
		err       error
		sponsorID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid sponsor id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	sponsorID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid sponsor id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	sponsor, err = models.GetSponsor(sponsorID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(sponsor)
}

func (c SponsorController) Create() revel.Result {
	var (
		sponsor models.Sponsor
		err     error
	)

	err = json.NewDecoder(c.Request.Body).Decode(&sponsor)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	sponsor, err = models.AddSponsor(sponsor)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(sponsor)
}

func (c SponsorController) Update() revel.Result {
	var (
		sponsor models.Sponsor
		err     error
	)
	err = json.NewDecoder(c.Request.Body).Decode(&sponsor)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = sponsor.UpdateSponsor()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(sponsor)
}

func (c SponsorController) Delete(id string) revel.Result {
	var (
		err       error
		sponsor   models.Sponsor
		sponsorID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid sponsor id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	sponsorID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid sponsor id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	sponsor, err = models.GetSponsor(sponsorID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = sponsor.DeleteSponsor()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}

func (c SponsorController) ShowList() revel.Result {
	sponsors, err := models.GetSponsors()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.Render(sponsors)
}