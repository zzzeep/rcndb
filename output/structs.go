package output

type ShowOptions struct {
	Domains         bool
	IPs             bool
	Ports           bool
	Urls            bool
	Status          bool
	Webserver       bool
	Content         bool
	LastScan        bool
	NoColor         bool
	NoTruncation    bool
	NoEmptyResponse bool

	FilterStatus uint
	FilterIP     string
	FilterPort   string
}

func (opt ShowOptions) CheckUnset() ShowOptions {
	if !opt.Domains && !opt.Urls && !opt.Status && !opt.IPs &&
		!opt.Ports && !opt.Webserver && !opt.Content && !opt.LastScan {
		opt.Domains = true
		opt.Urls = true
		opt.Status = true
		opt.IPs = true
		opt.Ports = true
		opt.Webserver = true
		opt.Content = true
		opt.LastScan = true
	}
	return opt
}

type TrackOptions struct {
	Domain  string
	Url     string
	All     bool
	NoColor bool
}
