package taskService

type Task struct {
	ID     uint   `gorm:"primaryKey"`
	Task   string `gorm:"not null"`
	IsDone bool   `json:"is_done"`
}
