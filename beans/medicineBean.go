package beans

import "time"

type Medicine struct {
	Id             int       `gorm:"primary_key"`
	TabletName     string    `json:"table_name"`
	TimeToRemind   time.Time `json:"time_to_remind"`
	DaysToRemind   time.Time `json:"days_to_remind"`
	MedicineType   string    `json:"medicine_type"`
	MedicineDosage string    `json:"medicine_dosage"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}
