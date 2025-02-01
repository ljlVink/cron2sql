package utils

import(
    log "github.com/sirupsen/logrus"
)

func CheckErr(err error,errorinfo string){
	if err!=nil{
		log.Fatalln("Fatal error :",errorinfo,err)
	}
}