package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"gitlab.com/syaifuddin.teddy/test-case-majoo/api/models"
)

var users = []models.User{
	models.User{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: "password",
	},
}

var products = []models.Product{
	models.Product{
		Name:        "Product 1",
		Description: "Just A Sample Product",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title Tes",
		Content: "Hello world",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.Product{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Product{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Product{}).Create(&products[i]).Error
		if err != nil {
			log.Fatalf("cannot seed products table: %v", err)
		}

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
