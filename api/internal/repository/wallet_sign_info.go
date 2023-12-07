package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type WalletSignInfoRepository interface {
	CreateNewWalletNonce(walletAddress string, nonce string) (*WalletSignInfoModel, error)
	GetWalletNonce(walletAddress string) (*WalletSignInfoModel, error)
}

type WalletSignInfoModel struct {
	WalletAddress       string    `gorm:"primaryKey;column:wallet_address" json:"walletAddress"`
	Nonce               string    `json:"nonce"`
	NonceCreatedAt      time.Time `json:"nonceCreatedAt"`
	Signature           string    `json:"signature"`
	SignatureVerifiedAt time.Time `json:"signatureVerifiedAt"`
}

type walletSignInfoImpl struct {
	db *gorm.DB
}

func (WalletSignInfoModel) TableName() string {
	return "wallet_sign_info"
}

func NewWalletSignInfoDbImpl(db *gorm.DB) WalletSignInfoRepository {
	return &walletSignInfoImpl{
		db: db,
	}
}

func (w *walletSignInfoImpl) CreateNewWalletNonce(walletAddress string, nonce string) (*WalletSignInfoModel, error) {
	signInfo := &WalletSignInfoModel{
		WalletAddress:  walletAddress,
		Nonce:          nonce,
		NonceCreatedAt: time.Now(),
	}
	r := w.db.Create(signInfo)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to insert into db: %v", r.Error)
	}
	return signInfo, nil
}

func (w *walletSignInfoImpl) GetWalletNonce(walletAddress string) (*WalletSignInfoModel, error) {
	result := &WalletSignInfoModel{}
	r := w.db.Where("wallet_address = ?", walletAddress).First(result)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return result, nil
}

func (w *walletSignInfoImpl) UpdateWalletSignature(walletAddress string, signature string) error {
	r := w.db.Model(&WalletSignInfoModel{}).Where("wallet_address = ?", walletAddress).Updates(&WalletSignInfoModel{
		Signature:           signature,
		SignatureVerifiedAt: time.Now(),
	})
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return fmt.Errorf("failed to update db: %v", r.Error)
	}

	return nil
}
