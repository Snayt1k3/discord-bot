package metrics

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// MetricsMiddleware is a Gin middleware that records basic HTTP metrics.
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()

		path := c.FullPath()
		if path == "" {
			path = "unknown"
		}
		group := extractGroup(path)

		status := strconv.Itoa(c.Writer.Status())

		HttpRequestsTotal.WithLabelValues(group, c.Request.Method, path, status).Inc()
		HttpRequestDuration.WithLabelValues(group, path).Observe(duration)
	}
}

// extractGroup вырезает группу после /settings/guild/:guild_id/
func extractGroup(path string) string {
	if path == "" {
		return "unknown"
	}

	// ищем начало settings-группы
	idx := strings.Index(path, "/settings/guild/:guild_id/")
	if idx == -1 {
		return "other"
	}

	// всё после /settings/guild/:guild_id/
	rest := strings.TrimPrefix(path[idx+len("/settings/guild/:guild_id/"):], "/")

	// если пусто, значит просто settings
	if rest == "" {
		return "settings"
	}

	// берём только первый сегмент, например "roles", "welcome", "automode"
	parts := strings.SplitN(rest, "/", 2)
	return parts[0]
}
