package orders

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderId     uint         `gorm:"column:order_id;primaryKey;autoIncrement" json:"order_id"`
	City        string       `gorm:"column:city" json:"city"`
	District    string       `gorm:"column:district" json:"district"`
	Address     string       `gorm:"column:address" json:"address"`
	OrderTime   time.Time    `gorm:"column:order_time" json:"orderTime"`
	PickUpDrop  string       `gorm:"column:pick_up_drop" json:"pickUpDrop"`
	PickUpTime  time.Time    `gorm:"column:pick_up_time" json:"pickUpTime"`
	Weekday     time.Weekday `gorm:"column:weekday" json:"weekday"`
	Group       string       `gorm:"column:group" json:"group"`
	Amount      float64      `gorm:"column:amount" json:"amount"`
	Distance    float64      `gorm:"column:distance" json:"distance"`
	IsException bool         `gorm:"column:is_exception" json:"isException"`
	RepeatCount int          `gorm:"column:repeat_count" json:"repeatCount"`
}
