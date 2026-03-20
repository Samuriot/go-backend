package models

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type Service struct {
	URL          *url.URL
	description  string
	Alive        bool
	Mutex        sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

type ServicePool struct {
	Services []*Service
	current  uint64
}
