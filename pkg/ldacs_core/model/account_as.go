// 自动生成模板AccountAs
package model

import (
	"ldacs_sim_sgw/internal/global"

	"time"
)

// 飞机站账户 结构体  AccountAs
type AccountAs struct {
	global.PREFIX_MODEL
	AsPlaneId int           `json:"as_plane_id" form:"as_plane_id" gorm:"column:as_plane_id;comment:;"` //飞机注册号
	AsFlight  int           `json:"as_flight" form:"as_flight" gorm:"column:as_flight;comment:;"`       //执飞航班号
	AsDate    *time.Time    `json:"as_date" form:"as_date" gorm:"column:as_date;comment:;"`             //执飞日期
	AsSac     int           `json:"as_sac" form:"as_sac" gorm:"column:as_sac;comment:;" `               //飞机站SAC
	AsState   int           `json:"as_state" form:"as_state" gorm:"column:as_state;comment:;"`
	Planeid   AccountPlane  `json:"plane_id" form:"plane_id" gorm:"foreignKey:AsPlaneId"`
	Flight    AccountFlight `json:"flight" form:"flight" gorm:"foreignKey:AsFlight"`
	CreatedBy uint          `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint          `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint          `gorm:"column:deleted_by;comment:删除者"`
}

type AccountAsOptions struct {
	AsPlaneIds []AccountPlane  `json:"plane_ids"`
	AsFlights  []AccountFlight `json:"flights"`
}

// TableName 飞机站账户 AccountAs自定义表名 account_as
func (AccountAs) TableName() string {
	return "account_as"
}
