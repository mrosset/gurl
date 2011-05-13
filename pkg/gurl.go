package gurl

import (
	"fmt"
	"http"
	"os"
	"path"
	term "github.com/kless/go-term/term"
)

var (
	printf  = fmt.Printf
	println = fmt.Println
)

type Client struct {
	client *http.Client
}

func (v *Client) Download(destdir string, url string) os.Error {
	winsize, err := term.GetWinsize()
	if err != nil {
		return err
	}
	if v.client == nil {
		v.client = new(http.Client)
	}
	request, err := buildRequest("GET", url)
	if err != nil {
		return err
	}
	response, err := v.client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return err
	}
	fpath := path.Join(destdir, path.Base(url))
	f, err := os.Create(fpath)
	defer f.Close()
	segments := response.ContentLength / (int64(winsize.Col) - 4)

	var bcount int64
	for {
		b := make([]byte, 1024)
		read, err := response.Body.Read(b)
		if err != nil && err != os.EOF {
			return err
		}
		if err == os.EOF {
			break
		}
		_, err = f.Write(b[0:read])
		if err != nil {
			return err
		}
		if bcount >= segments {
			printf("#")
			bcount = 0
			continue
		}
		bcount += int64(read)
	}
	println()
	//_, err = io.Copy(f, response.Body)
	return err
}

func buildRequest(method, url string) (*http.Request, os.Error) {
	var err os.Error
	req := new(http.Request)
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	req.Header = http.Header{}
	req.Header.Set("Connection", "keep-alive")
	req.Method = method
	req.URL, err = http.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return req, nil
}
