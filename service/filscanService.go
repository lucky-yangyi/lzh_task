package service

import (
	"lzh-tsak/api"
	"lzh-tsak/consts"
	"lzh-tsak/dao"
	"lzh-tsak/model"
	"time"
)

func GetFilsInfo(ownerId string, date string) (resp model.FilscanInfo, err error) {
	var (
		start time.Time
		end   time.Time
	)
	if date == "" {
		t, err := time.Parse(consts.YYYYMMDD, time.Now().AddDate(0, 0, -1).Format(consts.YYYYMMDD))
		if err != nil {
			return resp, err
		}
		start = t
		end = start.AddDate(0, 0, 1)
	} else {
		t, err := time.Parse(consts.YYYYMMDD, date)
		if err != nil {
			return resp, err
		}
		start = t
		end = start.AddDate(0, 0, 1)
	}
	reward, err := dao.GetReward(ownerId, start, end)
	if err != nil {
		if date == "" {
			reward, err = api.GetRewardInfo(ownerId)
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	//p,err := strconv.ParseInt(reward.Reward, 10, 64)
	//if err != nil {
	//	return
	//}

	resp = model.FilscanInfo{
		QualityAdjPower: reward.Power,
		Reward24Hours:   reward.Reward,
	}
	return
}

func GetListPower() (resp model.MiningCost, err error) {
	return api.GetMiningCost()
}
