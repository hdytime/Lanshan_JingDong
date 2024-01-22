package main

import (
	"Lanshan_JingDong/api/boot/initialize"
	"Lanshan_JingDong/api/boot/router"
)

func main() {
	initialize.SetupViper()
	initialize.SetupLogger()
	initialize.SetupDatabase()
	router.InitRouter()
}
