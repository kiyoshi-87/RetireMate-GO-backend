package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Userdata struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email              string             `json:"email"`
	Total_expenses     string             `json:"total_expenses"`
	Total_earnings     string             `json:"total_earnings"`
	Estimated_networth string             `json:"estimated_networth"`
	Targeted_networth  string             `json:"targeted_networth"`
	Cashflow           string             `json:"cashflow"`
	Stocks             string             `json:"stocks"`
	Real_estate        string             `json:"real_estate"`
	Gold               string             `json:"gold"`
	FD                 string             `json:"fd"`
	FD_interests       string             `json:"fd_interests"`
	Time_to_retire     string             `json:"time_to_retire"`
}
