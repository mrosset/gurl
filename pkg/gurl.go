package gurl

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"util/console"
	"util/file"
)

var (
	Debug  = false
	client = new(http.Client)
)

func Download(rawurl, destdir string) (err error) {
	if !file.Exists(destdir) {
		return fmt.Errorf("dir %s does not exists.", destdir)
	}
	req, err := buildRequest("GET", rawurl)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("Error status " + res.Status)
	}
	defer res.Body.Close()
	fpath := path.Join(destdir, path.Base(rawurl))
	fd, err := os.Create(fpath)
	defer fd.Close()
	pw := console.NewProgressBarWriter(path.Base(rawurl), res.ContentLength, fd)
	_, err = io.Copy(pw, res.Body)
	return err
}

/*
func doProgressBar(start time.Time, downloaded, totalDownload int64, file string) {
	twidth, err := TermWidth()
	if err != nil {
		log.Fatal(err)
	}
	var (
		width    int = int((int64(twidth) / 2)) - 9
		percent  int = int((downloaded * 100) / totalDownload)
		progress int = int((width * percent) / 100)
		bps      int
	)
	tick := time.Now().Sub(start)
	if tick == 0 {
		tick++
	}
	bps = int(downloaded / int64(tick.Seconds()))
	bar := strings.Repeat("#", progress)
	stats := fmt.Sprintf("%3.3s%% %9.9s", strconv.Itoa(percent), speed(bps))
	fmt.Fprintf(buf, "\r%-*.*s [%-*s] %s", width, width, file, width, bar, stats)
	buf.Flush()
}

func doProgress(start time.Time, downloaded, totalDownload int64, file string) {
	var (
		percent int = int((downloaded * 100) / totalDownload)
		bps     int
	)
	tick := time.Now().Sub(start)
	if tick == 0 {
		tick++
	}
	bps = int(downloaded / int64(tick.Seconds()))
	fmt.Fprintf(buf, "\r%-40.40s %3.3s%% %v", file, strconv.Itoa(percent), speed(bps))
	buf.Flush()
}
*/
// Receiving objects:   2% (41013/2050606), 14.90 MiB | 1.03 MiB/s

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
