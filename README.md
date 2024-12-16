# go_retry

**go_retry** 是一个用 Go 语言编写的重试库，旨在提供简单而强大的重试机制，以增强你的应用程序的健壮性。

## 特点

- **指数级重试**：自动根据指数退避策略进行重试，直到成功或达到最大重试次数。
- **灵活配置**：允许自定义最大重试次数、初始等待时间和最大等待时间。
- **易于集成**：轻松集成到任何 Go 项目中，无需复杂的设置。

## 安装

使用 `go get` 命令安装：

```sh
go get github.com/SenLinLeo/go_retry
```

## 使用案例

以下是如何使用 `go_retry` 进行指数级重试的示例：

```go
package main

import (
	"fmt"
	"time"

	"github.com/SenLinLeo/go_retry"
)

func main() {
	err := go_retry.ExponentialBackoffRetry(func() error {
		fmt.Println("Attempting operation...")
		// 模拟操作，这里返回 nil 表示成功
		return nil
	}, go_retry.RetryOptions{
		MaxRetries:  3,
		InitialWait: 4 * time.Second,
		MaxWait:     8 * time.Second,
	})

	if err != nil {
		fmt.Println("Operation failed after retries:", err)
	} else {
		fmt.Println("Operation succeeded!")
	}
}
```

## 贡献

我们欢迎任何形式的贡献！无论是修复 bug、添加新功能还是改进文档，你的帮助都是宝贵的。请随时提交 pull request 或 issue。


## 为什么选择 go_retry？

- **简洁的 API**：易于理解和使用。
- **强大的功能**：提供指数级重试，帮助你处理暂时性错误。

