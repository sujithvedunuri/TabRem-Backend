package beans

import (
	"gorm.io/gorm"
)

type Medicine struct {
	ID             uint   `gorm:"primaryKey"`
	TabletName     string `json:"tablet_name"`
	MedicineType   string `json:"medicine_type"`
	MedicineDosage string `json:"medicine_dosage"`
	// TimeToRemind   time.Time `json:"time_to_remind"`
	// DaysToRemind time.Time `json:"days_to_remind"`
	CreatedAt int64          `gorm:"autoCreateTime"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
