package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models/mongodb"
	"time"
)

type Video struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Link        string        `json:"link" bson:"link"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}

func newVideoCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("videos")
}

// AddVideo insert a new Video into database and returns
// last inserted video on success.
func AddVideo(m Video) (video Video, err error) {
	c := newVideoCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdateVideo update a Video into database and returns
// last nil on success.
func (m Video) UpdateVideo() error {
	c := newVideoCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"link": m.Link, "title": m.Title, "description": m.Description, "updatedAt": time.Now()},
	})
	return err
}

// DeleteVideo Delete Video from database and returns
// last nil on success.
func (m Video) DeleteVideo() error {
	c := newVideoCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetVideos Get all Video from database and returns
// list of Video on success
func GetVideos() ([]Video, error) {
	var (
		videos []Video
		err    error
	)

	c := newVideoCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&videos)
	return videos, err
}

// GetVideo Get a Video from database and returns
// a Video on success
func GetVideo(id bson.ObjectId) (Video, error) {
	var (
		video Video
		err   error
	)

	c := newVideoCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&video)
	return video, err
}
