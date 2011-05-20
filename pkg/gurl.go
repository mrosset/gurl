package gurl

import (
	"bufio"
	"bytes"
	"fmt"
	"http"
	"os"
	"path"
	term "github.com/kless/go-term/term"
	"io"
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
	progressHandle func(int64, int64, int64, chan int)
}

func (v *Client) Download(destdir string, url string) os.Error {
	if v.client == nil {
		v.client = new(http.Client)
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
			doProgress(start, downloaded, response.ContentLength, fpath)
		default:
		}
		if err == os.EOF {
			doProgress(start, downloaded, response.ContentLength, fpath)
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
	var bps int64
	tick := time.Seconds() - start
	if tick != 0 {
		bps = downloaded / tick
	}
	frac := float32(downloaded) / float32(totalDownload)
	percent := downloaded * 100 / totalDownload
	tail := fmt.Sprintf("] %vKB/s %v%% %v ", bps/1024, percent, file)
	line := new(bytes.Buffer)
	line.WriteString("\r[")
	progress := int(frac * (float32(winsize.Col) - float32(len(tail)) - 1))
	for i := 0; i < progress; i++ {
		line.WriteByte('#')
	}
	llen := len(line.Bytes())
	for i := 0; i < int(winsize.Col)-len(tail)-llen; i++ {
		line.WriteByte(' ')
	}
	line.WriteString(tail)
	io.Copy(buf, line)
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
