package worker

import (
	"errors"
	my "github.com/kaatinga/assets"
	"sync"
)

// модели данных для списка урлов для "обзвона"
type URL struct {
	Name string
	URL  string
	Body string
}

type URLs struct {
	mx   sync.RWMutex
	List []URL
}

func (dDoSURLs *URLs) GetURL() URL {
	dDoSURLs.mx.RLock()
	defer dDoSURLs.mx.RUnlock()

	return dDoSURLs.List[my.GetRandomByte(byte(len(dDoSURLs.List)))]
}

func (dDoSURLs *URLs) AddURL(name, url, body string) error {
	dDoSURLs.mx.Lock()
	defer dDoSURLs.mx.Unlock()

	if len(dDoSURLs.List) > 255 {
		return errors.New("the limit of URLs is reached")
	}

	if url[:4] != "http" {
		return errors.New("incorrect URL")
	}

	dDoSURLs.List = append(dDoSURLs.List, URL{Name: name, URL: url, Body: body})
	return nil
}
