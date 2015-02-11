package redis

import (
	"errors"

	"github.com/xuyu/goredis"
)

/*
	This REDIS implementation is a sharable component that can be switched out if
	a problem with the 3rd party comes up.  Other implementation ideas:

	https://github.com/gosexy/redis
	- Good mature API

	https://github.com/shipwire/redis
	- Uses io.Reader/Writer to make large data usage more efficient.
*/

// API is the inteface used to access the Redis cache
type API interface {
	// Get returns a value, if found.  If no value is found, nil is returned.
	Get(key string) ([]byte, error)
	// Set takes an interface and serializes to a string to be stored.
	Set(key string, data []byte) error
}

// DefaultAPI is the internal concrete implementation of the RedisAPI
type DefaultAPI struct {
	Debug bool
	URL   string

	client *goredis.Redis
}

// NewDefaultAPI returns a new instance of the API
func NewDefaultAPI(options ...func(*DefaultAPI)) API {

	// defaults
	api := DefaultAPI{}

	// setup options
	for _, o := range options {
		o(&api)
	}

	// open the connection and send a ping
	c, err := goredis.DialURL(api.URL)
	if err != nil {
		panic("Could not open the Redis data store: " + err.Error())
	}
	api.client = c
	if err := api.client.Ping(); err != nil {
		panic("Could not PING the redis data store: " + err.Error())
	}

	return &api
}

// Get returns a value, if found.  If no value is found, nil is returned.
func (api *DefaultAPI) Get(key string) ([]byte, error) {

	r, err := api.client.ExecuteCommand("GET", key)
	if err != nil {
		return nil, err
	}

	if r.Type != goredis.BulkReply {
		return nil, errors.New("Expected BulkReply for binary by the redis client.")
	}

	return r.BytesValue()
}

// Set takes an interface and serializes to a string to be stored.
func (api *DefaultAPI) Set(key string, data []byte) error {

	_, err := api.client.ExecuteCommand("SET", key, data)
	return err
}
