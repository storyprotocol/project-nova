package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type WalletMerkleProofRepository interface {
	GetMerkleProof(walletAddress string, allowlistId string) (*WalletMerkleProofModel, error)
}

type WalletMerkleProofModel struct {
	ID            string `gorm:"primaryKey;column:id"`
	AllowlistId   string
	WalletAddress string
	Proof         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (WalletMerkleProofModel) TableName() string {
	return "wallet_merkle_proof"
}

func NewWalletMerkleProofDbImpl(db *gorm.DB) WalletMerkleProofRepository {
	return &walletMerkleProofDbImpl{
		db: db,
	}
}

type walletMerkleProofDbImpl struct {
	db *gorm.DB
}

func (s *walletMerkleProofDbImpl) GetMerkleProof(walletAddress string, allowlistId string) (*WalletMerkleProofModel, error) {
	result := &WalletMerkleProofModel{}
	r := s.db.Where("wallet_address = ? and allowlist_id = ?", walletAddress, allowlistId).First(&result)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return result, nil
}
