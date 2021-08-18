package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"lzh-tsak/model"
	"time"
)


var db *gorm.DB

const (
	TableReward = "reward"
)

func init() {
	//用户名：密码^@tcp(地址:3306)/数据库
	var err error
	db,err = gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/blog?charset=utf8")
	if err!=nil {
		fmt.Println(err)
		return
	}
	db.SingularTable(true)
	return
}

func GetDB()*gorm.DB {
	if db == nil{
		return nil
	}
	return db
}


func GetOwnerIdList()(minersList []model.Miners,err error){
	//return []string{"f0154944","f0688083"},nil
	db.Raw(`select id,owner_id from miners`).Scan(&minersList)
	return
}

func CreateReward(reward model.Reward)error {
	return GetDB().Create(&reward).Error
}


func GetReward(ownerId string ,start,end time.Time )(data model.Reward,err error){
	err = GetDB().Table(TableReward).Where("owner_id = ? and created_at >= ? and created_at < ? ",ownerId,start,end).Find(&data).Error
	return
}