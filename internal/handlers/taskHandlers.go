package handlers

import (
	"golang.org/x/net/context"
	"petproject/internal/taskService"
	"petproject/internal/web/tasks"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

func (h *TaskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := uint(request.Id)
	taskUpdates := request.Body

	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return tasks.PatchTasksId404Response{}, nil
	}

	var existingTask *taskService.Task
	for _, t := range allTasks {
		if t.ID == id {
			existingTask = &t
			break
		}
	}

	if existingTask == nil {
		return tasks.PatchTasksId404Response{}, nil
	}

	if taskUpdates.Task != nil {
		existingTask.Task = *taskUpdates.Task
	}

	if taskUpdates.IsDone != nil {
		existingTask.IsDone = *taskUpdates.IsDone
	}

	updatedTask, err := h.Service.UpdateTaskByID(id, *existingTask)
	if err != nil {
		return tasks.PatchTasksId404Response{}, nil
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := uint(request.Id)

	err := h.Service.DeleteTaskByID(id)
	if err != nil {
		return tasks.DeleteTasksId404Response{}, nil
	}

	return tasks.DeleteTasksId204Response{}, nil
}
