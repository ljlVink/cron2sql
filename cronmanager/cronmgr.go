package cronmanager

import (
    "context"
    log "github.com/sirupsen/logrus"
    "cron2sql/task"
    "cron2sql/structs"
    "github.com/robfig/cron/v3"
)

// CronManager 管理 Cron 作业
type CronManager struct {
    cron        *cron.Cron
    taskService task.TaskService
    entries     map[uint]cron.EntryID // 记录任务 ID 和 Cron 条目 ID 的映射
}


// NewCronManager 创建一个新的 CronManager
func NewCronManager(taskService task.TaskService) *CronManager {
    return &CronManager{
        cron:        cron.New(cron.WithSeconds()),
        taskService: taskService,
        entries:     make(map[uint]cron.EntryID), // 初始化映射
    }
}


// Start 启动 CronManager
func (m *CronManager) Start() {
    m.entries = make(map[uint]cron.EntryID)
    
    // 加载所有任务并添加到 Cron 调度器
    tasks, err := m.taskService.GetAllTasks(context.Background())
    if err != nil {
        log.Fatalf("Failed to load tasks: %v", err)
    }

    for _, t := range tasks {
        m.AddTask(t)
    }

    // 启动 Cron 调度器
    m.cron.Start()
    log.Println("CronManager started")
}


// Stop 停止 CronManager
func (m *CronManager) Stop() {
    m.cron.Stop()
    log.Println("CronManager stopped")
}

// AddTask 添加任务到 Cron 调度器
func (m *CronManager) AddTask(t structs.Task) {
    entryID, err := m.cron.AddFunc(t.CronJob, func() {
        log.Printf("Executing task: %s (ID: %d)\n", t.Description, t.ID)
        log.Printf("Command: %s\n", t.Command)
    })
    if err != nil {
        log.Printf("Failed to add task %d to cron: %v", t.ID, err)
    } else {
        m.entries[t.ID] = entryID // 记录任务 ID 和 Cron 条目 ID
        log.Printf("Task id:%d added to cron scheduler", t.ID)
    }
}


// RemoveTask 从 Cron 调度器中移除任务
func (m *CronManager) RemoveTask(id uint) {
    if entryID, exists := m.entries[id]; exists {
        m.cron.Remove(entryID) // 通过 EntryID 移除任务
        delete(m.entries, id)  // 从映射中删除记录
        log.Printf("Task id:%d removed from cron scheduler", id)
    } else {
        log.Printf("Task id:%d not found in cron scheduler", id)
    }
}
