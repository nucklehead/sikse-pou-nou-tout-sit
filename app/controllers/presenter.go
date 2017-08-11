package controllers

import (
	"encoding/json"
	"errors"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
)

type PresenterController struct {
	*revel.Controller
}

func (c PresenterController) Index() revel.Result {
	var (
		presenters []models.Presenter
		err        error
	)
	presenters, err = models.GetPresenters()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(presenters)
}

func (c PresenterController) Show(id string) revel.Result {
	var (
		presenter   models.Presenter
		err         error
		presenterID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid presenter id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	presenterID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid presenter id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	presenter, err = models.GetPresenter(presenterID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(presenter)
}

func (c PresenterController) Create() revel.Result {
	var (
		presenter models.Presenter
		err       error
	)

	err = json.NewDecoder(c.Request.Body).Decode(&presenter)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	presenter, err = models.AddPresenter(presenter)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(presenter)
}

func (c PresenterController) Update() revel.Result {
	var (
		presenter models.Presenter
		err       error
	)
	err = json.NewDecoder(c.Request.Body).Decode(&presenter)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = presenter.UpdatePresenter()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(presenter)
}

func (c PresenterController) Delete(id string) revel.Result {
	var (
		err         error
		presenter   models.Presenter
		presenterID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid presenter id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	presenterID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid presenter id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	presenter, err = models.GetPresenter(presenterID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = presenter.DeletePresenter()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}

func (c PresenterController) ShowList() revel.Result {
	presenters, err := models.GetPresenters()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.Render(presenters)
}