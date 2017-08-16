package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models/mongodb"
	"time"
)

type Presenter struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
	Twitter   string        `json:"twitter" bson:"twitter"`
	About     string        `json:"about" bson:"about"`
	Location  string        `json:"location" bson:"location"`
	Email     string        `json:"email" bson:"email"`
	Phone     string        `json:"phone" bson:"phone"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func newPresenterCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("presenters")
}

// AddPresenter insert a new Presenter into database and returns
// last inserted presenter on success.
func AddPresenter(m Presenter) (presenter Presenter, err error) {
	c := newPresenterCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdatePresenter update a Presenter into database and returns
// last nil on success.
func (m Presenter) UpdatePresenter() error {
	c := newPresenterCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"firstName": m.FirstName, "lastName": m.LastName, "twitter": m.Twitter, "about": m.About, "location": m.Location, "email": m.Email, "phone": m.Phone, "updatedAt": time.Now()},
	})
	return err
}

// DeletePresenter Delete Presenter from database and returns
// last nil on success.
func (m Presenter) DeletePresenter() error {
	c := newPresenterCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetPresenters Get all Presenter from database and returns
// list of Presenter on success
func GetPresenters() ([]Presenter, error) {
	var (
		presenters []Presenter
		err        error
	)

	c := newPresenterCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&presenters)
	return presenters, err
}

// GetPresenter Get a Presenter from database and returns
// a Presenter on success
func GetPresenter(id bson.ObjectId) (Presenter, error) {
	var (
		presenter Presenter
		err       error
	)

	c := newPresenterCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&presenter)
	return presenter, err
}
