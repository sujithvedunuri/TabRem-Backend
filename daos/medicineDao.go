package daos

import (
	"sujith/tabRemBackend/beans"
	"sujith/tabRemBackend/resources/database"
)

func FetchMedicineDetails() []beans.Medicine {
	var medicines []beans.Medicine
	database.Db.Find(&medicines)
	return medicines
}

func FetchMedicineById() beans.Medicine {
	var medicines beans.Medicine
	database.Db.Where("id=?").First(&medicines)
	return medicines
}

func AddMedicineDetials() {
	//add details to database
}
