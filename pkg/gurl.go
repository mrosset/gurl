package gurl

import (
	"bufio"
	"fmt"
	"http"
	"os"
	"path"
	term "github.com/kless/go-term/term"
	"strconv"
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

func (v *Client) Download(destdir string, url string) (err os.Error) {
	defer func() {
		if recover() != nil {
		}
	}()
	if v.client == nil {
		v.client = new(http.Client)
	}
	if v.ProgressHandle == nil {
		v.ProgressHandle = doProgress
	}
	req, err := buildRequest("GET", url)
	if err != nil {
		return err
	}
	debugRequest(req)
	res, err := v.client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return os.NewError("Error status " + res.Status)
	}
	defer res.Body.Close()
	debugResponse(res)
	fpath := path.Join(destdir, path.Base(url))
	f, err := os.Create(fpath)
	defer f.Close()
	var downloaded int64
	start := time.Seconds()
	tick := time.Tick(1e09)
	for {
		b := make([]byte, 1024)
		read, err := res.Body.Read(b)
		if err != nil && err != os.EOF {
			return err
		}
		downloaded += int64(read)
		select {
		case <-tick:
			if res.ContentLength > 0 {
				v.ProgressHandle(start, downloaded, res.ContentLength, fpath)
			}
		default:
		}
		if err == os.EOF {
			if res.ContentLength > 0 {
				v.ProgressHandle(start, downloaded, res.ContentLength, fpath)
			} else {
				fmt.Printf("%v done.\n", fpath)
			}
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
		width    int = int((int64(winsize.Col) / 2)) - 9
		percent  int = int((downloaded * 100) / totalDownload)
		progress int = int((width * percent) / 100)
		bps      int
	)
	tick := time.Seconds() - start
	if tick == 0 {
		tick++
	}
	bps = int(downloaded / tick)
	bar := strings.Repeat("#", progress)
	stats := fmt.Sprintf("%3.3s%% %9.9s", strconv.Itoa(percent), speed(bps))
	fmt.Fprintf(buf, "\r%-*.*s [%-*s] %s", width, width, file, width, bar, stats)
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

func speed(bint int) string {
	var (
		b float32 = float32(bint)
	)
	switch {
	case b < 1024:
		return fmt.Sprintf("%vB/s", b)
	case b < 1024*1000:
		return fmt.Sprintf("%5.1fKB/s", b/1024)
	case b < 1024*1024*1000:
		return fmt.Sprintf("%5.1fMB/s", b/1024/1024)
	default:
		return fmt.Sprintf("%5.1fGB/s", b/1024/1024/1024)
	}
	return fmt.Sprintf("%5.1fGB/s", b/1024/1024/1024)
}

func debugRequest(req *http.Request) {
	if Debug {
		b, _ := http.DumpRequest(req, false)
		os.Stderr.Write(b)
	}
}

func debugResponse(res *http.Response) {
	if Debug {
		b, _ := http.DumpResponse(res, false)
		os.Stderr.Write(b)
	}
}
