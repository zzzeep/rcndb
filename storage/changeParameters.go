package storage

import (
	"time"

	"github.com/zzzeep/rcndb/util"
)

type ChangeType string

const (
	Host          ChangeType = "Host"
	Port          ChangeType = "Port"
	WebServer     ChangeType = "WebServer"
	StatusCode    ChangeType = "Status Code"
	ContentLength ChangeType = "Content Length"
	ContentType   ChangeType = "Content Type"
	Title         ChangeType = "Title"
	BodyMD5       ChangeType = "Body MD5"
	HeaderMD5     ChangeType = "Header MD5"
)

type ChangeParameters struct {
	OldValue *string
	NewValue *string
	Type     ChangeType
	Date     time.Time
}

func ChangeParametersFromString(oldV *string, newV *string, t ChangeType) ChangeParameters {
	return ChangeParameters{
		OldValue: oldV,
		NewValue: newV,
		Type:     t,
		Date:     time.Now(),
	}
}

func ChangeParametersFromUint(oldV uint, newV uint, t ChangeType) ChangeParameters {
	oldVStr := util.Uint2Str(oldV)
	newVStr := util.Uint2Str(newV)
	return ChangeParametersFromString(&oldVStr, &newVStr, t)
}
