package naturalearth

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/everystreet/go-shapefile"
	"github.com/gosuri/uiprogress"
)

type File interface {
	io.Closer
	Open() (*shapefile.ZipScanner, error)
}

func NewFile(uri, label string, cli *http.Client) File {
	url, err := url.Parse(uri)
	if err == nil && url.Scheme != "" {
		return &RemoteFile{
			client: http.DefaultClient,
			url:    url,
			label:  label,
		}
	}

	return &LocalFile{
		path: uri,
	}
}

type RemoteFile struct {
	client *http.Client
	url    *url.URL
	label  string
}

func (f *RemoteFile) Open() (*shapefile.ZipScanner, error) {
	resp, err := f.client.Get(f.url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	size := int(f.size())
	bar := uiprogress.AddBar(size)

	bar.AppendCompleted()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return f.label
	})

	buf := bytes.Buffer{}
	done := make(chan struct{})
	defer close(done)

	go func() {
		tick := time.NewTicker(10 * time.Millisecond)
		defer tick.Stop()

		for {
			select {
			case <-done:
				return
			case <-tick.C:
				bar.Set(buf.Len())
			}
		}
	}()

	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, err
	}
	bar.Set(size)
	return shapefile.NewZipScanner(bytes.NewReader(buf.Bytes()), resp.ContentLength, path.Base(f.url.Path))
}

func (f *RemoteFile) Close() error {
	return nil
}

func (f *RemoteFile) size() uint {
	resp, err := f.client.Head(f.url.String())
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	size, err := strconv.ParseUint(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return 0
	}
	return uint(size)
}

type LocalFile struct {
	path string
	file *os.File
}

func (f *LocalFile) Open() (*shapefile.ZipScanner, error) {
	r, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}

	stat, err := r.Stat()
	if err != nil {
		return nil, err
	}

	return shapefile.NewZipScanner(r, stat.Size(), filepath.Base(f.path))
}

func (f *LocalFile) Close() error {
	if f.file == nil {
		return nil
	}
	return f.file.Close()
}
