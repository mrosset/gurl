package gurl

import (
	"fmt"
	"github.com/str1ngs/util/console"
	"github.com/str1ngs/util/file"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

var (
	Debug          = false
	client         = new(http.Client)
	ProgressPrefix = ""
)

// Interface to handle Gets
type Downloader interface {
	Get(out string, url string) error
}

// Provides Downloader interface for http and https
type Http struct {
}

// Provides Downloader Get method
func (h Http) Get(p string, url string) error {
	return nil
}

// Download a slice of URL's to destination directory
// TODO: make this concurrent
func DownloadAll(destdir string, rawurls []string) (err error) {
	for _, rawurl := range rawurls {
		err = Download(destdir, rawurl)
		if err != nil {
			return
		}
	}
	return
}

// Returns a new Downloader. With the right Downloader interface for scheme.
// TODO: this breaks Download singleton. however we could use NewDownloader(url).Get()
func NewDownloader(url string) {
}

// We should use an interface for downloading. This would allow for different protocols.
// eg. git:// ftp://

// Download url to destdir
func OldDownload(destdir, rawurl string) (err error) {
	if !file.Exists(destdir) {
		return fmt.Errorf("dir %s does not exists.", destdir)
	}
	// building a request every time is annoying. Why even use this
	req, err := buildRequest("GET", rawurl)
	if err != nil {
		return
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("Error status %s %s", res.Status, rawurl)
	}
	defer res.Body.Close()
	fpath := path.Join(destdir, path.Base(rawurl))
	fd, err := os.Create(fpath)
	defer fd.Close()
	prefix := path.Base(rawurl)
	if ProgressPrefix != "" {
		prefix = ProgressPrefix
	}
	pw := console.NewProgressBarWriter(prefix, res.ContentLength, fd)
	defer pw.Close()
	_, err = io.Copy(pw, res.Body)
	return
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
