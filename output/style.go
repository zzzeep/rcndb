package output

import (
	"net"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/zzzeep/rcndb/storage"
	"github.com/zzzeep/rcndb/util"
)

func styleDate(date time.Time) string {
	if (date == time.Time{}) {
		return ""
	}
	return date.Local().Format("2006-01-02")
}

func styleStatusCode(status uint) string {
	if status == 0 {
		return ""
	}

	if status >= 200 && status < 300 {
		return text.FgGreen.Sprint(status)
	} else if status >= 300 && status < 400 {
		return text.FgYellow.Sprint(status)
	} else if status >= 400 && status < 600 {
		return text.FgRed.Sprint(status)
	} else {
		return ""
	}
}

func styleIP(ipString string) string {
	ip := net.ParseIP(ipString)
	private := ip.IsPrivate()

	if private {
		return text.FgHiRed.Sprint(ipString)
	} else {
		return text.FgBlue.Sprint(ipString)
	}
}

func styleUrl(opt ShowOptions, u storage.URL) string {
	res := ""
	if opt.Domains {
		res += u.Domain.Domain
	}
	if opt.Urls {
		res += u.Path
	}

	truncLenght := 40
	if len(res) > truncLenght && !opt.NoTruncation {
		res = res[0:truncLenght]
		res += "..."
	}
	return text.FgHiWhite.Sprint(res)
}

func styleWebserver(ws *string) string {
	return text.FgHiCyan.Sprint(util.Ptr2Str(ws))
}
