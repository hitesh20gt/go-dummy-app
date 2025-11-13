package models

type EarmarkStatusRequest struct {
	DebitAccount  string `bson:"debitAccount" json:"debitAccount"`
	AccountBranch string `bson:"accountBranch" json:"accountBranch"`
}
