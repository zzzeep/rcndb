package storage

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GetAllDomains() ([]Domain, error) {
	db := open()
	var domains []Domain
	err := db.Find(&domains).Error
	if err != nil {
		return nil, err
	}

	return domains, nil
}

func GetDomain(domain string) (*Domain, error) {
	db := open()
	d := Domain{
		Domain: domain,
	}
	err := db.Where(&d).Error
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func GetDomainById(id uint) (*Domain, error) {
	db := open()
	d := &Domain{}
	err := db.First(d, id).Error
	if err != nil {
		return nil, err
	}
	return d, nil
}

func GetDomainUrls(d Domain) ([]URL, error) {
	db := open()
	var urls []URL
	err := db.Where(&URL{DomainID: d.ID}).Find(&urls).Error
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func GetAllUrls() ([]URL, error) {
	db := open()
	var urls []URL
	err := db.Joins("Domain").Find(&urls).Error
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func GetDomainChanges(d *Domain) ([]DomainChange, error) {
	db := open()
	err := db.Find(d).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("domain not found", d.Domain)
		return nil, err
	}

	var dChanges []DomainChange
	err = db.Joins("Domain").Where(&DomainChange{DomainID: d.ID}).Find(&dChanges).Error
	if err != nil {
		return nil, err
	}

	return dChanges, nil
}

func GetUrlChanges(d Domain, url string) ([]UrlChange, error) {
	db := open()
	u := URL{}
	err := db.Where(&URL{DomainID: d.ID}).Where(&URL{Path: url}).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("url not found", d.Domain+url)
		return nil, err
	}

	var uChanges []UrlChange
	err = db.Joins("Url").Where(&UrlChange{UrlID: u.ID}).Find(&uChanges).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return uChanges, nil
}
