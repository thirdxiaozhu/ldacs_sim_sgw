// 自动生成模板AirStation
package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/internal/global"
	"time"
)

// AirStation 飞机站 结构体  AirStation
type AirStation struct {
	global.PREFIX_MODEL
	AsPlaneId  int           `json:"as_plane_id" form:"as_plane_id" gorm:"column:as_plane_id;comment:;"binding:"required"` //飞机注册号
	AsFlight   int           `json:"as_flight" form:"as_flight" gorm:"column:as_flight;comment:;"binding:"required"`       //航班号
	FlightDate *time.Time    `json:"flight_date" form:"flight_date" gorm:"column:flight_date;comment:;"binding:"required"` //执飞日期
	AsSac      string        `json:"as_sac" form:"as_sac" gorm:"column:as_sac;comment:;"binding:"required"`                //飞机站SAC
	Planeid    AccountPlane  `json:"plane_id" form:"plane_id" gorm:"foreignKey:AsPlaneId"`
	Flight     AccountFlight `json:"flight" form:"flight" gorm:"foreignKey:AsFlight"`
	CreatedBy  uint          `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint          `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint          `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 飞机站 AirStation自定义表名 air_station
func (AirStation) TableName() string {
	return "air_station"
}
