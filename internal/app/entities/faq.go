package entities

import "time"

type Faq struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	UniqHash  string    `gorm:"unique" json:"uniq_hash"`
	Question  string    `gorm:"not null" json:"question"`
	Answer    string    `gorm:"not null" json:"answer"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	//DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Faqs []Faq
