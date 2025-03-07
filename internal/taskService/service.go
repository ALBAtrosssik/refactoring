package taskService

type TaskService struct {
	taskRepo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{taskRepo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.taskRepo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.taskRepo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	return s.taskRepo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.taskRepo.DeleteTaskByID(id)
}
