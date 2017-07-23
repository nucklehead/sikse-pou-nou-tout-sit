package models

import (

)

type Account struct {
	Username      string
	Password      string
	Email         string
	Phone         string

}

func newAccount(username, password, email, phone string) Account {
	return Account{username, password, email, phone}
}