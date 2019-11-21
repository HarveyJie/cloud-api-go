package monitor
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/aws/session"
)
var CloudWatchClient *cloudwatch.CloudWatch
func InitCloudWatchClient(region string) {
	sess := session.Must(session.NewSession())
	CloudWatchClient= cloudwatch.New(sess, &aws.Config{Region: aws.String(region)})
}

