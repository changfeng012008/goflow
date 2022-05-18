package runtime

import (
	redisStateStore "github.com/changfeng_012008/goflow/core/redis-statestore"
	"github.com/changfeng_012008/goflow/core/sdk"
)

func initStateStore(redisURI string) (stateStore sdk.StateStore, err error) {
	stateStore, err = redisStateStore.GetRedisStateStore(redisURI)
	return stateStore, err
}
