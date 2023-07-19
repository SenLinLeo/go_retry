# go_retry
go语言重试函数
- 指数级重试函数(返回error重试)
  使用案例，欢迎交流
  ```
	ExponentialBackoffRetry(func() error {
		fmt.Println("ExponentialBackoffRetry")
		return nil
	}, RetryOptions{
		MaxRetries:  3,
		InitialWait: 4 * time.Second,
		MaxWait:     8 * time.Second,
	})
```
