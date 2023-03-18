package auth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func RecoverAddress(message string, signature string) (string, error) {
	decodedSignature, err := hexutil.Decode(signature)
	if err != nil {
		return "", err
	}
	// Support both formats of recovery bit (27/28 or 0/1)
	if decodedSignature[crypto.RecoveryIDOffset] == 27 || decodedSignature[crypto.RecoveryIDOffset] == 28 {
		decodedSignature[64] -= 27
	}

	messageByte := []byte(message)
	messageHash := accounts.TextHash(messageByte) //crypto.Keccak256Hash(messageByte)
	publicKey, err := crypto.SigToPub(messageHash, decodedSignature)
	if err != nil {
		return "", err
	}

	address := crypto.PubkeyToAddress(*publicKey)
	return address.String(), nil
}

// For Testing only, pKey is from input so won't be stored anywhere
func SignMessage(message string, pKey string) (string, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	messageByte := []byte(fullMessage)
	messageHash := crypto.Keccak256Hash(messageByte)

	privateKey, err := crypto.HexToECDSA(pKey)
	if err != nil {
		return "", err
	}
	signatureHash, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(signatureHash), nil
}
