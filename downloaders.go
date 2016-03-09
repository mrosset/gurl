package gurl

import (
	"errors"
	"fmt"
	"github.com/str1ngs/util/console"
	"github.com/str1ngs/util/file"
	"io"
	"net/http"
	"os"
	"path"
)

var (
	client = new(http.Client)
)

// Provides an Unknown Downloader
type Unknown struct {
}

func (u Unknown) Get() error {
	return errors.New("Unknown URL scheme")
}

// Provides Downloader interface for http and https
type Http struct {
	destdir string
	rawurl  string
}

// Provides Downloader Get method
// TODO: simplify this method
func (h Http) Get() error {
	if !file.Exists(h.destdir) {
		return fmt.Errorf("dir %s does not exists.", h.destdir)
	}
	// building a request every time is annoying. Why even use this
	req, err := buildRequest("GET", h.rawurl)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Error status %s %s", res.Status, h.rawurl)
	}
	defer res.Body.Close()
	fpath := path.Join(h.destdir, path.Base(h.rawurl))
	fd, err := os.Create(fpath)
	defer fd.Close()
	prefix := path.Base(h.rawurl)
	if ProgressPrefix != "" {
		prefix = ProgressPrefix
	}
	pw := console.NewProgressBarWriter(prefix, res.ContentLength, fd)
	defer pw.Close()
	_, err = io.Copy(pw, res.Body)
	return nil
}
