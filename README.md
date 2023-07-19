# go_retry
go语言重试函数
- 指数级重试函数(返回error重试)
  ```
  ExponentialBackoffRetry(func() error {
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
```
