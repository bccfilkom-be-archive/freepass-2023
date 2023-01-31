package Config

import (
	"fmt"
	"freepass/Model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(
		mysql.Open(
			fmt.Sprintf(
				"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASS"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_NAME"),
			),
		),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(
		&Model.Admin{},
		&Model.Class{},
		&Model.Course{},
		&Model.Student{},
		&Model.User{},
	); err != nil {
		log.Fatal(err.Error())
	}

	courses := map[string]Model.Course{
		"Pemrograman Dasar":               {Name: "Pemrograman Dasar", Credits: 4, Code: "CIF2007"},
		"Pemrograman Berorientasi Objek":  {Name: "Pemrograman Berorientasi Objek", Credits: 4, Code: "CIF2008"},
		"Keamanan Informasi":              {Name: "Keamanan Informasi", Credits: 4, Code: "CIF2009"},
		"Sistem Multimedia":               {Name: "Sistem Multimedia", Credits: 4, Code: "CIF2010"},
		"Pemrograman Web":                 {Name: "Pemrograman Web", Credits: 5, Code: "CIF2011"},
		"Analisis dan Perancangan Sistem": {Name: "Analisis dan Perancangan Sistem", Credits: 5, Code: "CIF2012"},
		"Pengantar Pembelajaran Mesin":    {Name: "Pengantar Pembelajaran Mesin", Credits: 10, Code: "CIF2013"},
		"Bahasa Indonesia":                {Name: "Bahasa Indonesia", Credits: 1, Code: "CIF2014"},
	}

	for _, course := range courses {
		db.Save(&course)
	}

	return db
}
