/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/9/12 14:00
*  @To:
 */

package models

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

// Single 主
type Single struct {
	Id           int        `primaryKey:"true" json:"id"`
	Name         string     `json:"name"`
	BridgeDeckID int        `gorm:"default:0;comment:'桥面id'" json:"bridgeDeckID"`
	BridgeDeck   BridgeDeck `gorm:"foreignKey:BridgeDeckID"` //桥面系要素选择
}

// BA 桥面铺装 数据
type BA struct {
	ID   int     `json:"id"`
	BA1  float64 `json:"BA1"`  //网裂或龟裂
	BA2  float64 `json:"BA2"`  //波浪及车辙
	BA3  float64 `json:"BA3"`  //坑槽
	BA4  float64 `json:"BA4"`  //碎裂或破碎
	BA5  float64 `json:"BA5"`  //坑洞
	BA6  float64 `json:"BA6"`  //桥面贯通横缝
	BA7  float64 `json:"BA7"`  //桥面贯通纵缝
	MdpH float64 `json:"MDPh"` //总扣分
}

// BB 桥头平顺
type BB struct {
	ID   int     `json:"id"`
	BB1  float64 `json:"BB1"`  //桥头沉降
	BB2  float64 `json:"BB2"`  //台背下沉
	MdpH float64 `json:"MDPh"` //总扣分
}

// BC 伸缩装置
type BC struct {
	ID   int     `json:"id"`
	BC1  float64 `json:"BC1"`  //螺帽松动
	BC2  float64 `json:"BC2"`  //缝内沉积物阻塞
	BC3  float64 `json:"BC3"`  //止水带破损、老化
	BC4  float64 `json:"BC4"`  //钢材料破损
	BC5  float64 `json:"BC5"`  //接缝处铺装碎边
	BC6  float64 `json:"BC6"`  //接缝处高差
	BC7  float64 `json:"BC7"`  //钢材翘曲变形
	BC8  float64 `json:"BC8"`  //结构缝宽异常
	BC9  float64 `json:"BC9"`  //伸缩缝处异常声响
	MdpH float64 `json:"MDPh"` //总扣分
}

// BD 排水系统
type BD struct {
	ID   int     `json:"id"`
	BD1  float64 `json:"BD1"`  //泄水管堵塞
	BD2  float64 `json:"BD2"`  //残缺脱落
	BD3  float64 `json:"BD3"`  //桥面积水
	BD4  float64 `json:"BD4"`  //防水层
	MdpH float64 `json:"MDPh"` //总扣分
}

// BE 人行道
type BE struct {
	ID   int     `json:"id"`
	BE1  float64 `json:"BE1"`  //网裂
	BE2  float64 `json:"BE2"`  //松动或变形
	BE3  float64 `json:"BE3"`  //残缺
	MdpH float64 `json:"MDPh"` //总扣分
}

// BF 栏杆及护栏
type BF struct {
	ID   int     `json:"id"`
	BF1  float64 `json:"BF1"`  //露筋锈蚀
	BF2  float64 `json:"BF2"`  //松动错位
	BF3  float64 `json:"BF3"`  //丢失残缺
	MdpH float64 `json:"MDPh"` //总扣分
}

type BridgeDeck struct {
	Id       int  `primaryKey:"true" json:"id"`
	BAId     uint `gorm:"default:0;comment:'桥面铺装id'" json:"BAId"`
	BBId     uint `gorm:"default:0;comment:'桥头平顺id'" json:"BBId"`
	BCId     uint `gorm:"default:0;comment:'伸缩装置id'" json:"BCId"`
	BDId     uint `gorm:"default:0;comment:'排水系统id'" json:"BDId"`
	BEId     uint `gorm:"default:0;comment:'排水系统id'" json:"BEId"`
	BFId     uint `gorm:"default:0;comment:'人行道id'" json:"BFId"`
	ElementA BA   `gorm:"foreignKey:BAId"` //桥面铺装
	ElementB BB   `gorm:"foreignKey:BBId"` //桥头平顺
	ElementC BC   `gorm:"foreignKey:BCId"` //伸缩装置
	ElementD BD   `gorm:"foreignKey:BDId"` //排水系统
	ElementE BE   `gorm:"foreignKey:BEId"` //人行道
	ElementF BF   `gorm:"foreignKey:BFId"` //栏杆及护栏

}

type Ips struct {
	gorm.Model
	IPList ipList `json:"ip"`
}

type ipList []string

func (p ipList) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan 实现方法
func (p *ipList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &p)
}
