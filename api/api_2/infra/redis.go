package infra

import (
	"context"
	"crypto/tls"
	"fmt"
	"module/util"
	"time"

	"github.com/redis/go-redis/v9"
)

// 環境変数名
const (
	ENV_KEY__REDIS_DNS_NAME = "REDIS_DNS_NAME"
	ENV_KEY__REDIS_PORT     = "REDIS_PORT"
	ENV_KEY__REDIS_PASSWORD = "REDIS_PASSWORD"
	ENV_KEY__REDIS_DB_NO    = "REDIS_DB_NO"
	ENV_KEY__REIDS_USERNAME = "REDIS_USERNAME"
)

// 環境変数を格納する静的変数
var (
	REDIS_ADDR     string
	REDIS_DB_NO    int
	REDIS_PASSWORD string
	REDIS_USERNAME string
	CONNECTIONS    []*RedisClient
)

func init() {
	fmt.Println("REDIS ENV START")
	redis_dns_name := util.GetEnv(ENV_KEY__REDIS_DNS_NAME, "localhost")
	redis_port := util.GetEnv(ENV_KEY__REDIS_PORT, "6379")

	REDIS_ADDR = fmt.Sprintf("%s:%s", redis_dns_name, redis_port)

	REDIS_DB_NO = util.ToInt(
		util.GetEnv(ENV_KEY__REDIS_DB_NO, "0"),
	)

	fmt.Println("REIDS ADDR ::: ", REDIS_ADDR)

	REDIS_PASSWORD = util.GetEnv(ENV_KEY__REDIS_PASSWORD, "")
	REDIS_USERNAME = util.GetEnv(ENV_KEY__REIDS_USERNAME, "default")

	CONNECTIONS = make([]*RedisClient, 0)

}

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
	status bool
}

func (rc *RedisClient) Set(key string, val interface{}, ttl_count_by_minute int) bool {
	err := rc.client.Set(rc.NewContext(), key, val, time.Minute*time.Duration(ttl_count_by_minute))
	return err == nil
}

func (rc *RedisClient) Get(key string) (string, error) {
	ctx := rc.NewContext()
	return rc.client.Get(ctx, key).Result()
}

func (rc *RedisClient) GetClient() *redis.Client {
	return rc.client
}

func (rc *RedisClient) NewContext() context.Context {
	return context.Background()
}

func (rc *RedisClient) ConnectionComplite() bool {

	err := rc.client.Ping(rc.NewContext()).Err()
	return err == nil
}

func (rc *RedisClient) HealthCheck() bool {
	ctx, cancel := context.WithTimeout(rc.NewContext(), 50*time.Minute)
	defer cancel()

	if rc.client == nil {
		fmt.Println("REDIS Client is none")
		return false
	}

	_, err := rc.client.Ping(ctx).Result()
	return err == nil
}

func (rc *RedisClient) ConnectionErr() string {
	wati_time := 30 * time.Second
	ctx, cancel := context.WithTimeout(rc.NewContext(), wati_time)
	defer cancel()
	err := rc.client.Ping(ctx).Err()
	if err != nil {
		fmt.Println("redis to Ping ConnectionErr", err.Error())
		return err.Error()
	}

	return "OK"
}

func (rc *RedisClient) Close() {
	rc.client.Close()
}

func RedisConnectionsClose() {
	for _, redis_clinet := range CONNECTIONS {
		redis_clinet.Close()
	}
}

//--------------------------------------------------//

func NewRedisClient() (*RedisClient, error) {
	fmt.Println("NEW REDIS CLIENT CREATED")
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		DB:       REDIS_DB_NO,
		Password: REDIS_PASSWORD,
		Username: REDIS_USERNAME,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			fmt.Println("valkey接続完了")
			return nil
		},
		TLSConfig:    &tls.Config{},
		PoolSize:     10,
		MinIdleConns: 5,
	})

	ctx := context.Background()
	redis_clinet := &RedisClient{
		client: client,
		ctx:    ctx,
		status: false,
	}

	go func(client_redis *RedisClient) {
		fmt.Println("redisとの接続を確認:::", client_redis.HealthCheck())
	}(redis_clinet)

	CONNECTIONS = append(CONNECTIONS, redis_clinet)
	return redis_clinet, nil
}
