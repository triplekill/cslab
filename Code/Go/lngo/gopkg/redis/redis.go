package main

import (
	"fmt"
	"time"

	"github.com/exfly/lngo/model"
	gocache "github.com/go-redis/cache"
	goredis "github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
	cache "github.com/huntsman-li/go-cache"
	_ "github.com/huntsman-li/go-cache/redis"
	"github.com/vmihailenco/msgpack"
)

// go get github.com/gomodule/redigo/redis
func objToRedisSet() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		// handle error
	}
	defer c.Close()
	v, err := c.Do("SET", "name", "red")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	v, err = redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

	var p1, p2 struct {
		Title  string `redis:"title"`
		Author string `redis:"author"`
		Body   string `redis:"body"`
	}

	p1.Title = "Example"
	p1.Author = "Gary"
	p1.Body = "Hello"

	if _, err := c.Do("HMSET", redis.Args{}.Add("id1").AddFlat(&p1)...); err != nil {
		fmt.Println(err)
		return
	}

	m := map[string]string{
		"title":  "Example2",
		"author": "Steve",
		"body":   "Map",
	}

	if _, err := c.Do("HMSET", redis.Args{}.Add("id2").AddFlat(m)...); err != nil {
		fmt.Println(err)
		return
	}

	for _, id := range []string{"id1", "id2"} {

		v, err := redis.Values(c.Do("HGETALL", id))
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := redis.ScanStruct(v, &p2); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%+v\n", p2)
	}
}

func complexTRedis() {
	p := model.People{
		ID:   "myid",
		Name: "myname",
		Age:  10,
		Toys: []model.Toy{
			model.Toy{
				ID:   "toy1",
				Name: "toy1name",
			},
			model.Toy{
				ID:   "toy2",
				Name: "toy2name",
			},
		},
	}
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		// handle error
	}
	defer c.Close()
	if _, err := c.Do("HMSET", redis.Args{}.Add("people").AddFlat(&p)...); err != nil {
		fmt.Println(err)
		return
	}
	id := "people"
	v, err := redis.Values(c.Do("HGETALL", id))
	if err != nil {
		fmt.Println(err)
		return
	}
	p2 := model.People{}
	if err := redis.ScanStruct(v, &p2); err != nil {
		fmt.Println(err) // redigo can't struct one complex struct
		return
	}

	fmt.Printf("%+v\n", p2)
}

// go get -u github.com/huntsman-li/go-cache
func cacheWithRedis() {
	ca, err := cache.Cacher(cache.Options{
		Adapter:       "redis",
		AdapterConfig: "addr=127.0.0.1:6379,prefix=cache",
		OccupyMode:    true,
	})

	if err != nil {
		panic(err)
	}
	p := model.People{
		ID:   "myid",
		Name: "myname",
		Age:  10,
		Toys: []model.Toy{
			model.Toy{
				ID:   "toy1",
				Name: "toy1name",
			},
			model.Toy{
				ID:   "toy2",
				Name: "toy2name",
			},
		},
	}
	ca.Put("liyan", "cache", 60)
	ca.Put("cache_people", &p, 20)
	// pg := model.People{}
	fmt.Printf("%T %v", ca.Get("cache_people"), ca.Get("cache_people"))

}

// gocache "github.com/go-redis/cache"
// goredis "github.com/go-redis/redis"
func cacheWithRedis2() {
	p := model.People{
		ID:   "myid",
		Name: "myname",
		Age:  10,
		Toys: []model.Toy{
			model.Toy{
				ID:   "toy1",
				Name: "toy1name",
			},
			model.Toy{
				ID:   "toy2",
				Name: "toy2name",
			},
		},
	}
	ring := goredis.NewRing(&goredis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
		},
	})
	codec := &gocache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
	key := "mykey"
	codec.Set(&gocache.Item{
		Key:        key,
		Object:     p,
		Expiration: time.Hour,
	})

	var wanted model.People
	if err := codec.Get(key, &wanted); err == nil {
		fmt.Println(wanted)
		fmt.Println(wanted.ID)
	}
}
func redisUse() {
	p := model.People{
		ID:   "myid",
		Name: "myname",
		Age:  10,
		Toys: []model.Toy{
			model.Toy{
				ID:   "toy1",
				Name: "toy1name",
			},
			model.Toy{
				ID:   "toy2",
				Name: "toy2name",
			},
		},
	}

	client := goredis.NewClient(&goredis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	err = client.Set("goredis_people", p, 0).Err() // There is a err, can't set one obj
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == goredis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
func main() {
	// objToRedisSet()
	// complexTRedis() // have one complex and thinkable err
	// cacheWithRedis()
	// cacheWithRedis2()
	redisUse()
}
