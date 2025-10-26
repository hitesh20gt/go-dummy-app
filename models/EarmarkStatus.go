package models

import "math/big"

type EarmarkStatus struct {
	ID               string     `bson:"_id,omitempty" json:"id"`
	EarmarkAmount    string     `bson:"earmarkAmount" json:"earmarkAmount"`
	EarmarkCurrency  string     `bson:"earmarkCurrency" json:"earmarkCurrency"`
	DebitAccount     *big.Float `bson:"debitAccount" json:"debitAccount"`
	BusinessDate     string     `bson:"businessDate" json:"businessDate"`
	AccountBranch    string     `bson:"accountBranch" json:"accountBranch"`
	EarmarkReference string     `bson:"earmarkReference" json:"earmarkReference"`
	SourceSystem     string     `bson:"sourceSystem" json:"sourceSystem"`
	Status           string     `bson:"status" json:"status"`
}
