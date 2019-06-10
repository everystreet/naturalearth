package naturalearth

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/mercatormaps/go-shapefile"
)

type Source interface {
	io.Closer
	Open() (*shapefile.ZipScanner, error)
}

func NewSource(uri string, cli *http.Client) Source {
	url, err := url.Parse(uri)
	if err == nil && url.Scheme != "" {
		return &RemoteFile{
			client: http.DefaultClient,
			url:    url,
		}
	}

	return &LocalFile{
		path: uri,
	}
}

type RemoteFile struct {
	client *http.Client
	url    *url.URL
}

func (f *RemoteFile) Open() (*shapefile.ZipScanner, error) {
	resp, err := f.client.Get(f.url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return shapefile.NewZipScanner(bytes.NewReader(buf), resp.ContentLength, path.Base(f.url.Path))
}

func (f *RemoteFile) Close() error {
	return nil
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
