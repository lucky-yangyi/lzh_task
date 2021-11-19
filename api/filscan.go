package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"lzh-tsak/consts"
	"lzh-tsak/model"
	"net/http"
	"strconv"
)
//Obtaining node revenue
func GetRewardInfo(ownerId string) (response model.Reward, err error) {
	body := fmt.Sprintf(consts.PowerBody, ownerId)
	btr := bytes.NewReader([]byte(body))
	r, err := http.Post(consts.FilsUrl, consts.ContentType, btr)
	if err != nil {
		log.Println(consts.FilsUrl, err)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	resp := model.PowerResp{}
	if err = json.Unmarshal(b, &resp); err != nil {
		panic(err)
	}
	body = fmt.Sprintf(consts.RewardBody, ownerId)
	btr = bytes.NewReader([]byte(body))
	r, err = http.Post(consts.FilsUrl, consts.ContentType, btr)
	if err != nil {
		log.Println(err)
		return
	}
	b, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	reward := model.RewardResp{}
	if err = json.Unmarshal(b, &reward); err != nil {
		log.Println(err)
		return
	}
	prd, err := strconv.ParseFloat(reward.Result.Reward, 64)
	if err != nil {
		return
	}
	//re := prd / math.Pow(10.0000, 18)
	//
	//p1, err := strconv.ParseFloat(resp.Result.Extra.Power, 64)
	//if err != nil {
	//	return
	//}
	//r1 := p1 / math.Pow(1024.00, 5)
	bp, err :=  strconv.ParseInt(resp.Result.Extra.Power, 10, 64)
	if err != nil{
		return
	}
	bt, err :=  strconv.ParseInt(reward.Result.Reward, 10, 64)
	if err != nil{
		return
	}
	response.Power =  bp
	response.Reward = bt
	return
}
//Obtain the cost of the miner's computing power
func GetMiningCost() (response model.MiningCost, err error) {
	btr := bytes.NewReader([]byte(consts.PledgeperteraBody))
	r, err := http.Post(consts.FilsUrl, consts.ContentType, btr)
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	resp := model.PledgePerTeraResp{}
	if err = json.Unmarshal(b, &resp); err != nil {
		return
	}
	btr = bytes.NewReader([]byte(consts.GascostperBody))
	r, err = http.Post(consts.FilsUrl, consts.ContentType, btr)
	if err != nil {
		return
	}
	b, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	gasresp := model.GascostperResp{}
	if err = json.Unmarshal(b, &gasresp); err != nil {
		panic(err)
	}
	var avg float64
	for _, v := range gasresp.Result.BasefeePoints {
		if v.GasCost32g != "" {
			g32, err := strconv.ParseFloat(v.GasCost32g, 64)
			if err != nil {
				continue
			}
			avg += g32
		}
	}
	l := len(gasresp.Result.BasefeePoints)
	avg = avg / float64(l)
	per, err := strconv.ParseFloat(resp.Result.Data.PledgePerTera, 64)
	if err != nil {
		return
	}
	response = model.MiningCost{
		SectorSize:       "32GiB",
		GasCostPerTiB:    fmt.Sprintf("%.5f", avg),
		PledgeCostPerTiB: fmt.Sprintf("%.5f", per),
		TotalCostPerTiB:  fmt.Sprintf("%.5f", per+avg),
	}
	return
}
