package handlers

import (
	"golang.org/x/net/context"
	"petproject/internal/userService"
	"petproject/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allTasks, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, tsk := range allTasks {
		user := users.User{
			Id:       &tsk.ID,
			Email:    &tsk.Email,
			Password: &tsk.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := h.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func (h *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := uint(request.Id)
	userUpdates := request.Body

	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return users.PatchUsersId404Response{}, nil
	}

	var existingUser *userService.User
	for _, t := range allUsers {
		if t.ID == id {
			existingUser = &t
			break
		}
	}

	if existingUser == nil {
		return users.PatchUsersId404Response{}, nil
	}

	if userUpdates.Email != nil {
		existingUser.Email = *userUpdates.Email
	}

	if userUpdates.Password != nil {
		existingUser.Password = *userUpdates.Password
	}

	updatedUser, err := h.Service.UpdateUserByID(id, *existingUser)
	if err != nil {
		return users.PatchUsersId404Response{}, nil
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return response, nil
}

func (h *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := uint(request.Id)

	err := h.Service.DeleteUserByID(id)
	if err != nil {
		return users.DeleteUsersId404Response{}, nil
	}

	return users.DeleteUsersId204Response{}, nil
}
