package controllers

import (
	"github.com/revel/revel"
	"github.com/nucklehead/sikse-pou-nou-tout-sit/app/models"
	"net/http"
	uuid "github.com/hashicorp/go-uuid"
)

var Comments = map[string]models.Comment{}

type CommentController struct {
	*revel.Controller
}

func (c CommentController) Create(comment models.Comment) revel.Result {
	id, _ := uuid.GenerateUUID()
	comment.ID = id
	Comments[id] = comment
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(comment)
}


func (c CommentController) Read(id string) revel.Result {
	return c.RenderJSON(Comments[id])
}

func (c CommentController) Update(id string, comment models.Comment) revel.Result {
	Comments[id] = comment
	return c.RenderJSON(comment)
}

func (c CommentController) Delete(id string) revel.Result {
	delete(Comments, id)
	return c.RenderJSON("")
}

func (c CommentController) List() revel.Result {
	return c.RenderJSON(Comments)
}

func (c CommentController) ShowList() revel.Result {
	return c.Render(Comments)
}
