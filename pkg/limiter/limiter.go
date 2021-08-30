package limiter

import (
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimitInterface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBucket(rules ...BucketRule) LimitInterface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type BucketRule struct {
	Key      string
	Interval time.Duration
	Capacity int64
	Quantum  int64
}
