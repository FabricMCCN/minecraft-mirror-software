package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/FabricMCCN/minecraft-mirror-software/logger"
)

type DownloadInfo struct {
	DownloadUrl     string
	DownloadStorage string
	Hash            string
	Size            int
}

type DownloadUtil struct {
}

func (d DownloadUtil) Download(url string, size int, storePath string) (string, error) {
	if s, p := d.download(url, storePath); s {
		return p.Name(), nil
	}
	return "", fmt.Errorf("download failed: %s", url)
}

func (d DownloadUtil) DownloadWithHashCheck(url string, size int, storePath string, hash string) (string, error) {
	f, err := d.Download(url, size, storePath)
	if err != nil {
		return "", err
	}

	ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}

	return f, nil
}

func (d DownloadUtil) download(url string, storePath string) (status bool, file *os.File) {
	var f *os.File
	var err error

	if storePath == "" {
		f, err = ioutil.TempFile("", "*")
	} else {
		f, err = os.Open(storePath)
	}

	defer func() {
		e := recover()
		if e != nil {
			logger.Error(fmt.Errorf("%w", e))
		}
	}()
	if err != nil {
		panic(err)
	}

	client := new(http.Client)
	client.Timeout = time.Second * 15

	var content []byte
	var resp *http.Response
	content, _ = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(fmt.Errorf("%w", err))
		return false, f
	}
	f.Write(content)

	return true, f
}
