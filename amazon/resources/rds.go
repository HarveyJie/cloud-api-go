package resources

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/aws/session"
)

func ListRDS() ( *rds.DescribeDBInstancesOutput, error) {
	sess := session.Must(session.NewSession())
	svc :=rds.New(sess,&aws.Config{Region: aws.String("cn-north-1")})
	result,err:=svc.DescribeDBInstances(&rds.DescribeDBInstancesInput{})
	if err !=nil{
		fmt.Println(err)
	}

    return  result,err

}