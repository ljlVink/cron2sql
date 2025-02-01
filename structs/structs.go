package structs

import (
	"time"
)

type TaskRequest struct {
	ID          uint   `json:"id"`      // 任务 ID
	Option      string `json:"option"`  // 操作类型：create/get/delete
	CronJob     string `json:"cronjob"` // cron time
	Description string `json:"desc"`    // 任务描述
	Command     string `json:"command"` // 命令
}

type ListTaskResponse struct {
	Total int         `json:"total"` // task总数
	Tasks []Task_json `json:"tasks"`
}

type Task_json struct {
	ID          uint      `json:"id"`      // id
	CronJob     string    `json:"cronjob"` // Cron 表达式
	Command     string    `json:"command"` // 任务命令
	Description string    `json:"desc"`    // 任务描述
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间
}

type Task struct {
	ID          uint      `gorm:"primaryKey"` // 主键
	CronJob     string    `gorm:"size:50"`    // Cron 表达式
	Command     string    `gorm:"size:255"`   // 任务命令
	Description string    `gorm:"size:255"`   // 任务描述
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间
}
