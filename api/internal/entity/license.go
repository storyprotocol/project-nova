package entity

import (
	"fmt"
	"strconv"
)

type License struct {
	ID                int64  `json:"id"`
	FranchiseId       int64  `json:"franchiseId"`
	IpAssetId         int64  `json:"ipAssetId"`
	IpAssetName       string `json:"ipAssetName"`
	OwnerAddress      string `json:"ownerAddress"`
	CollectionAddress string `json:"collectionAddress"`
	Scope             string `json:"scope"`
	Duration          string `json:"duration"`
	Right             string `json:"right"`
	ImageUrl          string `json:"imageUrl"`
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
