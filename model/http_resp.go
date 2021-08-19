package model

/*调用第三方返回体*/
//获取每TiBGas消耗
type GascostperResp struct {
	Result struct {
		BasefeePoints []struct {
			GasCost32g string `json:"gas_cost_32g"`
		} `json:"basefee_points"`
	}
}

type PledgePerTeraResp struct {
	Result struct {
		Data struct {
			PledgePerTera string `json:"pledge_per_tera"`
			//GasIn32g      string  `json:"gas_in_32g"`
			//AddPowerIn32g string  `json:"add_power_in_32g"`
		}
	}
}

type PowerResp struct {
	Result struct {
		Extra struct {
			Power string `json:"power"`
		} `json:"extra"`
	} `json:"result"`
}

type RewardResp struct {
	Result struct {
		Reward string `json:"blocks_rewards"`
	}
}
