package proxy

import (
//	"net/http"
//	"github.com/go-zoo/bone"
//	"mittwald.de/servicegateway/config"
//	"github.com/mailgun/oxy/testutils"
//	consul "github.com/hashicorp/consul/api"
//	"strings"
//	"regexp"
//	"fmt"
//	"mittwald.de/servicegateway/auth"
)

//type ProxyBuilder struct {
//	Configuration *config.Configuration
//	ProxyHandler *ProxyHandler
//	Cache *Cache
//	RateThrottler *RateThrottler
//	AuthDecorator auth.AuthDecorator
//	Consul *consul.Client
//}
//
//func NewProxyBuilder(cfg *config.Configuration, proxy *ProxyHandler, cache *Cache, throttler *RateThrottler, authDecorator auth.AuthDecorator) (*ProxyBuilder, error) {
//	consulConfig := consul.DefaultConfig()
//	consulConfig.Address = fmt.Sprintf("%s:%d", cfg.Consul.Host, cfg.Consul.Port)
//
//	consul, err := consul.NewClient(consulConfig)
//	if err != nil {
//		return nil, err
//	}
//
//	return &ProxyBuilder{
//		Configuration: cfg,
//		ProxyHandler: proxy,
//		Cache: cache,
//		RateThrottler: throttler,
//		AuthDecorator: authDecorator,
//		Consul: consul,
//	}, nil
//}
//
//func (b *ProxyBuilder) BuildHandler(mux *bone.Mux, name string, appCfg config.Application) error {
//	routes := make(map[string]http.HandlerFunc)
//
//	backendUrl := appCfg.Backend.Url
//	if backendUrl == "" && appCfg.Backend.Service != "" {
//		if appCfg.Backend.Tag != "" {
//			backendUrl = fmt.Sprintf("http://%s.%s.service.consul", appCfg.Backend.Tag, appCfg.Backend.Service)
//		} else {
//			backendUrl = fmt.Sprintf("http://%s.service.consul", appCfg.Backend.Service)
//		}
//	}
//
//	if appCfg.Routing.Type == "path" {
//		var handler http.HandlerFunc = func(rw http.ResponseWriter, req *http.Request) {
//			sanitizedPath := strings.Replace(req.URL.Path, appCfg.Routing.Path, "", 1)
//			proxyUrl := backendUrl + sanitizedPath
//			b.ProxyHandler.HandleProxyRequest(rw, req, proxyUrl, name, &appCfg)
//		}
//
//		routes[appCfg.Routing.Path+"/*"] = handler
//	} else if appCfg.Routing.Type == "pattern" {
//		re := regexp.MustCompile(":([a-zA-Z0-9]+)")
//		for pattern, target := range appCfg.Routing.Patterns {
//			parameters := re.FindAllStringSubmatch(pattern, -1)
//			var patternHandler http.HandlerFunc = func(rw http.ResponseWriter, req *http.Request) {
//				targetUrl := backendUrl + target
//				for _, paramName := range parameters {
//					targetUrl = strings.Replace(targetUrl, paramName[0], bone.GetValue(req, paramName[1]), -1)
//				}
//
//				req.URL = testutils.ParseURI(targetUrl)
//				b.ProxyHandler.HandleProxyRequest(rw, req, targetUrl, name, &appCfg)
//			}
//
//			routes[pattern] = patternHandler
//		}
//	}
//
//	if err := b.AuthDecorator.RegisterRoutes(mux); err != nil {
//		return err
//	}
//
//	for route, handler := range routes {
//		safeHandler := handler
//		unsafeHandler := handler
//
//		if appCfg.Caching.Enabled {
//			safeHandler = b.Cache.DecorateHandler(handler)
//
//			if appCfg.Caching.AutoFlush {
//				unsafeHandler = b.Cache.DecorateUnsafeHandler(handler)
//			}
//		}
//
//		if ! appCfg.Auth.Disable {
//			safeHandler = b.AuthDecorator.DecorateHandler(safeHandler, &appCfg)
//			unsafeHandler = b.AuthDecorator.DecorateHandler(unsafeHandler, &appCfg)
//		}
//
//		if appCfg.RateLimiting {
//			safeHandler = b.RateThrottler.DecorateHandler(safeHandler)
//			unsafeHandler = b.RateThrottler.DecorateHandler(unsafeHandler)
//		}
//
//		mux.GetFunc(route, safeHandler)
//		mux.HeadFunc(route, safeHandler)
//		mux.OptionsFunc(route, safeHandler)
//
//		mux.PostFunc(route, unsafeHandler)
//		mux.PutFunc(route, unsafeHandler)
//		mux.DeleteFunc(route, unsafeHandler)
//	}
//
//	return nil
//}
