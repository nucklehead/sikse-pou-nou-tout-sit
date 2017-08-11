package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models/mongodb"
	"time"
)

type Sponsor struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}

func newSponsorCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("sponsors")
}

// AddSponsor insert a new Sponsor into database and returns
// last inserted sponsor on success.
func AddSponsor(m Sponsor) (sponsor Sponsor, err error) {
	c := newSponsorCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdateSponsor update a Sponsor into database and returns
// last nil on success.
func (m Sponsor) UpdateSponsor() error {
	c := newSponsorCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"name": m.Name, "description": m.Description, "updatedAt": time.Now()},
	})
	return err
}

// DeleteSponsor Delete Sponsor from database and returns
// last nil on success.
func (m Sponsor) DeleteSponsor() error {
	c := newSponsorCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetSponsors Get all Sponsor from database and returns
// list of Sponsor on success
func GetSponsors() ([]Sponsor, error) {
	var (
		sponsors []Sponsor
		err      error
	)

	c := newSponsorCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&sponsors)
	return sponsors, err
}

// GetSponsor Get a Sponsor from database and returns
// a Sponsor on success
func GetSponsor(id bson.ObjectId) (Sponsor, error) {
	var (
		sponsor Sponsor
		err     error
	)

	c := newSponsorCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&sponsor)
	return sponsor, err
}
