package models

type NetworkCardList struct {
	IsMaster bool   `json:"isMaster"`
	SubnetID string `json:"subnetID"`
}
