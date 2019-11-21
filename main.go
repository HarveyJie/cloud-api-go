package main

import "github.com/HarveyJie/cloud-api-go/amazon/resources"

func main()  {
/*
	now:=time.Now()
	later :=now.Add(-time.Minute)
	res ,_:=monitor.RdsConnFilterDBName(now,later,"accountdb01")
    fmt.Println(res)

 */
     resources.ListRedis()


}