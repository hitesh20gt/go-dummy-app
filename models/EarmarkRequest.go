package models

import "math/big"

type EarmarkRequest struct {
	ID               string     `bson:"_id,omitempty" json:"id"`
	RequestId        string     `bson:"requestId" json:"requestId"`
	EarmarkAmount    *big.Float `bson:"earmarkAmount" json:"earmarkAmount"`
	EarmarkCurrency  string     `bson:"earmarkCurrency" json:"earmarkCurrency"`
	DebitAccount     string     `bson:"debitAccount" json:"debitAccount"`
	BusinessDate     string     `bson:"businessDate" json:"businessDate"`
	AccountBranch    string     `bson:"accountBranch" json:"accountBranch"`
	EarmarkReference string     `bson:"earmarkReference" json:"earmarkReference"`
	SourceSystem     string     `bson:"sourceSystem" json:"sourceSystem"`
	CountryCode      string     `bson:"countryCode" json:"countryCode"`
	RequestType      string     `bson:"requestType" json:"requestType"`
	PaymentType      string     `bson:"paymentType" json:"paymentType"`
	EarmarkType      string     `bson:"earmarkType" json:"earmarkType"`
}
