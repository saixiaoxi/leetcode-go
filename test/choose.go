package main

import "time"

// 默认缓存设置
const (
	defaultCaching = false
)

// 连接结构体
type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

// 配置选项结构体
type options struct {
	timeout time.Duration
	caching bool
}

// Option 接口，要求实现 apply 方法
type Option interface {
	apply(*options)
}

// 函数类型起别名，类型是参数为 *options 返回为空的函数
type optionFunc func(*options)

// 实现 Option 接口的 apply 方法
func (f optionFunc) apply(o *options) {
	f(o)
}

// 创建一个设置超时的 Option
func WithTimeout(t time.Duration) Option {
	// 把参数为 *options 返回为空的函数类型转换为 optionFunc 类型
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

// 创建一个设置缓存的 Option
func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}
