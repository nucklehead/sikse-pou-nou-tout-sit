package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models/mongodb"
	"time"
)

type Account struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"password" bson:"password"`
	Email     string        `json:"email" bson:"email"`
	Phone     string        `json:"phone" bson:"phone"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func newAccountCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("accounts")
}

// AddAccount insert a new Account into database and returns
// last inserted account on success.
func AddAccount(m Account) (account Account, err error) {
	c := newAccountCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdateAccount update a Account into database and returns
// last nil on success.
func (m Account) UpdateAccount() error {
	c := newAccountCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"username": m.Username, "password": m.Password, "email": m.Email, "phone": m.Phone, "updatedAt": time.Now()},
	})
	return err
}

// DeleteAccount Delete Account from database and returns
// last nil on success.
func (m Account) DeleteAccount() error {
	c := newAccountCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetAccounts Get all Account from database and returns
// list of Account on success
func GetAccounts() ([]Account, error) {
	var (
		accounts []Account
		err      error
	)

	c := newAccountCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&accounts)
	return accounts, err
}

// GetAccount Get a Account from database and returns
// a Account on success
func GetAccount(id bson.ObjectId) (Account, error) {
	var (
		account Account
		err     error
	)

	c := newAccountCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&account)
	return account, err
}
