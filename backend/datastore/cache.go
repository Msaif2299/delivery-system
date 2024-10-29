package datastore

import "context"

type Cache interface {
	Get(ctx context.Context, key string, val interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
}
