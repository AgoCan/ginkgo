package redis

import (
	rdb "github.com/go-redis/redis/v8"
)

type Options struct {
	Addr     string
	Password string
	DB       int
}

type Client struct {
	Rdb *rdb.Client
}

func NewClient(o *Options) *Client {
	var c Client
	c.Rdb = rdb.NewClient(&rdb.Options{
		Addr:     o.Addr,
		Password: o.Password,
		DB:       o.DB,
	})
	return &c
}
