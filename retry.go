package retry

import (
	"fmt"
	"time"
)

// RetryOptions 重试选项
type RetryOptions struct {
	MaxRetries  int           // 最大重试次数
	InitialWait time.Duration // 初始等待时间
	MaxWait     time.Duration // 最大等待时间
}

// RetryFunc 重试函数类型
type RetryFunc func() error

// ExponentialBackoffRetry 指数级重试函数(返回error重试)
func ExponentialBackoffRetry(f RetryFunc, opts RetryOptions) error {
	var err error
	wait := opts.InitialWait
	for i := 0; i < opts.MaxRetries; i++ {
		if err = func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("panic occurred: %v", r)
				}
			}()
			return f()
		}(); err == nil {
			return nil
		}
		if i < opts.MaxRetries-1 {
			time.Sleep(wait)
			wait = time.Duration(float64(wait) * 2)
			if wait > opts.MaxWait {
				wait = opts.MaxWait
			}
		}
	}
	return fmt.Errorf("retries exhausted: %v", err)
}
