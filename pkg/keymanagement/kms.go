package keymanagement

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

type KeyManagementClient interface {
	Encrypt(message []byte, keyId string) ([]byte, error)
	Decrypt(encryptedBytes []byte, keyId string) ([]byte, error)
}

func NewKmsClient(region string, profile string) (KeyManagementClient, error) {
	var sess *session.Session
	var err error

	if profile != "" {
		sess, err = session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Profile:           "stage",
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create new kms session with profile %s: %w", profile, err)
		}
	} else {
		sess = session.Must(session.NewSession())
	}
	kms := kms.New(sess, aws.NewConfig().WithRegion(region))
	return &kmsClient{
		kms: kms,
	}, nil
}

type kmsClient struct {
	kms *kms.KMS
}

func (k *kmsClient) Encrypt(message []byte, keyId string) ([]byte, error) {
	output, err := k.kms.Encrypt(&kms.EncryptInput{
		Plaintext: message,
		KeyId:     &keyId,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt the message from kms %v", err)
	}

	return output.CiphertextBlob, nil
}

func (k *kmsClient) Decrypt(encryptedBytes []byte, keyId string) ([]byte, error) {
	output, err := k.kms.Decrypt(&kms.DecryptInput{
		CiphertextBlob: encryptedBytes,
		KeyId:          &keyId,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt the message from kms: %v", err)
	}

	return output.Plaintext, nil
}
