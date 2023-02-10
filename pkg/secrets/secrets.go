package secrets

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	yaml "gopkg.in/yaml.v2"
)

// FetchSecrets loads service secrets from aws secrets manager
func FetchSecrets(region string, appID string, config interface{}) error {
	//Create a Secrets Manager client
	session, err := session.NewSession()
	if err != nil {
		return err
	}
	svc := secretsmanager.New(session,
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(appID),
	}

	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
	result, err := svc.GetSecretValue(input)
	if err != nil {
		return err
	}

	if result.SecretString != nil {
		secretString := *result.SecretString
		if err := yaml.Unmarshal([]byte(secretString), config); err != nil {
			// if secrets can't be serialized
			return err
		}
	}

	return nil
}
