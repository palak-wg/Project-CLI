package main

import (
	"doctor-patient-cli/startApplication"
	"doctor-patient-cli/utils"
)

func main() {
	go utils.InitDB()
	defer utils.CloseDB()
	startApplication.StartApp()
}
