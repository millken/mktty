package common

import (
	"fmt"

	"gopkg.in/redis.v3"
)

type Session struct {
	appKey string
	client *redis.Client
}

func NewSession(appKey string) (*Session, error) {
	return &Session{
		appKey: appKey,
	}, nil
}

func (s *Session) SetRedis(client *redis.Client) error {
	s.client = client
	return nil
}

func (s *Session) Set(key string, value interface{}) error {
	return s.client.Set(s.Key(key), value, 0).Err()
}

func (s *Session) Get(key string) (interface{}, error) {
	val, err := s.client.Get(s.Key(key)).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("key:%s does not exists", s.Key(key))
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

func (s *Session) Key(key string) string {
	return key
}
