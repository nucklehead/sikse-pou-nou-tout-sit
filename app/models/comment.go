package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models/mongodb"
	"time"
)

type Comment struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	User      string        `json:"user" bson:"user"`
	Content   string        `json:"content" bson:"content"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func newCommentCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("comments")
}

// AddComment insert a new Comment into database and returns
// last inserted comment on success.
func AddComment(m Comment) (comment Comment, err error) {
	c := newCommentCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdateComment update a Comment into database and returns
// last nil on success.
func (m Comment) UpdateComment() error {
	c := newCommentCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"user": m.User, "content": m.Content, "updatedAt": time.Now()},
	})
	return err
}

// DeleteComment Delete Comment from database and returns
// last nil on success.
func (m Comment) DeleteComment() error {
	c := newCommentCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetComments Get all Comment from database and returns
// list of Comment on success
func GetComments() ([]Comment, error) {
	var (
		comments []Comment
		err      error
	)

	c := newCommentCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&comments)
	return comments, err
}

// GetComment Get a Comment from database and returns
// a Comment on success
func GetComment(id bson.ObjectId) (Comment, error) {
	var (
		comment Comment
		err     error
	)

	c := newCommentCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&comment)
	return comment, err
}
