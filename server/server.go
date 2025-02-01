package server

import (
	"context"
	"cron2sql/structs"
	"cron2sql/task"
	"cron2sql/utils"
    "cron2sql/cronmanager"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init(){
	gin.SetMode(gin.ReleaseMode)
}

func Startserver(taskService task.TaskService,cronManger *cronmanager.CronManager){
	r := gin.Default()
	r.POST("/task", func(c *gin.Context) {
        var req structs.TaskRequest

        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
            return
        }
        switch req.Option {
        case "create":
			log.Println("receive task create,id=" , req.Description, ",Command=",req.Command,"cronjob=",req.CronJob)
            newTask := &structs.Task{
                CronJob:     req.CronJob, // 默认每小时执行一次
                Command:     req.Command,
                Description: req.Description,
                CreatedAt:   time.Now(),
                UpdatedAt:   time.Now(),
            }
            err:=taskService.CreateTask(context.Background(),newTask)
            utils.CheckErr(err,"fatal error on Creating Task!!")
            cronManger.AddTask(*newTask) // 同步添加任务

            c.JSON(http.StatusOK, gin.H{
                "message": "Task created",
                "command": req.Command,
                "code" : 0,
            })
        case "get":
			log.Println("receive task get")
            tasks, err := taskService.GetAllTasks(context.Background())
            if err != nil {
                log.Println("Failed to list tasks:", err)
                return
            }
            taskList := make([]structs.Task_json, 0, len(tasks))

            for _, t := range tasks {
                taskList = append(taskList, structs.Task_json{
                    ID:          t.ID,
                    CronJob:     t.CronJob,
                    Command:     t.Command,
                    Description: t.Description,
                    CreatedAt:   t.CreatedAt,
                    UpdatedAt:   t.UpdatedAt,
                })
            }
            response := structs.ListTaskResponse{
                Total: len(taskList),
                Tasks: taskList,
            }
            c.JSON(http.StatusOK, response)
        case "delete":
			if req.ID==0{
				log.Println("task delete id==0 not allowed")
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "not allowed",
                    "code" : -1,
				})
				return
			}
            err:=taskService.DeleteTask(context.Background(),req.ID)
            cronManger.RemoveTask(req.ID)
            if err!=nil{
                c.JSON(http.StatusInternalServerError, gin.H{
                    "message": "Failed delete task",
                    "code" : -1,
                })    
            }
            c.JSON(http.StatusOK, gin.H{
                "message": "Deleted task",
                "code" : 0,
            })    


		default:
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid option"})
        }
    })
	r.Run(":8080")
}