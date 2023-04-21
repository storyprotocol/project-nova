package middleware

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/pkg/constant"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/prometheus/client_golang/prometheus"
)

const ()

func Prometheus(subsystem string) gin.HandlerFunc {
	reqCnt := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "requests_total",
			Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"code", "method", "url"},
	)
	if err := prometheus.Register(reqCnt); err != nil {
		logger.Fatal("requests_total could not be registered in Prometheus")
	}

	reqDur := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: subsystem,
			Name:      "request_duration_seconds",
			Help:      "The HTTP request latencies in seconds.",
		},
		[]string{"code", "method", "url"},
	)
	if err := prometheus.Register(reqDur); err != nil {
		logger.Fatal("request_duration_seconds could not be registered in Prometheus")
	}

	return func(c *gin.Context) {
		if c.Request.URL.Path == constant.MetricsPath {
			c.Next()
			return
		}

		start := time.Now()

		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		elapsed := float64(time.Since(start)) / float64(time.Second)

		// The URL path is not parameterized. For example for /v1/story/:franchiseId/:storyNum,
		// the URL path is /v1/story/1/1 or /v1/story/1/2, depending on the request input.
		// We need to parameterize them and make them the same label to reduce cardinality of
		// the metrics and increase readability.
		// In order to do that, we also need to separate prefix like /v1/ before parameterizing the value
		// and re-append it afterward
		prefix, url := separatePrefix(c.Request.URL.Path)
		for _, p := range c.Params {
			url = strings.Replace(url, p.Value, ":"+p.Key, 1)
		}
		url = prefix + url

		reqCnt.WithLabelValues(status, c.Request.Method, url).Inc()
		reqDur.WithLabelValues(status, c.Request.Method, url).Observe(elapsed)
	}
}

func separatePrefix(url string) (string, string) {
	// Match two groups
	// Prefix: It can be /admin/v1/ or /v1/
	// Path: The rest of the url path
	r := regexp.MustCompile(`(?P<Prefix>\D*\/v[0-9]+\/)(?P<Path>.*)`)
	matches := r.FindStringSubmatch(url)
	return matches[1], matches[2]
}
