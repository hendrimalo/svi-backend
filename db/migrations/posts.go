package migrations

import "time"

type Post struct {
	Id        int        `gorm:"primaryKey;autoIncrement;not null"`
	Title     string     `gorm:"type:varchar(200)"`
	Content   string     `gorm:"type:text"`
	Category  string     `gorm:"type:varchar(100)"`
	CreatedAt *time.Time `gorm:"type:timestamp"`
	UpdatedAt *time.Time `gorm:"type:timestamp"`
	Status    string     `gorm:"type:enum('publish', 'draft', 'trash')"`
}
