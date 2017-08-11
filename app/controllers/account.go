package controllers

import (
	"encoding/json"
	"errors"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"crypto/rand"
)

type AccountController struct {
	*revel.Controller
}

func (c AccountController) Login(account models.Account) revel.Result {
	result := make(map[string]interface{})
	result["loggedIn"] = false
	accounts, err := models.GetAccounts()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	for _, savedAccount := range accounts {
		if savedAccount.Username == account.Username && savedAccount.Password == account.Password {
			result["loggedIn"] = true
			b := make([]byte, 8)
			rand.Read(b)
			result["token"] = b
			result["id"] = savedAccount.ID
		}
	}
	return c.RenderJSON(result)
}

func (c AccountController) Index() revel.Result {
	var (
		accounts []models.Account
		err      error
	)
	accounts, err = models.GetAccounts()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(accounts)
}

func (c AccountController) Show(id string) revel.Result {
	var (
		account   models.Account
		err       error
		accountID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid account id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	accountID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid account id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	account, err = models.GetAccount(accountID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(account)
}

func (c AccountController) Create() revel.Result {
	var (
		account models.Account
		err     error
	)

	err = json.NewDecoder(c.Request.Body).Decode(&account)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	account, err = models.AddAccount(account)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(account)
}

func (c AccountController) Update() revel.Result {
	var (
		account models.Account
		err     error
	)
	err = json.NewDecoder(c.Request.Body).Decode(&account)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = account.UpdateAccount()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(account)
}

func (c AccountController) Delete(id string) revel.Result {
	var (
		err       error
		account   models.Account
		accountID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid account id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	accountID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid account id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	account, err = models.GetAccount(accountID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = account.DeleteAccount()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}

func (c AccountController) ShowList() revel.Result {
	accounts, err := models.GetAccounts()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.Render(accounts)
}
