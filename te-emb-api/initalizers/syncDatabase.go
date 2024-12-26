package initalizers

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"te-emb-api/models"

	"gorm.io/gorm"
)

func SyncDatabase() {
	// DB.AutoMigrate(&models.User{})
	DBL.AutoMigrate(&models.User{})

	cmd := exec.Command("cat", "/proc/sys/kernel/random/boot_id")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	var boot_id models.Boot_seq_record
	if result := DBL.Where("boot_id=?", strings.Trim(stdout.String(), "\n")).First(&boot_id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("fail to find ID: %s boot_id table %s", boot_id.Boot_Id, result.Error)
			store_path := os.Getenv("STORAGE")
			ams_sn := models.Boot_seq_record{
				Boot_Id: strings.Trim(stdout.String(), "\n"),
				Path:    store_path,
			}
			if res := DBL.Create(&ams_sn); res.Error != nil {
				log.Printf("fail to insert boot_id table %s", res.Error)
			}
		}
	}
}
