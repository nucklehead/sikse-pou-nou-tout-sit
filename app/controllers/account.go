package controllers

import (
	"github.com/revel/revel"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"crypto/rand"
	uuid "github.com/hashicorp/go-uuid"
)

var Accounts []models.Account

type AccountController struct {
	*revel.Controller
}

func (c AccountController) Create(account models.Account) revel.Result {
    id, _ := uuid.GenerateUUID()
    account.ID = id
    Accounts = append(Accounts, account)
    account.Password = ""
	return c.RenderJSON(account)
}

func (c AccountController) Login(account models.Account) revel.Result {
    result := make(map[string]interface{})
    result["loggedIn"] = false
    for _, savedAccount := range Accounts {
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
