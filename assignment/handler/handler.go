package handler

import (
	"github.com/assignment/models"
	"gofr.dev/pkg/gofr"
	gofrhttp "gofr.dev/pkg/gofr/http"
)

type handler struct {
	service UserInterface
}

func New(userInterface UserInterface) *handler {
	return &handler{service: userInterface}
}

func (h *handler) GetUserByName(ctx *gofr.Context) (interface{}, error) {
	name := ctx.Request.PathParam("name")
	if name == "" {
		return nil, gofrhttp.ErrorMissingParam{}
	}
	var user *models.User
	user, err := h.service.GetUserByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h *handler) AddUser(ctx *gofr.Context) (interface{}, error) {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return nil, gofrhttp.ErrorMissingParam{}
	}
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	err = h.service.AddUser(ctx, &user)
	if err != nil {
		return nil, err
	}
	return nil, nil

}
func (h *handler) DeleteUser(ctx *gofr.Context) (interface{}, error) {
	name := ctx.Request.PathParam("name")
	err := h.service.DeleteUser(ctx, name)
	if name == "" {
		return nil, gofrhttp.ErrorMissingParam{}
	}
	if err != nil {
		return nil, gofrhttp.ErrorEntityNotFound{}
	}
	return nil, nil
}

func (h *handler) UpdateEmail(ctx *gofr.Context) (interface{}, error) {
	name := ctx.Request.PathParam("name")
	var user models.User

	if err := ctx.Bind(&user); err != nil {
		return nil, gofrhttp.ErrorInvalidParam{}
	}
	err := user.Validate()
	if err != nil {
		return nil, gofrhttp.ErrorInvalidParam{}
	}
	err = h.service.UpdateEmail(ctx, name, user.Email)
	if err != nil {
		return nil, gofrhttp.ErrorInvalidParam{}
	}
	return nil, nil
}
