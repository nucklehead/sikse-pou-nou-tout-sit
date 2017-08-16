package controllers

import (
	"encoding/json"
	"errors"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
)

type VideoController struct {
	*revel.Controller
}

func (c VideoController) Index() revel.Result {
	var (
		videos []models.Video
		err    error
	)
	videos, err = models.GetVideos()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(videos)
}

func (c VideoController) Show(id string) revel.Result {
	var (
		video   models.Video
		err     error
		videoID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid video id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	videoID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid video id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	video, err = models.GetVideo(videoID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(video)
}

func (c VideoController) Create() revel.Result {
	var (
		video models.Video
		err   error
	)

	err = json.NewDecoder(c.Request.Body).Decode(&video)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	video, err = models.AddVideo(video)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(video)
}

func (c VideoController) Update() revel.Result {
	var (
		video models.Video
		err   error
	)
	err = json.NewDecoder(c.Request.Body).Decode(&video)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = video.UpdateVideo()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(video)
}

func (c VideoController) Delete(id string) revel.Result {
	var (
		err     error
		video   models.Video
		videoID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid video id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	videoID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid video id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	video, err = models.GetVideo(videoID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = video.DeleteVideo()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}

func (c VideoController) ShowList() revel.Result {
	videos, err := models.GetVideos()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.Render(videos)
}