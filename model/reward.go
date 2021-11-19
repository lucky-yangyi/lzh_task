package model

type Reward struct {
	Power     int64
	Reward    int64
	CreatedAt string
	OwnerId   string
}

type Miners struct {
	//Id  int  `json:"id" gorm:"index"`
	OwnerId string `json:"owner_id"`
}
