package entity

import (
	"fmt"
	"strconv"
)

type License struct {
	ID                int64  `json:"id"`
	FranchiseId       int64  `json:"franchiseId"`
	StoryId           int64  `json:"storyId"`
	StoryName         string `json:"storyName"`
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
	StoryId           string `json:"storyId"`
	StoryName         string `json:"storyName"`
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
	storyId, err := strconv.ParseInt(l.StoryId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert. story id: %s is not a valid int64. %v", l.StoryId, err)
	}
	licenseId, err := strconv.ParseInt(l.LicenseId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert. license id: %s is not a valid int64. %v", l.LicenseId, err)
	}

	license := &License{
		ID:                licenseId,
		FranchiseId:       franchiseId,
		StoryId:           storyId,
		StoryName:         l.StoryName,
		OwnerAddress:      l.OwnerAddress,
		CollectionAddress: l.CollectionAddress,
		Scope:             l.Scope,
		Duration:          l.Duration,
		Right:             l.Right,
		ImageUrl:          l.ImageUrl,
	}
	return license, nil
}
