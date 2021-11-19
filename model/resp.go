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
	QualityAdjPower int64 `json:"qualityAdjPower"`
	Reward24Hours   int64 `json:"reward24Hours"`
}
