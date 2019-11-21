package monitor
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
)
var CloudWatchClient *cloudwatch.CloudWatch
func InitCloudWatchClient(region string) {
	fmt.Println("12344")

	sess := session.Must(session.NewSession())
	CloudWatchClient= cloudwatch.New(sess, &aws.Config{Region: aws.String(region)})
}

