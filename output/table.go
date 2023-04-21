package output

import (
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/zzzeep/rcndb/storage"
	"github.com/zzzeep/rcndb/util"
)

func SimplePrintOut(opt ShowOptions) {
	opt = opt.CheckUnset()
	if opt.NoColor {
		text.DisableColors()
	}

	urls, err := storage.GetAllUrls()
	if err != nil {
		panic("Couldn't get urls")
	}

	tbl := table.NewWriter()
	tbl.Style().Options.DrawBorder = false
	tbl.Style().Options.SeparateColumns = false
	tbl.SetOutputMirror(os.Stdout)

	for _, u := range urls {
		if opt.NoEmptyResponse && u.StatusCode == 0 {
			continue
		}

		if opt.FilterStatus != 0 && u.StatusCode != opt.FilterStatus {
			continue
		}
		if len(opt.FilterIP) > 0 && !strings.Contains(util.Ptr2Str(u.Domain.Host), opt.FilterIP) {
			continue
		}
		if len(opt.FilterPort) > 0 && util.Ptr2Str(u.Domain.Port) != opt.FilterPort {
			continue
		}

		r := table.Row{}
		if opt.Domains || opt.Urls {
			r = append(r, styleUrl(opt, u))
		}
		if opt.Status {
			r = append(r, styleStatusCode(u.StatusCode))
		}
		if opt.IPs {
			r = append(r, styleIP(util.Ptr2Str(u.Domain.Host)))
		}
		if opt.Ports {
			r = append(r, util.Ptr2Str(u.Domain.Port))
		}
		if opt.Webserver {
			r = append(r, styleWebserver(u.Domain.Webserver))
		}
		if opt.Content {
			r = append(r, util.Ptr2Str(u.ContentType))
			r = append(r, text.FgMagenta.Sprint(u.ContentLength))
		}
		if opt.LastScan {
			r = append(r, styleDate(u.LastScan))
		}

		tbl.AppendRow(r)
	}

	tbl.Render()
}
