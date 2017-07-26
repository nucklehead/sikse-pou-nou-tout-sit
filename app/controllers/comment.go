package controllers

import (
	"github.com/revel/revel"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"crypto/rand"
	"net/http"
	uuid "github.com/hashicorp/go-uuid"
)

var Comments map[string]models.Comment

type CommentController struct {
	*revel.Controller
}

func (c CommentController) Create(comment models.Comment) revel.Result {
	id, _ := uuid.GenerateUUID()
	Comments[id] = comment
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(comment)
}


func (c CommentController) Read(commentID string) revel.Result {
	return c.RenderJSON(Comments[commentID])
}

func (c CommentController) Update(comment models.Comment) revel.Result {
	Comments[comment.ID] = comment
	return c.RenderJSON(comment)
}

func (c CommentController) Delete(commentID string) revel.Result {
	delete(Comments, commentID)
	return c.RenderJSON("")
}

func (c CommentController) List() revel.Result {
	return c.RenderJSON(Comments)
}

func (c CommentController) ShowList() revel.Result {
	return c.Render(Comments)
}
