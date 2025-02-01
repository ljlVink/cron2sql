package cmd

import (
    log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"context"
	"time"
	"cron2sql/server"
    "cron2sql/structs"
    "cron2sql/cronmanager"
	"cron2sql/task"
	"cron2sql/database"
	"cron2sql/utils"
)

var (
    taskService task.TaskService
    cronManager *cronmanager.CronManager
)

func init(){
	db, err := database.InitDB()
	utils.CheckErr(err,"database init failed")
	taskService = task.NewTaskService(db)
    cronManager = cronmanager.NewCronManager(taskService)
    cronManager.Start()
}
var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "start serve",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("start service..")
		server.Startserver(taskService,cronManager)
		newTask := &structs.Task{
            CronJob:     "0 * * * *", // 默认每小时执行一次
            Command:     "echo 'Hello, World!'",
            Description: "Print Hello World every hour",
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        }
        err := taskService.CreateTask(context.Background(), newTask)
        if err != nil {
            log.Println("Failed to create task:", err)
            return
        }
		

	},
}
