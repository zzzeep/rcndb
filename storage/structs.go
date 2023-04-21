package storage

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Domain     string `gorm:"unique;not null"`
	Host       *string
	Port       *string
	Webserver  *string
	LastChange time.Time
}

type URL struct {
	gorm.Model
	DomainID      uint
	Domain        Domain
	Path          string `gorm:"not null"`
	StatusCode    uint
	ContentType   *string
	ContentLength uint
	Title         *string
	BodyMD5       *string
	HeaderMD5     *string
	LastScan      time.Time
	ResponseTime  string
}

type DomainChange struct {
	DomainID uint
	Domain   Domain
	Change   ChangeParameters `gorm:"embedded"`
}

type UrlChange struct {
	UrlID  uint
	Url    URL
	Change ChangeParameters `gorm:"embedded"`
}
