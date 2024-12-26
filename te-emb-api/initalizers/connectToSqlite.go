package initalizers

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"te-emb-api/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBL *gorm.DB
var DBLS *gorm.DB
var DBLSB *gorm.DB

func ConnectToSQLITE() {
	var err error
	store_path := os.Getenv("STORAGE")
	dsn := fmt.Sprintf("%s/te-emb-api.db?_journal_mode=WAL", store_path)
	log.Printf("te-emb-api | sqlite: %s", dsn)
	DBL, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("%s", err.Error())
		panic("Failed to connect to db")
	}
	// 设置数据库连接池参数
	sqlDB, err := DBL.DB()
	if err != nil {
		log.Printf("%s", err.Error())
		panic("Failed to connect to db")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DBL.AutoMigrate(&models.Mix_sec_struct{})
	DBL.AutoMigrate(&models.Boot_seq_record{})
}

func ConnectToSqliteTimeseries() {
	var err error
	store_path := os.Getenv("STORAGE")
	cmd := exec.Command("cat", "/proc/sys/kernel/random/boot_id")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	exists_or_create(store_path)
	dsn := fmt.Sprintf("%s/%s.db?_journal_mode=WAL", store_path, strings.Trim(stdout.String(), "\n"))
	log.Printf("te-emb-api timeseries | sqlite: %s", dsn)
	DBLS, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("%s", err.Error())
		panic("Failed to connect to db")
	}
	// set dv connection pool number
	sqlDB, err := DBLS.DB()
	if err != nil {
		log.Printf("%s", err.Error())
		panic("Failed to connect to db")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DBLS.AutoMigrate(&models.Mix_sec_struct{})
	DBLS.AutoMigrate(&models.Ams_serial_number{})
}

func ConnectToBackupSqliteTimeseries(dbname string) {
	var err error
	store_path := os.Getenv("STORAGE")
	exists_or_create(store_path)
	dsn := fmt.Sprintf("%s/%s.db?_journal_mode=WAL", store_path, dbname)
	log.Printf("te-emb-api       timeseries | sqlite: %s", dsn)
	DBLSB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("%s", err.Error())
		panic("Failed to connect to db")
	}
	// set dv connection pool number
	sqlDB, err := DBLSB.DB()
	if err != nil {
		log.Printf("%s", err.Error())
		panic("Failed to connect to db")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DBLS.AutoMigrate(&models.Mix_sec_struct{})
	DBLS.AutoMigrate(&models.Ams_serial_number{})
}

func DisconnectBackupSqlite() {
	sqlDB, err := DBLSB.DB()
	if err != nil {
		log.Printf("%s", err.Error())
		panic("Failed to connect to db")
	}
	sqlDB.Close()
}

// exists returns whether the given file or directory exists
func exists_or_create(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
		return true, nil
	} else {
		return false, err
	}
}
