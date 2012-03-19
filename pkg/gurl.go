package gurl

import (
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
	Debug          = false
	client         = new(http.Client)
	ProgressPrefix = ""
)

func Download(rawurl, destdir string) (err error) {
	if !file.Exists(destdir) {
		return fmt.Errorf("dir %s does not exists.", destdir)
	}
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
	_, err = io.Copy(pw, res.Body)
	return
}

/*
// Receiving objects:   2% (41013/2050606), 14.90 MiB | 1.03 MiB/s
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
