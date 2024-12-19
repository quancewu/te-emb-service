package models

import (
	"time"

	"gorm.io/gorm"
)

type Mix_sec_struct struct {
	// gorm.Model
	ID              uint           `gorm:"primarykey" json:"-"`
	CreatedAt       time.Time      `gorm:"index" json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	System_Id       int            `json:"-" gorm:"uniqueIndex:system_id_time_uniq"`
	Serial_Number   string         `json:"-"`
	Timestamp       time.Time      `json:"timestamp" gorm:"index;uniqueIndex:system_id_time_uniq"`
	RSSI            *float64       `json:"rssi"`
	Timestamp_SU    time.Time      `json:"timestamp_su"`
	Temperature     *float64       `json:"temperature"`
	Relative_H      *float64       `json:"relative_humidity"`
	Pressure        *float64       `json:"pressure"`
	PM_1            *float64       `json:"pm_1"`
	PM_2p5          *float64       `json:"pm_2p5"`
	PM_10           *float64       `json:"pm_10"`
	CO_2            *float64       `json:"co_2"`
	Wind_Speed      *float64       `json:"wind_speed"`
	Wind_Direction  *float64       `json:"wind_direction"`
	Input_volt      *float64       `json:"input_voltage"`
	Solar_irrad     *float64       `json:"solar_irradirance"`
	Pyr_temperature *float64       `json:"pyr_temperature"`
	Bat             *float64       `json:"bat"`
	Raw             string         `json:"raw"`
	Flag            int            `json:"flag"`
}

type Boot_seq_record struct {
	ID        int            `gorm:"primarykey" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Boot_Id   string         `gorm:"index" json:"boot_id"`
	Path      string         `json:"path"`
}

type Ams_serial_number struct {
	ID            int            `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	System_Id     int            `json:"system_id" gorm:"uniqueIndex:system_id_uniq"`
	Serial_Number string         `json:"serial_number" gorm:"uniqueIndex:system_id_uniq"`
}

type AmsData struct {
	System_Id       int       `json:"system_id" gorm:"uniqueIndex:system_id_time_uniq"`
	Serial_Number   string    `json:"serial_number"`
	Timestamp       time.Time `json:"timestamp" gorm:"uniqueIndex:system_id_time_uniq"`
	RSSI            *float64  `json:"rssi"`
	Timestamp_SU    time.Time `json:"timestamp_su"`
	Temperature     *float64  `json:"temperature"`
	Relative_H      *float64  `json:"relative_humidity"`
	Pressure        *float64  `json:"pressure"`
	PM_1            *float64  `json:"pm_1"`
	PM_2p5          *float64  `json:"pm_2p5"`
	PM_10           *float64  `json:"pm_10"`
	CO_2            *float64  `json:"co_2"`
	Wind_Speed      *float64  `json:"wind_speed"`
	Wind_Direction  *float64  `json:"wind_direction"`
	Input_volt      *float64  `json:"input_voltage"`
	Solar_irrad     *float64  `json:"solar_irradirance"`
	Pyr_temperature *float64  `json:"pyr_temperature"`
	Bat             *float64  `json:"bat"`
	Raw             string    `json:"raw"`
	Flag            int       `json:"flag"`
}
