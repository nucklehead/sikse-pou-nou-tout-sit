package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models/mongodb"
	"time"
)

type Event struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Date        time.Time     `json:"date" bson:"date"`
	Description string        `json:"description" bson:"description"`
	Speaker     string        `json:"speaker" bson:"speaker"`
	Location    string        `json:"location" bson:"location"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}

func newEventCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("events")
}

// AddEvent insert a new Event into database and returns
// last inserted event on success.
func AddEvent(m Event) (event Event, err error) {
	c := newEventCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdateEvent update a Event into database and returns
// last nil on success.
func (m Event) UpdateEvent() error {
	c := newEventCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"title": m.Title, "date": m.Date, "description": m.Description, "speaker": m.Speaker, "location": m.Location, "updatedAt": time.Now()},
	})
	return err
}

// DeleteEvent Delete Event from database and returns
// last nil on success.
func (m Event) DeleteEvent() error {
	c := newEventCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetEvents Get all Event from database and returns
// list of Event on success
func GetEvents() ([]Event, error) {
	var (
		events []Event
		err    error
	)

	c := newEventCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&events)
	return events, err
}

// GetEvent Get a Event from database and returns
// a Event on success
func GetEvent(id bson.ObjectId) (Event, error) {
	var (
		event Event
		err   error
	)

	c := newEventCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&event)
	return event, err
}
