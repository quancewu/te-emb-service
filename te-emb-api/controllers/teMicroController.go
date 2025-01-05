package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"te-emb-api/initalizers"
	"te-emb-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AmsId struct {
	ID string `uri:"id" binding:"required"`
}

type AmsLatestBootId struct {
	Boot_Id int `form:"boot_id"`
}

type AmsQueryPara struct {
	After   time.Time `form:"after" time_format:"2006-01-02T15:04:05Z07:00"`
	Before  time.Time `form:"before" time_format:"2006-01-02T15:04:05Z07:00"`
	Boot_Id int       `form:"boot_id"`
}

type AmsPostPara struct {
	System_Id       int       `json:"system_id"`
	Serial_Number   string    `json:"serial_number"`
	Timestamp       time.Time `json:"timestamp" time_format:"2006-01-02T15:04:05Z07:00"`
	RSSI            *float64  `json:"rssi"`
	Timestamp_SU    time.Time `json:"timestamp_su" time_format:"2006-01-02T15:04:05Z07:00"`
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

type AmsDbResposeModel struct {
	Hostname string                   `json:"hostname"`
	Status   string                   `json:"status"`
	Data     []models.Boot_seq_record `json:"data"`
}

type AmsDeviceResposeModel struct {
	Hostname string                     `json:"hostname"`
	Status   string                     `json:"status"`
	Data     []models.Ams_serial_number `json:"data"`
}

type AmsResposeModel struct {
	System_Id     int                   `json:"system_id"`
	Serial_Number string                `json:"serial_number"`
	Status        string                `json:"status"`
	Data          models.Mix_sec_struct `json:"data"`
}

type AmsResposeModels struct {
	System_Id     int                     `json:"system_id"`
	Serial_Number string                  `json:"serial_number"`
	Status        string                  `json:"status"`
	After         time.Time               `json:"after"`
	Before        time.Time               `json:"before"`
	Data          []models.Mix_sec_struct `json:"data"`
}

func parseId(idstr string) int {
	var sn int
	fmt.Sscanf(idstr, "te-su-%d", &sn)
	sn = sn % 1000
	return sn
}

func TeAmsDB(c *gin.Context) {
	var ams_boot_seq []models.Boot_seq_record
	if res := initalizers.DBL.Find(&ams_boot_seq); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "no boot seq table"})
		return
	}
	hostname, _ := os.Hostname()
	c.JSON(http.StatusOK, AmsDbResposeModel{
		Hostname: hostname,
		Status:   "okay",
		Data:     ams_boot_seq,
	})
}

func TeAmsDevices(c *gin.Context) {
	var ams_sn_list []models.Ams_serial_number
	if res := initalizers.DBLS.Find(&ams_sn_list); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "no ams sn found"})
		return
	}
	hostname, _ := os.Hostname()
	c.JSON(http.StatusOK, AmsDeviceResposeModel{
		Hostname: hostname,
		Status:   "okay",
		Data:     ams_sn_list,
	})
}

func TeAmsData(c *gin.Context) {
	// Get the id request uri
	var amsId AmsId
	if err := c.ShouldBindUri(&amsId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	var amsBootId AmsLatestBootId
	if err := c.ShouldBind(&amsBootId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	}
	var query_db *gorm.DB
	if amsBootId.Boot_Id >= 0 {
		query_db = initalizers.DBLS
	} else if amsBootId.Boot_Id < -100 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "boot_id need larger than -100"})
		return
	} else if amsBootId.Boot_Id < 0 {
		var ams_boot_seq []models.Boot_seq_record
		if res := initalizers.DBL.Limit(100).Order("created_at desc").Find(&ams_boot_seq); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("no boot seq db %d", amsBootId.Boot_Id)})
			return
		}
		if len(ams_boot_seq)+amsBootId.Boot_Id < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("ams_boot_seq lens: %d, boot_id: %d", len(ams_boot_seq), amsBootId.Boot_Id)})
		}
		dbname := ams_boot_seq[-amsBootId.Boot_Id].Boot_Id
		log.Printf("dbname: %s\n", dbname)
		initalizers.ConnectToBackupSqliteTimeseries(dbname)
		defer initalizers.DisconnectBackupSqlite()
		query_db = initalizers.DBLSB
	}
	var intId int
	if intId = parseId(amsId.ID); intId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "id not found"})
		return
	} else {
		var ams_sn models.Ams_serial_number
		if res := query_db.Where(&models.Ams_serial_number{System_Id: intId}).First(&ams_sn); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("id: %s not found in table", amsId.ID)})
			return
		}
	}

	var amsGet models.Mix_sec_struct
	res := query_db.Last(&amsGet, "system_id=?", intId)
	// res := initalizers.DBL.Where("created_at BETWEEN ? AND ?", amsQueryPara.After, amsQueryPara.Before).Find(&amsGet)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Fail to get data"})
		return
	}

	c.JSON(http.StatusOK, AmsResposeModel{
		System_Id:     intId,
		Serial_Number: amsId.ID,
		Status:        "okay",
		Data:          amsGet,
	})
}

