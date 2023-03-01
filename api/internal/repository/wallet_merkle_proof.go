package repository

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WalletMerkleProofRepository interface {
	GetMerkleProof(walletAddress string, allowlistId string) (*WalletMerkleProofModel, error)
	CreateMerkleProof(walletAddress string, allowlistId string, proof string) (*WalletMerkleProofModel, error)
}

type WalletMerkleProofModel struct {
	ID            string    `gorm:"primaryKey;column:id" json:"id"`
	AllowlistId   string    `json:"allowlistId"`
	WalletAddress string    `json:"walletAddress"`
	Proof         string    `json:"proof"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
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
	walletAddress = strings.ToLower(walletAddress)

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

func (s *walletMerkleProofDbImpl) CreateMerkleProof(walletAddress string, allowlistId string, proof string) (*WalletMerkleProofModel, error) {
	walletAddress = strings.ToLower(walletAddress)
	model := &WalletMerkleProofModel{
		ID:            uuid.New().String(),
		WalletAddress: walletAddress,
		AllowlistId:   allowlistId,
		Proof:         proof,
	}
	r := s.db.Create(model)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to insert into db: %v", r.Error)
	}

	return model, nil
}
