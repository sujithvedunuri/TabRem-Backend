package beans

import "gorm.io/gorm"

type Medicine struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	TabletName     string  `json:"tablet_name" gorm:"unique"`
	MedicineType   string `json:"medicine_type"`
	MedicineDosage int `json:"medicine_dosage"`
 // TimeToRemind   time.Time `json:"time_to_remind"`
	DaysToRemind int `json:"days_to_remind"`
	CreatedAt   int64          `gorm:"autoCreateTime"`
	UpdatedAt   int64          `gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
