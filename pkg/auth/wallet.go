package auth

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func RecoverAddress(message string, signature string) (string, error) {
	decodedSignature, err := hexutil.Decode(signature)
	if err != nil {
		return "", err
	}
	// Support both formats of recovery bit (27/28 or 0/1)
	if decodedSignature[64] == 27 || decodedSignature[64] == 28 {
		decodedSignature[64] -= 27
	}

	messageByte := []byte(message)
	messageHash := crypto.Keccak256Hash(messageByte)
	publicKey, err := crypto.SigToPub(messageHash.Bytes(), decodedSignature)
	if err != nil {
		return "", err
	}

	address := crypto.PubkeyToAddress(*publicKey)
	return address.String(), nil
}

// For Testing
func SignMessage(message string) ([]byte, error) {
	messageByte := []byte(message)
	messageHash := crypto.Keccak256Hash(messageByte)

	privateKey, err := crypto.HexToECDSA("add your private key")
	if err != nil {
		return nil, err
	}
	signatureHash, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}

	return signatureHash, nil
}
