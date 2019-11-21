package monitor

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"fmt"
	"time"
)

func rdsFilerNameComm(startTime time.Time,endTime time.Time,dbName string,metricName string) (*cloudwatch.GetMetricStatisticsOutput,error) {
	rdsInput := &cloudwatch.GetMetricStatisticsInput{
		MetricName: aws.String(metricName),
		Namespace:  aws.String("AWS/RDS"),
		Period:     aws.Int64(60),
		EndTime:    aws.Time(startTime),
		StartTime:  aws.Time(endTime),
		Statistics: []*string{aws.String("Average")},
		Dimensions: []*cloudwatch.Dimension{
			&cloudwatch.Dimension{Name: aws.String("DBInstanceIdentifier"), Value: aws.String(dbName)},
		},
	}
	resp, error := CloudWatchClient.GetMetricStatistics(rdsInput)
	if error != nil {
		fmt.Println(error)
	}
	return resp,error
}

func RdsCPUUsagFilterDBName(startTime time.Time,endTime time.Time,dbName string) (*cloudwatch.GetMetricStatisticsOutput,error)  {
	return rdsFilerNameComm(startTime,endTime,dbName,"CPUUtilization")


}

func RdsConnFilterDBName(startTime time.Time,endTime time.Time,dbName string) (*cloudwatch.GetMetricStatisticsOutput,error) {
	return rdsFilerNameComm(startTime,endTime,dbName,"DatabaseConnections")

}

func RdsFreeMemFilterDBName(startTime time.Time,endTime time.Time,dbName string) (*cloudwatch.GetMetricStatisticsOutput,error) {
	return rdsFilerNameComm(startTime,endTime,dbName,"FreeableMemory")
}

func RdsWIOPSFilterDBName(startTime time.Time,endTime time.Time,dbName string) (*cloudwatch.GetMetricStatisticsOutput,error) {
	return rdsFilerNameComm(startTime,endTime,dbName,"WriteIOPS")
}

func RdsRIOPSFilterDBName(startTime time.Time,endTime time.Time,dbName string) (*cloudwatch.GetMetricStatisticsOutput,error) {
	return rdsFilerNameComm(startTime,endTime,dbName,"ReadIOPS")
}

func RdsNetworkRecFilterDBName(startTime time.Time,endTime time.Time,dbName string) (*cloudwatch.GetMetricStatisticsOutput,error) {
	return rdsFilerNameComm(startTime,endTime,dbName,"NetworkReceiveThroughput")
}

func RdsNetworkTransFilterDBName(startTime time.Time,endTime time.Time,dbName string) (*cloudwatch.GetMetricStatisticsOutput,error) {
	return rdsFilerNameComm(startTime,endTime,dbName,"NetworkTransmitThroughput")
}

func RdsReadLatencyFilterDBName(startTime time.Time,endTime time.Time,dbName string) (*cloudwatch.GetMetricStatisticsOutput,error) {
	return rdsFilerNameComm(startTime,endTime,dbName,"ReadLatency")
}

func RdsWriteLatencyFilterDBName(startTime time.Time,endTime time.Time,dbName string) (*cloudwatch.GetMetricStatisticsOutput,error) {
	return rdsFilerNameComm(startTime,endTime,dbName,"WriteLatency")
}

func RdsFreeDiskFilterDBName(startTime time.Time,endTime time.Time,dbName string) (*cloudwatch.GetMetricStatisticsOutput,error) {
	return rdsFilerNameComm(startTime,endTime,dbName,"FreeStorageSpace")
}
