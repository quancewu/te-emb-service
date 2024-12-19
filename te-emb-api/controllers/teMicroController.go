package controllers

import (
	"fmt"
	"net/http"
	"te-emb-api/initalizers"
	"te-emb-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

type AmsId struct {
	ID string `uri:"id" binding:"required"`
}

type AmsQueryPara struct {
	After  time.Time `form:"after" time_format:"2006-01-02T15:04:05Z07:00"`
	Before time.Time `form:"before" time_format:"2006-01-02T15:04:05Z07:00"`
}

type AmsPostPara struct {
	System_Id     int       `json:"system_id"`
	Serial_Number string    `json:"serial_number"`
	Timestamp     time.Time `json:"timestamp" time_format:"2006-01-02T15:04:05Z07:00"`
	RSSI          *float64  `json:"rssi"`
	Timestamp_SU  time.Time `json:"timestamp_su" time_format:"2006-01-02T15:04:05Z07:00"`
	Temperature   *float64  `json:"temperature"`
	Relative_H    *float64  `json:"relative_humidity"`
	Pressure      *float64  `json:"pressure"`
	PM_1          *float64  `json:"pm_1"`
}

type AmsResposeModel struct {
	System_Id     int                   `json:"system_id"`
	Serial_Number string                `json:"serial_number"`
	Data          models.Mix_sec_struct `json:"data"`
}

type AmsResposeModels struct {
	System_Id     int                     `json:"system_id"`
	Serial_Number string                  `json:"serial_number"`
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

func TeAmsData(c *gin.Context) {
	// Get the id request uri
	var amsId AmsId
	if err := c.ShouldBindUri(&amsId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	var intId int
	if intId = parseId(amsId.ID); intId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "id not found"})
		return
	} else {
		var ams_sn models.Ams_serial_number
		if res := initalizers.DBLS.Where(&models.Ams_serial_number{System_Id: intId}).First(&ams_sn); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("id: %s not found in table", amsId.ID)})
			return
		}
	}

	var amsGet models.Mix_sec_struct
	res := initalizers.DBLS.Last(&amsGet, "system_id=?", intId)
	// res := initalizers.DBL.Where("created_at BETWEEN ? AND ?", amsQueryPara.After, amsQueryPara.Before).Find(&amsGet)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Fail to get data"})
		return
	}

	c.JSON(http.StatusOK, AmsResposeModel{
		System_Id:     intId,
		Serial_Number: amsId.ID,
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
	var intId int
	if intId = parseId(amsId.ID); intId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "id not found"})
	} else {
		var ams_sn models.Ams_serial_number
		if res := initalizers.DBLS.Where(&models.Ams_serial_number{System_Id: intId}).First(&ams_sn); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "id not found in table", "error": res.Error})
			return
		}
	}

	var amsGet []models.Mix_sec_struct
	res := initalizers.DBLS.Where("timestamp BETWEEN ? AND ? and system_id=?", amsQueryPara.After, amsQueryPara.Before, intId).Find(&amsGet)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Fail to get data"})
		return
	}

	c.JSON(http.StatusOK, AmsResposeModels{
		System_Id:     intId,
		Serial_Number: amsId.ID,
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
		System_Id:     amsPostPara.System_Id,
		Serial_Number: amsPostPara.Serial_Number,
		Timestamp:     amsPostPara.Timestamp,
		RSSI:          amsPostPara.RSSI,
		Timestamp_SU:  amsPostPara.Timestamp_SU,
		Temperature:   amsPostPara.Temperature,
		Relative_H:    amsPostPara.Relative_H,
		Pressure:      amsPostPara.Pressure,
		PM_1:          amsPostPara.RSSI,
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
