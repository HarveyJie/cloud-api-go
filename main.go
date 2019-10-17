package main

import (
	"fmt"
	"github.com/HarveyJie/cloud-api-go/amazon/monitor"
	"time"
)

func main()  {

	now:=time.Now()
	later :=now.Add(-time.Minute)
	res ,_:=monitor.RdsConnFilterDBName(now,later,"accountdb01")
    fmt.Println(res)



}