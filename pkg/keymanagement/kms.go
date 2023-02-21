package keymanagement

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

type KeyManagementClient interface {
	Encrypt(message []byte) ([]byte, error)
	Decrypt(encryptedBytes []byte) ([]byte, error)
}

func NewKmsClient(region string) KeyManagementClient {
	sess := session.Must(session.NewSession())
	kms := kms.New(sess, aws.NewConfig().WithRegion(region))
	return &kmsClient{
		kms: kms,
	}
}

type kmsClient struct {
	kms *kms.KMS
}

func (k *kmsClient) Encrypt(message []byte) ([]byte, error) {
	output, err := k.kms.Encrypt(&kms.EncryptInput{Plaintext: message})
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt the message from kms %v", err)
	}

	return output.CiphertextBlob, nil
}

func (k *kmsClient) Decrypt(encryptedBytes []byte) ([]byte, error) {
	output, err := k.kms.Decrypt(&kms.DecryptInput{CiphertextBlob: encryptedBytes})
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt the message from kms: %v", err)
	}

	return output.Plaintext, nil
}
