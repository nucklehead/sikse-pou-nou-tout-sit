package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models/mongodb"
	"time"
)

type Option struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}

func newOptionCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("options")
}

// AddOption insert a new Option into database and returns
// last inserted option on success.
func AddOption(m Option) (option Option, err error) {
	c := newOptionCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdateOption update a Option into database and returns
// last nil on success.
func (m Option) UpdateOption() error {
	c := newOptionCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"name": m.Name, "description": m.Description, "updatedAt": time.Now()},
	})
	return err
}

// DeleteOption Delete Option from database and returns
// last nil on success.
func (m Option) DeleteOption() error {
	c := newOptionCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetOptions Get all Option from database and returns
// list of Option on success
func GetOptions() ([]Option, error) {
	var (
		options []Option
		err     error
	)

	c := newOptionCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&options)
	return options, err
}

// GetOption Get a Option from database and returns
// a Option on success
func GetOption(id bson.ObjectId) (Option, error) {
	var (
		option Option
		err    error
	)

	c := newOptionCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&option)
	return option, err
}
