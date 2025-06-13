package infra

import (
	"fmt"
	"testing"
	"time"
)

func Test_redisclient(t *testing.T) {
	defer RedisConnectionsClose()

	const key = "redis_key"

	client, err := NewRedisClient()
	if err != nil {
		t.Fatalf("Redisクライアントの作成に失敗: %v", err)
	}

	redis := client.GetClient()
	if redis == nil {
		t.Fatal("Redisクライアントがnilです")
	}

	err = redis.Set(client.ctx, key, "sample_value", time.Minute).Err()
	if err != nil {
		t.Fatalf("値の設定に失敗: %v", err)
	}

	fmt.Println(err)

	fmt.Println(
		REDIS_ADDR,
		REDIS_PASSWORD,
		REDIS_USERNAME,
	)

	if !client.ConnectionComplite() {
		t.Fatal("redis connection failed")
	}

	val, err := redis.Get(client.ctx, key).Result()
	if err != nil {
		t.Fatalf("値の取得に失敗: %v", err)
	}

	t.Logf("取得した値: %s", val)

	dunp_cmd := redis.Conn().Dump(client.ctx, key)
	fmt.Println(dunp_cmd.Bool())

}

func Test_client_method(t *testing.T) {

	client, _ := NewRedisClient()

	defer client.Close()
	const Key = "self_key"
	if client.Set(Key, "1234", 2) {
		fmt.Println("set is succsecd")
	}

	val, err := client.Get(Key)
	if err != nil {
		t.Fatalf("fatal:%s", err.Error())
	}

	fmt.Print(val)
}
