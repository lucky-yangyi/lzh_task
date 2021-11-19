package task

import (
	"github.com/robfig/cron"
	"log"
	"lzh-tsak/api"
	"lzh-tsak/consts"
	"lzh-tsak/dao"
	"time"
)

//A scheduled task to obtain node benefits
func TaskStart() {
	c := cron.New() //精确到秒
	if err := c.AddFunc("0 30 00 * * *", SpiderFilscan); err != nil {
		panic(err)
	}
	c.Start()
}

func SpiderFilscan() {
	list, err := dao.GetOwnerIdList()
	if err != nil {
		log.Println("获取onwerid失败", err)
		return
	}
	for _, v := range list {
		info, err := api.GetRewardInfo(v.OwnerId)
		if err != nil {
			continue
		}
		info.CreatedAt = time.Now().AddDate(0, 0, -1).Format(consts.YYYYMMDD)
		info.OwnerId = v.OwnerId
		err = dao.CreateReward(info)
		if err != nil {
			log.Println(err)
			continue
		}
	}

}
