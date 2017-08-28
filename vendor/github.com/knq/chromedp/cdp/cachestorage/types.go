package cachestorage

import "github.com/mailru/easyjson"

// Code generated by chromedp-gen. DO NOT EDIT.

// CacheID unique identifier of the Cache object.
type CacheID string

// String returns the CacheID as string value.
func (t CacheID) String() string {
	return string(t)
}

// DataEntry data entry.
type DataEntry struct {
	Request      string  `json:"request"`      // Request url spec.
	Response     string  `json:"response"`     // Response status text.
	ResponseTime float64 `json:"responseTime"` // Number of seconds since epoch.
}

// Cache cache identifier.
type Cache struct {
	CacheID        CacheID `json:"cacheId"`        // An opaque unique id of the cache.
	SecurityOrigin string  `json:"securityOrigin"` // Security origin of the cache.
	CacheName      string  `json:"cacheName"`      // The name of the cache.
}

// CachedResponse cached response.
type CachedResponse struct {
	Headers easyjson.RawMessage `json:"headers"`
	Body    string              `json:"body"` // Entry content, base64-encoded.
}
