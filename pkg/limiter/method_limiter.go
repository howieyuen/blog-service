package limiter

import (
	"strings"
	
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

var _ LimitInterface = MethodLimiter{}

type MethodLimiter struct {
	*Limiter
}

func (m MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

func (m MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := m.limiterBuckets[key]
	return bucket, ok
	
}

func (m MethodLimiter) AddBucket(rules ...BucketRule) LimitInterface {
	for _, rule := range rules {
		if _, ok := m.GetBucket(rule.Key); !ok {
			m.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.Interval, rule.Capacity, rule.Quantum)
		}
	}
	return m
}

func NewMethodLimiter() LimitInterface {
	return MethodLimiter{
		Limiter: &Limiter{
			limiterBuckets: map[string]*ratelimit.Bucket{},
		},
	}
}
