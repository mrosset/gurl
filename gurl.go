package gurl

import (
	"http"
	"io"
	"os"
	"path"
)

type Client struct {
	client *http.Client
}


func (v *Client) Download(destdir string, url string) os.Error {
	if v.client == nil {
		v.client = new(http.Client)
	}
	r, finalurl, err := v.client.Get(url)
	if err != nil {
		return err
	}
	fpath := path.Join(destdir, path.Base(finalurl))
	f, err := os.Open(fpath, os.O_WRONLY|os.O_CREAT, 0644)
	defer f.Close()
	_, err = io.Copy(f, r.Body)
	if err != nil {
		return err
	}
	return err
}
