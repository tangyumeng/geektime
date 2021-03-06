package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	gorma "github.com/hhxsv5/go-redis-memory-analysis"
)

var ctx = context.Background()

var rc *redis.Client

func main() {
	rc = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	write("len10_count10k", genValue(10), 10000)
	write("len10_count50k", genValue(10), 50000)
	write("len10_count500k", genValue(10), 500000)

	write("len1k_count10k", genValue(1000), 10000)
	write("len1k_count50k", genValue(1000), 50000)
	write("len1k_count500k", genValue(1000), 500000)

	write("len5k_count10k", genValue(5000), 10000)
	write("len5k_count50k", genValue(5000), 50000)
	write("len5k_count500k", genValue(5000), 500000)

	analysisMemory()
}

func write(key, value string, count int) {
	for i := 0; i < count; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		scmd := rc.Set(ctx, k, value, 0)
		err := scmd.Err()
		if err != nil {
			fmt.Println(scmd.String())
		}
	}
}

func genValue(len int) string {
	a := make([]byte, len)

	for i := 0; i < len; i++ {
		a[i] = 'v'
	}
	return string(a)
}

func analysisMemory() {
	analys, err := gorma.NewAnalysisConnection("localhost", 6379, "")

	if err != nil {
		fmt.Println("something wroing: ", err)
		return
	}
	defer analys.Close()

	analys.Start([]string{":"})

	err = analys.SaveReports("./reports")
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}
