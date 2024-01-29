// 自动生成模板AccountAs
package model

import (
	"ldacs_sim_sgw/internal/global"

	"time"
)

// 飞机站账户 结构体  AccountAs
type AccountAs struct {
	global.PREFIX_MODEL
	AsPlaneId      int           `json:"as_plane_id" form:"as_plane_id" gorm:"column:as_plane_id;comment:;"` //飞机注册号
	AsFlight       int           `json:"as_flight" form:"as_flight" gorm:"column:as_flight;comment:;"`       //执飞航班号
	AsDate         *time.Time    `json:"as_date" form:"as_date" gorm:"column:as_date;comment:;"`             //执飞日期
	AsCurrState    int           `json:"as_state" form:"as_state" gorm:"column:as_state;comment:;"`
	StateID        uint          `json:"state_id" form:"state_id" gorm:"column:state_id;comment:;"`
	Planeid        AccountPlane  `json:"plane_id" form:"plane_id" gorm:"foreignKey:AsPlaneId;references:ID"`
	Flight         AccountFlight `json:"flight" form:"flight" gorm:"foreignKey:AsFlight;references:ID"`
	State          *State        `json:"state" form:"state" gorm:"foreignKey:StateID;references:ID"`
	DeprecatedTime *time.Time    `json:"deprecated_time" form:"deprecated_time" gorm:"column:deprecated_time;comment:;"` //执飞日期
	CreatedBy      uint          `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint          `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint          `gorm:"column:deleted_by;comment:删除者"`
	DeprecatedBy   uint          `gorm:"column:deprecated_by;comment:删除者"`
}

type AccountAsOptions struct {
	AsPlaneIds []AccountPlane  `json:"plane_ids"`
	AsFlights  []AccountFlight `json:"flights"`
}

// TableName 飞机站账户 AccountAs自定义表名 account_as
func (AccountAs) TableName() string {
	return "account_as"
}
