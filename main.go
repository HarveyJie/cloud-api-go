package main

import (
	"github.com/HarveyJie/cloud-api-go/amazon/monitor"
	"time"
)
func main()  {

    monitor.InitCloudWatchClient("cn-north-1")

	now:=time.Now()
	later :=now.Add(-time.Minute)
	monitor.RdsConnFilterDBName(now,later,"accountdb01")


}