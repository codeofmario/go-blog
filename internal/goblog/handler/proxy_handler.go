package handler

import (
	"github.com/gin-gonic/gin"
	"goblog.com/goblog/internal/goblog/config"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ProxyHandler interface {
	ServePublicBucket(c *gin.Context)
}

type ProxyHandlerImpl struct {
	settings *config.Settings
}

func NewProxyHandler(settings *config.Settings) ProxyHandler {
	return &ProxyHandlerImpl{settings: settings}
}

func (h *ProxyHandlerImpl) ServePublicBucket(c *gin.Context) {
	remote, err := url.Parse(h.settings.MinioPublicAddr)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = "/" + h.settings.PublicBucketName + "/" + c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
