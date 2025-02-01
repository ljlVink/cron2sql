/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
    log "github.com/sirupsen/logrus"

)

var rootCmd = &cobra.Command{
	Use:   "cron2sql",
	Short: "Cron jobs to sql",
	Long: `A simple example of reading sql databases to apply tasks to cron. With the ability to add/delete/modify jobs`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initlogs()
	rootCmd.AddCommand(serverCmd)
}

func initlogs(){
	log.SetFormatter(&log.TextFormatter{
        FullTimestamp: true,          // 显示完整时间戳
		TimestampFormat: "2006-01-02 15:04:05-07:00", // 自定义时间戳格式
        ForceColors:   true,          // 强制启用颜色
        DisableColors: false,         // 不禁用颜色
    })
	log.Info("logrus init success")
}
