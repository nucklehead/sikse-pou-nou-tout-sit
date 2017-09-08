package controllers

import (
	"encoding/json"
	"errors"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
)

type BannerController struct {
	*revel.Controller
}

func (c BannerController) Index() revel.Result {
	var (
		banners []models.Banner
		err     error
	)
	banners, err = models.GetBanners()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(banners)
}

func (c BannerController) Show(id string) revel.Result {
	var (
		banner   models.Banner
		err      error
		bannerID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid banner id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	bannerID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid banner id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	banner, err = models.GetBanner(bannerID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(banner)
}

func (c BannerController) Create() revel.Result {
	var (
		banner models.Banner
		err    error
	)

	err = json.NewDecoder(c.Request.Body).Decode(&banner)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	banner, err = models.AddBanner(banner)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(banner)
}

func (c BannerController) Update() revel.Result {
	var (
		banner models.Banner
		err    error
	)
	err = json.NewDecoder(c.Request.Body).Decode(&banner)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = banner.UpdateBanner()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(banner)
}

func (c BannerController) Delete(id string) revel.Result {
	var (
		err      error
		banner   models.Banner
		bannerID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid banner id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	bannerID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid banner id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	banner, err = models.GetBanner(bannerID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = banner.DeleteBanner()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}

func (c BannerController) ShowList() revel.Result {
	banners, err := models.GetBanners()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.Render(banners)
}
