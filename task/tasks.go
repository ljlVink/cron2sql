package task

import (
    "context"
    "errors"
    "gorm.io/gorm"
    "cron2sql/structs"
)

// Task 定义任务表结构

// TaskService 定义任务服务的接口
type TaskService interface {
    CreateTask(ctx context.Context, task *structs.Task) error
    GetTaskByID(ctx context.Context, id uint) (*structs.Task, error)
    GetAllTasks(ctx context.Context) ([]structs.Task, error)
    UpdateTask(ctx context.Context, task *structs.Task) error
    DeleteTask(ctx context.Context, id uint) error
}

// TaskServiceImpl 实现 TaskService 接口
type TaskServiceImpl struct {
    db *gorm.DB
}

// NewTaskService 创建一个新的任务服务
func NewTaskService(db *gorm.DB) TaskService {
    return &TaskServiceImpl{db: db}
}

// CreateTask 创建任务
func (s *TaskServiceImpl) CreateTask(ctx context.Context, task *structs.Task) error {
    return s.db.WithContext(ctx).Create(task).Error
}

// GetTaskByID 根据 ID 查询任务
func (s *TaskServiceImpl) GetTaskByID(ctx context.Context, id uint) (*structs.Task, error) {
    var task structs.Task
    err := s.db.WithContext(ctx).First(&task, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil // 任务不存在
    }
    return &task, err
}

// GetAllTasks 查询所有任务
func (s *TaskServiceImpl) GetAllTasks(ctx context.Context) ([]structs.Task, error) {
    var tasks []structs.Task
    err := s.db.WithContext(ctx).Find(&tasks).Error
    return tasks, err
}

// UpdateTask 更新任务
func (s *TaskServiceImpl) UpdateTask(ctx context.Context, task *structs.Task) error {
    return s.db.WithContext(ctx).Save(task).Error
}

// DeleteTask 删除任务
func (s *TaskServiceImpl) DeleteTask(ctx context.Context, id uint) error {
    return s.db.WithContext(ctx).Delete(&structs.Task{}, id).Error
}