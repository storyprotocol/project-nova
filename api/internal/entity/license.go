package entity

import (
	"fmt"
	"strconv"
)

type License struct {
	ID                int64  `json:"id,omitempty"`
	FranchiseId       int64  `json:"franchiseId,omitempty"`
	IpAssetId         int64  `json:"ipAssetId,omitempty"`
	IpAssetName       string `json:"ipAssetName,omitempty"`
	OwnerAddress      string `json:"ownerAddress,omitempty"`
	CollectionAddress string `json:"collectionAddress,omitempty"`
	Scope             string `json:"scope,omitempty"`
	Duration          string `json:"duration,omitempty"`
	Right             string `json:"right,omitempty"`
	ImageUrl          string `json:"imageUrl,omitempty"`
}

type LicensesTheGraphResponse struct {
	LicensesGranted []*LicenseTheGraph `json:"licenseGranteds"`
}

type LicenseTheGraph struct {
	ID                string `json:"id"`
	LicenseId         string `json:"licenseId"`
	FranchiseId       string `json:"franchiseId"`
	IpAssetId         string `json:"storyId"`
	IpAssetName       string `json:"storyName"`
	OwnerAddress      string `json:"owner"`
	CollectionAddress string `json:"collectionAddress"`
	Scope             string `json:"scope"`
	Duration          string `json:"duration"`
	Right             string `json:"right"`
	ImageUrl          string `json:"imageUrl"`
}

func (l *LicenseTheGraph) ToLicense() (*License, error) {
	franchiseId, err := strconv.ParseInt(l.FranchiseId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert. franchise id: %s is not a valid int64. %v", l.FranchiseId, err)
	}
	ipAssetId, err := strconv.ParseInt(l.IpAssetId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert. story id: %s is not a valid int64. %v", l.IpAssetId, err)
	}
	licenseId, err := strconv.ParseInt(l.LicenseId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert. license id: %s is not a valid int64. %v", l.LicenseId, err)
	}

	license := &License{
		ID:                licenseId,
		FranchiseId:       franchiseId,
		IpAssetId:         ipAssetId,
		IpAssetName:       l.IpAssetName,
		OwnerAddress:      l.OwnerAddress,
		CollectionAddress: l.CollectionAddress,
		Scope:             l.Scope,
		Duration:          l.Duration,
		Right:             l.Right,
		ImageUrl:          l.ImageUrl,
	}
	return license, nil
}
