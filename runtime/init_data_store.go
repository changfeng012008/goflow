package runtime

import (
	redisDataStore "github.com/changfeng012008/goflow/core/redis-datastore"
	"github.com/changfeng012008/goflow/core/sdk"
)

func initDataStore(redisURI string) (dataStore sdk.DataStore, err error) {
	dataStore, err = redisDataStore.GetRedisDataStore(redisURI)
	return dataStore, err
}
