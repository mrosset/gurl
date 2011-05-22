package gurl

import (
	"bufio"
	"fmt"
	"http"
	"os"
	"path"
	term "github.com/kless/go-term/term"
	"strings"
	"time"
)

var (
	printf  = fmt.Printf
	fprintf = fmt.Fprintf
	println = fmt.Println
	Debug   = false
	buf     = bufio.NewWriter(os.Stderr)
)

type Client struct {
	client         *http.Client
	ProgressHandle func(int64, int64, int64, string)
}

func (v *Client) Download(destdir string, url string) os.Error {
	if v.client == nil {
		v.client = new(http.Client)
	}
	if v.ProgressHandle == nil {
		v.ProgressHandle = doProgress
	}
	request, err := buildRequest("GET", url)
	if err != nil {
		return err
	}
	if Debug {
		b, err := http.DumpRequest(request, false)
		if err != nil {
			return err
		}
		os.Stderr.Write(b)
	}
	response, err := v.client.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return os.NewError("Error status " + response.Status)
	}
	defer response.Body.Close()
	if Debug {
		b, err := http.DumpResponse(response, false)
		if err != nil {
			return err
		}
		os.Stderr.Write(b)
	}
	fpath := path.Join(destdir, path.Base(url))
	f, err := os.Create(fpath)
	defer f.Close()
	var downloaded int64
	start := time.Seconds()
	tick := time.Tick(1e09)
	for {
		b := make([]byte, 1024)
		read, err := response.Body.Read(b)
		if err != nil && err != os.EOF {
			return err
		}
		downloaded += int64(read)
		select {
		case <-tick:
			v.ProgressHandle(start, downloaded, response.ContentLength, fpath)
		default:
		}
		if err == os.EOF {
			v.ProgressHandle(start, downloaded, response.ContentLength, fpath)
			break
		}
		_, err = f.Write(b[0:read])
		if err != nil {
			return err
		}
	}
	fmt.Println()
	return err
}

func doProgress(start, downloaded, totalDownload int64, file string) {
	winsize, _ := term.GetWinsize()
	var (
		width    int64 = int64(winsize.Col)
		percent  int64 = (downloaded * 100) / totalDownload
		progress int64 = (width * percent) / 100
	)
	bar := strings.Repeat("#", int(progress))
	pad := strings.Repeat(" ", int(width)-int(progress))
	fmt.Fprintf(buf, "\r%s%s", bar, pad)
	buf.Flush()
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
