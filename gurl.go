package gurl

import (
	"net/http"
	"net/url"
)

var (
	Debug          = false
	ProgressPrefix = ""
)

// Interface to handle Gets
type Downloader interface {
	Get() error
}

// Returns a new Downloader. With the right Downloader interface for scheme.
func NewDownloader(destdir, rawurl string) Downloader {

	url, _ := url.Parse(rawurl)
	switch url.Scheme {
	case "http":
		return Http{destdir, rawurl}
	}
	return Unknown{}
}

// We should use an interface for downloading. This would allow for different protocols.
// eg. git:// ftp://

// Download a slice of URL's to destination directory
// TODO: make this concurrent
func DownloadAll(destdir string, rawurls []string) (err error) {
	/*for _, rawurl := range rawurls {
		err = Download(destdir, rawurl)
		if err != nil {
			return
		}
	}*/
	return
}

func Download(destdir, rawurl string) (err error) {
	return NewDownloader(destdir, rawurl).Get()
}

func buildRequest(method, rawurl string) (*http.Request, error) {
	var err error
	req := new(http.Request)
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	req.Header = http.Header{}
	req.Header.Set("Connection", "keep-alive")
	req.Method = method
	req.URL, err = url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	return req, nil
}
