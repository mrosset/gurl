package gurl

import (
	"fmt"
	"http"
	"os"
	"path"
	//term "github.com/kless/go-term/term"
	"syscall"
	"time"
	"unsafe"
)

var (
	printf  = fmt.Printf
	fprintf = fmt.Fprintf
	println = fmt.Println
	debug   = false
)

const (
	TIOCGWINSZ = 0x5413
)

type Client struct {
	client *http.Client
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWinSize() (*winsize, os.Error) {
	ws := new(winsize)
	r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(r1) == -1 {
		return nil, os.NewSyscallError("getWinSize", int(errno))
	}
	return ws, nil
}

func (v *Client) Download(destdir string, url string) os.Error {
	winsize, err := getWinSize()
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
	if debug {
		b, err := http.DumpRequest(request, false)
		if err != nil {
			return err
		}
		os.Stderr.Write(b)
	}
	response, err := v.client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return err
	}
	if debug {
		b, err := http.DumpResponse(response, false)
		if err != nil {
			return err
		}
		os.Stderr.Write(b)
	}
	fpath := path.Join(destdir, path.Base(url))
	f, err := os.Create(fpath)
	defer f.Close()
	_ = response.ContentLength / int64(winsize.Col)
	var tdown int64
	start := time.Seconds()
	start = start
	println("started", start)
	//start := time.Seconds()
	tick := time.NewTicker(1e09)
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
		tdown += int64(read)
		frac := float32(tdown) / float32(response.ContentLength)
		percent := tdown * 100 / response.ContentLength
		spercent := fmt.Sprintf("\r%v%% ", percent)
		line := spercent
		progress := int(frac*float32(winsize.Col)) - len(line) + 1
		for i := 0; i < progress; i++ {
			line += "#"
		}
		l := line
		for i := 0; i < int(winsize.Col)-len(l); i++ {
			line += " "
		}
		fprintf(os.Stderr, "%v", line)
		go func() {
			select {
			case <-tick.C:
				println("tick")
			}
		}()
	}
	tick.Stop()
	println()
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
