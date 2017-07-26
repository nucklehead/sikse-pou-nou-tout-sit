package models

import (
	"time"
	uuid "github.com/hashicorp/go-uuid"

)

type DBElement struct {
	ID      string
}

type Account struct {
	DBElement
	Username      string
	Password      string
	Email         string
	Phone         string
}
type Option struct {
	DBElement
	Name      string
	Description      string
}

type Comment struct {
	DBElement
	User      string
	Content      string
}

type Video struct {
	DBElement
	Link      string
	Title      string
	Description      string
}

type Sponsor struct {
	DBElement
	Name      string
	Description      string
}

type Event struct {
	DBElement
	Title      string
	Date      time.Time
	Description      string
	Speaker      string
	Location      time.Location
}

func newAccount(username, password, email, phone string) Account {
	id, _ := uuid.GenerateUUID()
	return Account{DBElement{id}, username, password, email, phone}
}

func newOption(name, description string) Option {
	id, _ := uuid.GenerateUUID()
	return Option{DBElement{id},name, description}
}

func newComment(user, content string) Comment {
	id, _ := uuid.GenerateUUID()
	return Comment{DBElement{id},user, content}
}

func newVideo(link, title, description string) Video {
	id, _ := uuid.GenerateUUID()
	return Video{DBElement{id},link, title, description}
}

func newSponsor(name, description string) Sponsor {
	id, _ := uuid.GenerateUUID()
	return Sponsor{DBElement{id},name, description}
}

func newEvent(title, description, speaker string, date time.Time, location time.Location) Event {
	id, _ := uuid.GenerateUUID()
	return Event{DBElement{id},title, date, description, speaker, location}
}