func TeAmsDatas(c *gin.Context) {
	// Get the id request uri
	var amsId AmsId
	if err := c.ShouldBindUri(&amsId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	var amsQueryPara AmsQueryPara
	if err := c.ShouldBind(&amsQueryPara); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	log.Printf("boot_id %d\n", amsQueryPara.Boot_Id)
	var query_db *gorm.DB
	if amsQueryPara.Boot_Id >= 0 {
		query_db = initalizers.DBLS
	} else if amsQueryPara.Boot_Id < -100 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "boot_id need larger than -100"})
		return
	} else if amsQueryPara.Boot_Id < 0 {
		var ams_boot_seq []models.Boot_seq_record
		if res := initalizers.DBL.Limit(100).Order("created_at desc").Find(&ams_boot_seq); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("no boot seq db %d", amsQueryPara.Boot_Id)})
			return
		}
		if len(ams_boot_seq)+amsQueryPara.Boot_Id < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("ams_boot_seq lens: %d, boot_id: %d", len(ams_boot_seq), amsQueryPara.Boot_Id)})
		}
		dbname := ams_boot_seq[-amsQueryPara.Boot_Id].Boot_Id
		log.Printf("dbname: %s\n", dbname)
		initalizers.ConnectToBackupSqliteTimeseries(dbname)
		defer initalizers.DisconnectBackupSqlite()
		query_db = initalizers.DBLSB
	}

	var intId int
	if intId = parseId(amsId.ID); intId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "id not found"})
	} else {
		var ams_sn models.Ams_serial_number
		if res := query_db.Where(&models.Ams_serial_number{System_Id: intId}).First(&ams_sn); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "id not found in table", "error": res.Error})
			return
		}
	}

	var amsGet []models.Mix_sec_struct
	res := query_db.Where("timestamp BETWEEN ? AND ? and system_id=?", amsQueryPara.After, amsQueryPara.Before, intId).Find(&amsGet)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Fail to get data"})
		return
	}

	c.JSON(http.StatusOK, AmsResposeModels{
		System_Id:     intId,
		Serial_Number: amsId.ID,
		Status:        "okay",
		After:         amsQueryPara.After,
		Before:        amsQueryPara.Before,
		Data:          amsGet,
	})
}

func TeAmsDataInsert(c *gin.Context) {
	var amsId AmsId
	if err := c.ShouldBindUri(&amsId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	var amsPostPara AmsPostPara
	if err := c.ShouldBind(&amsPostPara); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	amsPost := models.Mix_sec_struct{
		System_Id:       amsPostPara.System_Id,
		Serial_Number:   amsPostPara.Serial_Number,
		Timestamp:       amsPostPara.Timestamp,
		RSSI:            amsPostPara.RSSI,
		Timestamp_SU:    amsPostPara.Timestamp_SU,
		Temperature:     amsPostPara.Temperature,
		Relative_H:      amsPostPara.Relative_H,
		Pressure:        amsPostPara.Pressure,
		PM_1:            amsPostPara.RSSI,
		PM_2p5:          amsPostPara.PM_2p5,
		PM_10:           amsPostPara.PM_10,
		CO_2:            amsPostPara.CO_2,
		Wind_Speed:      amsPostPara.Wind_Speed,
		Wind_Direction:  amsPostPara.Wind_Direction,
		Input_volt:      amsPostPara.Input_volt,
		Solar_irrad:     amsPostPara.Solar_irrad,
		Pyr_temperature: amsPostPara.Pyr_temperature,
		Raw:             "",
		Flag:            201,
	}
	result := initalizers.DBL.Create(&amsPost)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to insert amsPost",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messsage": "success add ams data",
	})
}
