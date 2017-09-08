package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models/mongodb"
	"time"
)

type Banner struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Link        string        `json:"link" bson:"link"`
	Caption     string        `json:"caption" bson:"caption"`
	Description string        `json:"description" bson:"description"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}

func newBannerCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("banners")
}

// AddBanner insert a new Banner into database and returns
// last inserted banner on success.
func AddBanner(m Banner) (banner Banner, err error) {
	c := newBannerCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdateBanner update a Banner into database and returns
// last nil on success.
func (m Banner) UpdateBanner() error {
	c := newBannerCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"link": m.Link, "caption": m.Caption, "description": m.Description, "updatedAt": time.Now()},
	})
	return err
}

// DeleteBanner Delete Banner from database and returns
// last nil on success.
func (m Banner) DeleteBanner() error {
	c := newBannerCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetBanners Get all Banner from database and returns
// list of Banner on success
func GetBanners() ([]Banner, error) {
	var (
		banners []Banner
		err     error
	)

	c := newBannerCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&banners)
	return banners, err
}

// GetBanner Get a Banner from database and returns
// a Banner on success
func GetBanner(id bson.ObjectId) (Banner, error) {
	var (
		banner Banner
		err    error
	)

	c := newBannerCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&banner)
	return banner, err
}
