package storage

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/zzzeep/rcndb/parser"
	"gorm.io/gorm"
)

func ImportHttpxRecords(records []parser.HttpxRecord) {
	db := open()
	for _, rec := range records {
		u, err := url.Parse(rec.Url)
		if err != nil {
			fmt.Println("couldn't parse url:", rec.Url)
			return
		}

		d := Domain{
			Domain:     u.Host,
			Host:       &rec.Host,
			Port:       &rec.Port,
			Webserver:  &rec.Webserver,
			LastChange: time.Now(),
		}
		err = db.Create(&d).Error

		if IsUniqueContraintError(err) {
			db.Where(&Domain{Domain: d.Domain}).First(&d)

			HandleDomainUpdate(&d, rec)
		}

		existing := URL{}
		err = db.Where(&URL{DomainID: d.ID, Path: rec.Path}).First(&existing).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			lastScanTime, _ := time.Parse("2006-01-02", strings.Split(rec.Timestamp, "T")[0])

			db.Create(&URL{
				Domain:        d,
				Path:          rec.Path,
				StatusCode:    rec.StatusCode,
				ContentType:   &rec.ContentType,
				ContentLength: rec.ContentLength,
				Title:         &rec.Title,
				BodyMD5:       &rec.Hash.Body_md5,
				HeaderMD5:     &rec.Hash.Header_md5,
				ResponseTime:  rec.Time,
				LastScan:      lastScanTime,
			})
		} else {
			HandleUrlUpdate(&existing, rec)
		}
	}
}

func ImportDomainList(domains []string) {
	db := open()
	for _, domain := range domains {
		if len(domain) == 0 {
			continue
		}

		if !strings.HasPrefix(domain, "http") {
			domain = "https://" + domain
		}

		u, _ := url.Parse(domain)

		d := Domain{
			Domain:     u.Host,
			LastChange: time.Now(),
		}
		err := db.Create(&d).Error

		if IsUniqueContraintError(err) {
			db.Where(&Domain{Domain: d.Domain}).First(&d)
		}

		path := u.Path + u.RawQuery
		if len(path) == 0 {
			path = "/"
		}

		existing := URL{}
		err = db.Where(&URL{DomainID: d.ID}).Where(&URL{Path: path}).First(&existing).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&URL{
				DomainID: d.ID,
				Path:     path,
			})
		}
	}
}

func IsUniqueContraintError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "UNIQUE constraint failed")
}
