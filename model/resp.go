package model

type Response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type MiningCost struct {
	SectorSize       string //32GiB
	GasCostPerTiB    string //每TiBGas消耗 /7天平均值
	PledgeCostPerTiB string //每TiB质押量
	TotalCostPerTiB  string //每TiB新增算力成本
}

type FilscanInfo struct {
	QualityAdjPower string `json:"qualityAdjPower"`
	Reward24Hours   string `json:"reward24Hours"`
}
