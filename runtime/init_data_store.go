package runtime

import (
	redisDataStore "github.com/changfeng_012008/goflow/core/redis-datastore"
	"github.com/changfeng_012008/goflow/core/sdk"
)

func initDataStore(redisURI string) (dataStore sdk.DataStore, err error) {
	dataStore, err = redisDataStore.GetRedisDataStore(redisURI)
	return dataStore, err
}
