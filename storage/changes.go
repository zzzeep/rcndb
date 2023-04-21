package storage

import (
	"strings"
	"time"

	"github.com/zzzeep/rcndb/parser"
	"github.com/zzzeep/rcndb/util"
)

func HandleUrlUpdate(u *URL, rec parser.HttpxRecord) {
	db := open()

	lastScanTime, _ := time.Parse("2006-01-02", strings.Split(rec.Timestamp, "T")[0])
	if u.LastScan.Before(lastScanTime) {
		u.LastScan = lastScanTime
		u.ResponseTime = rec.Time

		if NeedsChange(u.StatusCode, rec.StatusCode) {
			ch := ChangeParametersFromUint(u.StatusCode, rec.StatusCode, StatusCode)
			StoreUrlChange(*u, ch)
			u.StatusCode = rec.StatusCode
		}

		if NeedsChange(u.ContentLength, rec.ContentLength) {
			ch := ChangeParametersFromUint(u.ContentLength, rec.ContentLength, ContentLength)
			StoreUrlChange(*u, ch)
			u.ContentLength = rec.ContentLength
		}

		if NeedsChange(util.Ptr2Str(u.ContentType), rec.ContentType) {
			ch := ChangeParametersFromString(u.ContentType, &rec.ContentType, ContentType)
			StoreUrlChange(*u, ch)
			u.ContentType = &rec.ContentType
		}

		if NeedsChange(util.Ptr2Str(u.Title), rec.Title) {
			ch := ChangeParametersFromString(u.Title, &rec.Title, Title)
			StoreUrlChange(*u, ch)
			u.Title = &rec.Title
		}

		if NeedsChange(util.Ptr2Str(u.BodyMD5), rec.Hash.Body_md5) {
			ch := ChangeParametersFromString(u.BodyMD5, &rec.Hash.Body_md5, BodyMD5)
			StoreUrlChange(*u, ch)
			u.BodyMD5 = &rec.Hash.Body_md5
		}

		if NeedsChange(util.Ptr2Str(u.HeaderMD5), rec.Hash.Header_md5) {
			ch := ChangeParametersFromString(u.HeaderMD5, &rec.Hash.Header_md5, BodyMD5)
			StoreUrlChange(*u, ch)
			u.HeaderMD5 = &rec.Hash.Header_md5
		}
	}
	db.Save(u)
}

func HandleDomainUpdate(d *Domain, rec parser.HttpxRecord) {
	db := open()
	if NeedsChange(util.Ptr2Str(d.Host), rec.Host) {
		ch := ChangeParametersFromString(d.Host, &rec.Host, Host)
		StoreDomainChange(*d, ch)
		d.Host = &rec.Host
	}

	if NeedsChange(util.Ptr2Str(d.Port), rec.Port) {
		ch := ChangeParametersFromString(d.Port, &rec.Port, Port)
		StoreDomainChange(*d, ch)
		d.Port = &rec.Port
	}

	if NeedsChange(util.Ptr2Str(d.Webserver), rec.Webserver) {
		ch := ChangeParametersFromString(d.Webserver, &rec.Webserver, WebServer)
		StoreDomainChange(*d, ch)
		d.Webserver = &rec.Webserver
	}
	db.Save(d)
}

func StoreDomainChange(d Domain, ch ChangeParameters) {
	db := open()

	db.Create(&DomainChange{
		Domain: d,
		Change: ch,
	})
}

func StoreUrlChange(u URL, ch ChangeParameters) {
	db := open()

	db.Create(&UrlChange{
		Url:    u,
		Change: ch,
	})
}

func NeedsChange(oldV any, newV any) bool {
	if newV != nil && oldV != newV {
		return true
	}

	return false
}
