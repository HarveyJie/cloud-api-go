package resources

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
)

func ListRedis() (result *elasticache.DescribeCacheClustersOutput) {
	sess := session.Must(session.NewSession())
	svc :=elasticache.New(sess,&aws.Config{Region: aws.String("cn-north-1")})
	result,err:=svc.DescribeCacheClusters(&elasticache.DescribeCacheClustersInput{})
	if err !=nil{
		fmt.Println(err)
	}
	//fmt.Println(result)
	return result
}