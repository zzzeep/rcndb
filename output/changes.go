package output

import (
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/zzzeep/rcndb/storage"
	"github.com/zzzeep/rcndb/util"
)

var tbl table.Writer

func InitTable(noColor bool) {
	tbl = table.NewWriter()
	if noColor {
		text.DisableColors()
	}
	tbl.Style().Options.DrawBorder = false
	tbl.Style().Options.SeparateColumns = false
	tbl.SetOutputMirror(os.Stdout)
}

func PrintDomainChanges(dChanges []storage.DomainChange) {
	for _, d := range dChanges {
		r := table.Row{}
		r = append(r, text.FgWhite.Sprint(d.Domain.Domain))
		r = appendChanges(r, d.Change)
		tbl.AppendRow(r)
	}
}

func PrintURLChanges(uChanges []storage.UrlChange) {
	for _, u := range uChanges {
		d, err := storage.GetDomainById(u.Url.DomainID)
		if err != nil {
			continue
		}

		r := table.Row{}
		r = append(r, text.FgWhite.Sprint(d.Domain+u.Url.Path))
		r = appendChanges(r, u.Change)
		tbl.AppendRow(r)
	}
}

func appendChanges(r table.Row, ch storage.ChangeParameters) table.Row {
	fmtOldValue := formatValue(ch.OldValue, ch.Type)
	fmtNewValue := formatValue(ch.NewValue, ch.Type)

	r = append(r, fmtOldValue)
	r = append(r, ">")
	r = append(r, fmtNewValue)
	r = append(r, ch.Type)
	r = append(r, ch.Date.Format("2006-01-02"))
	return r
}

func Render() {
	tbl.Render()
}

func formatValue(value *string, t storage.ChangeType) string {
	switch t {
	case storage.Host:
		return styleIP(util.Ptr2Str(value))
	case storage.WebServer:
		return styleWebserver(value)
	case storage.ContentLength:
		return text.FgMagenta.Sprint(util.Ptr2Str(value))
	case storage.StatusCode:
		number, err := strconv.ParseUint(util.Ptr2Str(value), 10, 64)
		if err != nil {
			number = 0
		}
		return styleStatusCode(uint(number))
	default:
		return util.Ptr2Str(value)
	}
}
