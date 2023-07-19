package retry

import (
	"errors"
	"testing"
	"time"
)

// InitialWait = 2 ^ (n - 1) * t
// MaxWait = 2 ^ n * t
// 其中，n 表示最大重试次数，t 表示初始等待时间。根据这个公式，可以计算出初始等待时间和最大等待时间的合理值，以便在重试时进行设置。
func TestExponentialBackoffRetry(t *testing.T) {
	var count int
	err := ExponentialBackoffRetry(func() error {
		count++
		t.Logf("time: [%v] count: [%d]", time.Now(), count)
		if count < 3 {
			return errors.New("retry")
		}
		return nil
	}, RetryOptions{
		MaxRetries:  3,
		InitialWait: 4 * time.Second,
		MaxWait:     8 * time.Second,
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if count != 3 {
		t.Errorf("unexpected count: %d", count)
	}
}

func TestExponentialBackoffRetryPanic(t *testing.T) {
	var count int
	err := ExponentialBackoffRetry(func() error {
		count++
		t.Logf("time: [%v] count: [%d]", time.Now(), count)
		if count < 3 {
			panic("x")
		}
		return nil
	}, RetryOptions{
		MaxRetries:  3,
		InitialWait: 4 * time.Second,
		MaxWait:     8 * time.Second,
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if count != 3 {
		t.Errorf("unexpected count: %d", count)
	}
}

func TestExponentialBackoffRetryError(t *testing.T) {
	var count int
	err := ExponentialBackoffRetry(func() error {
		count++
		t.Logf("time: [%v] count: [%d]", time.Now(), count)
		if count < 3 {
			return errors.New("failed")
		}
		return nil
	}, RetryOptions{
		MaxRetries:  3,
		InitialWait: 4 * time.Second,
		MaxWait:     8 * time.Second,
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if count != 3 {
		t.Errorf("unexpected count: %d", count)
	}
}
