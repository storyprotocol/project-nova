package utils

import (
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// Some of the code is from https://goethereumbook.org/en/util-go/

// SanitizeAddress validate the address and convert it to lower case
func SanitizeAddress(address string) (string, error) {
	if !IsValidAddress(address) {
		return "", fmt.Errorf("invalid address %s", address)
	}
	return strings.ToLower(address), nil
}

// IsValidNumberString validate if it's a valid number string
func IsValidNumberString(s string) bool {
	re := regexp.MustCompile("^[0-9]+$")
	return re.MatchString(s)
}

// IsValidTransaction validate hex address
func IsValidTransaction(tx string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{72}$")
	return re.MatchString(tx)
}

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case int:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}
