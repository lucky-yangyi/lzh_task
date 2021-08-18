package model

type Reward struct {
	Power     string
	Reward    string
	CreatedAt string
	OwnerId string
}

type Miners struct {
	Id  int  `json:"id" gorm:"index"`
	OwnerId string `json:"owner_id"`
}
