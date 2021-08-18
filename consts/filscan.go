package consts


const (
	FilsUrl = `https://api.filscan.io:8700/rpc/v1`
	PowerBody = `{"id":1,"jsonrpc":"2.0","params":["0.5y","%s"],"method":"filscan.OrePoolCurve"}`
	RewardBody = `{"id": 1, "jsonrpc": "2.0", "params": [["%s"], "1d", 3],"method": "filscan.FilscanStatisticalIndicatorsUnite"}`
	ContentType = `application/json`
	PledgeperteraBody = `{"id":1,"jsonrpc":"2.0","params":[],"method":"filscan.StatChainInfo"}`
	GascostperBody    =  `{"id": 1,"jsonrpc": "2.0","params": ["1w"],"method": "filscan.BaseFeeTrend"}`
)


