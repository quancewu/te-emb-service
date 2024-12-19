package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"te-redis-service/initalizers"
	"te-redis-service/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func float64_nan_check_print(value *float64) string {
	if value != nil {
		if math.IsNaN(*value) {
			return "null"
		} else {
			converted := fmt.Sprintf("%.2f", *value)
			return converted
		}
	} else {
		return "null"
	}
}

func InsertAmsData(m models.Message) error {
	// log.Printf("Data received: %s", m.Content)
	var ams models.AmsData
	if payload, err := json.Marshal(m.Metadata); err != nil {
		return nil
	} else {
		if err = json.Unmarshal(payload, &ams); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			return nil
		}
	}

	// log.Printf("AmsData ID: %s", ams.Serial_Number)
	// log.Printf("AmsData  T: %v", float64_nan_check_print(ams.Temperature))
	// log.Printf("AmsData RH: %v", float64_nan_check_print(ams.Relative_H))

	amsPost := models.Mix_sec_struct{
		System_Id:       ams.System_Id,
		Serial_Number:   ams.Serial_Number,
		Timestamp:       ams.Timestamp,
		RSSI:            ams.RSSI,
		Timestamp_SU:    ams.Timestamp_SU,
		Temperature:     ams.Temperature,
		Relative_H:      ams.Relative_H,
		Pressure:        ams.Pressure,
		PM_1:            ams.PM_1,
		PM_2p5:          ams.PM_2p5,
		PM_10:           ams.PM_10,
		CO_2:            ams.CO_2,
		Wind_Speed:      ams.Wind_Speed,
		Wind_Direction:  ams.Wind_Direction,
		Input_volt:      ams.Input_volt,
		Solar_irrad:     ams.Solar_irrad,
		Pyr_temperature: ams.Pyr_temperature,
		Bat:             ams.Bat,
		Raw:             ams.Raw,
		Flag:            ams.Flag,
	}
	// check := UpSert(models.Ams_serial_number{
	// 	System_Id:     ams.System_Id,
	// 	Serial_Number: ams.Serial_Number,
	// })
	// if check != nil {
	// 	log.Printf("fail to upsert %v", check)
	// }
	var ams_sn models.Ams_serial_number
	if result := initalizers.DBLS.Where("system_id=?", ams.System_Id).First(&ams_sn); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("fail to find sn: %s ams_serial_numbers table %s", ams.Serial_Number, result.Error)
			ams_sn := models.Ams_serial_number{
				System_Id:     ams.System_Id,
				Serial_Number: ams.Serial_Number,
			}
			if res := initalizers.DBLS.Create(&ams_sn); res.Error != nil {
				log.Printf("fail to insert ams sn ams_serial_numbers table %s", res.Error)
			}
		}
	}

	result := initalizers.DBLS.Create(&amsPost)
	if result.Error != nil {
		log.Printf("fail to insert ams %v", result.Error)
	}
	return nil
}

func UpSert(m models.Ams_serial_number) error {
	return initalizers.DBLS.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&m).Error
}
