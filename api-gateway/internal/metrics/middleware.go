package metrics

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware is a Gin middleware that records basic HTTP metrics.
func Middleware() gin.HandlerFunc {
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

	// убираем версию API
	path = strings.TrimPrefix(path, "/api/v1/")

	// ищем "settings/guild/"
	idx := strings.Index(path, "settings/guild/")
	if idx == -1 {
		return "other"
	}

	// убираем всё до guild/
	rest := path[idx+len("settings/guild/"):]
	// отбрасываем guild_id (оно может быть числом или :guild_id)
	parts := strings.SplitN(rest, "/", 2)
	if len(parts) < 2 {
		return "settings"
	}

	// берём первый сегмент после guild_id
	afterGuild := parts[1]
	if afterGuild == "" {
		return "settings"
	}

	group := strings.SplitN(afterGuild, "/", 2)[0]
	return group
}
