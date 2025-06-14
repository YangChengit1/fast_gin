package middleware

// 令牌桶限流
import (
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
	"time"
)

// LimitMiddleware 工厂函数
func LimitMiddleware(limit int) gin.HandlerFunc {
	return NewLimiter(limit, 1*time.Second).Middleware
}

// Limiter 限流器结构体
type Limiter struct {
	limit      int                // 限制的请求数量
	duration   time.Duration      // 时间窗口长度
	timestamps map[string][]int64 // 存储每个IP的请求时间戳 {IP: [时间戳1, 时间戳2,...]}
}

// NewLimiter 工厂函数创建限流器实例
// 参数:
//   - limit: 时间窗口内允许的最大请求数
//   - duration: 时间窗口长度(如1分钟、1小时等)
func NewLimiter(limit int, duration time.Duration) *Limiter {
	return &Limiter{
		limit:      limit,
		duration:   duration,
		timestamps: make(map[string][]int64),
	}
}

// Middleware Gin限流中间件，实际处理限流的逻辑函数
func (l *Limiter) Middleware(c *gin.Context) {
	// 1. 获取客户端IP作为唯一标识
	ip := c.ClientIP()

	// 2. 初始化该IP的时间戳记录(如果不存在)
	if _, ok := l.timestamps[ip]; !ok {
		l.timestamps[ip] = make([]int64, 0)
	}

	// 3. 获取当前时间戳(秒级)
	now := time.Now().Unix()

	// 4. 清理过期的请求记录(滑动窗口核心逻辑)
	// 只保留在时间窗口内的请求记录
	validTimestamps := make([]int64, 0)
	for _, ts := range l.timestamps[ip] {
		// 如果时间戳在时间窗口内(now - duration <= ts <= now)
		if ts >= now-int64(l.duration.Seconds()) {
			validTimestamps = append(validTimestamps, ts)
		}
	}
	l.timestamps[ip] = validTimestamps

	// 5. 检查当前请求数是否超过限制
	if len(l.timestamps[ip]) >= l.limit {
		// 返回429 Too Many Requests错误
		res.FailWithMsg("Too Many Requests", c)
		c.Abort() // 终止请求处理链
		return
	}

	// 6. 记录当前请求的时间戳
	l.timestamps[ip] = append(l.timestamps[ip], now)

	// 7. 继续处理后续中间件和请求处理函数
	c.Next()
}
