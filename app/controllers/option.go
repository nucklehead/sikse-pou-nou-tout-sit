package controllers

import (
	"encoding/json"
	"errors"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
)

type OptionController struct {
	*revel.Controller
}

func (c OptionController) Index() revel.Result {
	var (
		options []models.Option
		err     error
	)
	options, err = models.GetOptions()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(options)
}

func (c OptionController) Show(id string) revel.Result {
	var (
		option   models.Option
		err      error
		optionID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid option id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	optionID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid option id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	option, err = models.GetOption(optionID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(option)
}

func (c OptionController) Create() revel.Result {
	var (
		option models.Option
		err    error
	)

	err = json.NewDecoder(c.Request.Body).Decode(&option)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	option, err = models.AddOption(option)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(option)
}

func (c OptionController) Update() revel.Result {
	var (
		option models.Option
		err    error
	)
	err = json.NewDecoder(c.Request.Body).Decode(&option)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = option.UpdateOption()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(option)
}

func (c OptionController) Delete(id string) revel.Result {
	var (
		err      error
		option   models.Option
		optionID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid option id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	optionID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid option id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	option, err = models.GetOption(optionID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = option.DeleteOption()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}

func (c OptionController) ShowList() revel.Result {
	options, err := models.GetOptions()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.Render(options)
}