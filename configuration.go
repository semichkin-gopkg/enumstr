package enumstr

import "github.com/semichkin-gopkg/conf"

type Config[E e, S s] struct {
	DefaultEnum E
	Formatter   func(string) string
}

func WithDefaultEnum[E e, S s](enum E) conf.Updater[Config[E, S]] {
	return func(c *Config[E, S]) {
		c.DefaultEnum = enum
	}
}

func WithFormatter[E e, S s](formatter func(string) string) conf.Updater[Config[E, S]] {
	return func(c *Config[E, S]) {
		c.Formatter = formatter
	}
}
