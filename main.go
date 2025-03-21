package main

import (
	"VincentLimarus/log-activity/configs"
	"VincentLimarus/log-activity/controllers/helpers"
	"VincentLimarus/log-activity/routers"
)

func init(){
	configs.LoadEnvVariables()
	configs.ConnectToDB()
	configs.ConnectToMongo()
	helpers.GenerateDummyOrders()
}

func main() {
	r := routers.RoutersConfiguration()
	r.Run(":3000")
}
