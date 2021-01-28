package dynamo

// snippet-start:[dynamodb.go.list_tables.imports]
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"fmt"
)

// snippet-end:[dynamodb.go.list_tables.imports]

// GetItemsByTableName retrieves all items by table name
func GetItemsByTableName() {
	// snippet-start:[dynamodb.go.list_tables.session]
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	// snippet-end:[dynamodb.go.list_tables.session]

	tableName := "First-Table"
	// userName := "Bryam Pacheco"

	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
