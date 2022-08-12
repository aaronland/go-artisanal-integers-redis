package service

import (
	"context"
	"fmt"
	"github.com/aaronland/go-artisanal-integers/service"
	"github.com/gomodule/redigo/redis"
	"net/url"
	"strconv"
	"sync"
)

type RedisService struct {
	service.Service
	pool      *redis.Pool
	key       string
	offset    int64
	increment int64
	mu        *sync.Mutex
}

func NewRedisService(ctx context.Context, uri string) (service.Service, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	q := u.Query()

	dsn := q.Get("dsn")

	pool := &redis.Pool{
		MaxActive: 1000,
		Dial: func() (redis.Conn, error) {

			// https://www.iana.org/assignments/uri-schemes/prov/redis

			c, err := redis.DialURL(dsn)

			if err != nil {
				return nil, err
			}

			return c, err
		},
	}

	mu := new(sync.Mutex)

	s := RedisService{
		pool:      pool,
		key:       "integers",
		increment: 2,
		offset:    1,
		mu:        mu,
	}

	return &s, nil
}

func (s *RedisService) SetLastInt(ctx context.Context, i int64) error {

	last, err := s.LastInt(ctx)

	if err != nil {
		return err
	}

	if i < last {
		return fmt.Errorf("integer value too small")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	conn := s.pool.Get()
	defer conn.Close()

	_, err = conn.Do("SET", s.key, i)
	return err
}

func (s *RedisService) SetOffset(ctx context.Context, i int64) error {
	s.offset = i
	return nil
}

func (s *RedisService) SetIncrement(ctx context.Context, i int64) error {
	s.increment = i
	return nil
}

func (s *RedisService) LastInt(ctx context.Context) (int64, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	conn := s.pool.Get()
	defer conn.Close()

	redis_rsp, err := conn.Do("GET", s.key)

	if err != nil {
		return -1, err
	}

	b, err := redis.Bytes(redis_rsp, nil)

	if err != nil {
		return -1, err
	}

	i, err := strconv.ParseInt(string(b), 10, 64)

	if err != nil {
		return -1, err
	}

	return i, nil
}

func (s *RedisService) NextInt(ctx context.Context) (int64, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	conn := s.pool.Get()
	defer conn.Close()

	redis_rsp, err := conn.Do("INCRBY", s.key, s.increment)

	if err != nil {
		return -1, err
	}

	i, err := redis.Int64(redis_rsp, nil)

	if err != nil {
		return -1, err
	}

	return i, nil
}

func (s *RedisService) Close(ctx context.Context) error {
	return nil
}
