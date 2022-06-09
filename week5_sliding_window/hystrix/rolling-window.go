package hystrix

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type RollingWindow struct {
	sync.RWMutex
	broken bool
	//滑动窗口size
	swSize          int
	buckets         []*Bucket
	reqThreshold    int
	failedThreshold float64
	lastBreakTime   time.Time
	//熔断恢复的时间间隔
	brokeRecoverInterval time.Duration
}

func NewRollingWindow(
	size int,
	reqThreshold int,
	failedThreshold float64,
	brokeRecoverInterval time.Duration,
) *RollingWindow {
	return &RollingWindow{
		swSize:               size,
		buckets:              make([]*Bucket, 0, size),
		reqThreshold:         reqThreshold,
		failedThreshold:      failedThreshold,
		brokeRecoverInterval: brokeRecoverInterval,
	}
}

func (r *RollingWindow) AppendBucket() {
	r.Lock()
	defer r.Unlock()

	r.buckets = append(r.buckets, NewBucket())
	if !(len(r.buckets) < r.swSize+1) {
		r.buckets = r.buckets[1:]
	}
}

func (r *RollingWindow) GetBucket() *Bucket {
	if len(r.buckets) == 0 {
		r.AppendBucket()
	}
	return r.buckets[len(r.buckets)-1]
}

func (r *RollingWindow) RecordReqResult(resutl bool) {
	r.GetBucket().Record(resutl)
}

func (r *RollingWindow) ShowAllBucket() {
	for _, v := range r.buckets {
		fmt.Printf("id [%v] - total: [%d] failed: [%d]\n", v.Timestamp, v.Total, v.Failed)
	}
}

func (r *RollingWindow) Launch() {
	go func() {
		for {
			r.AppendBucket()
			time.Sleep(time.Millisecond * 100)
		}
	}()
}

//判断需要熔断
func (r *RollingWindow) BreakJudge() bool {
	r.RLock()

	defer r.Unlock()
	total := 0
	failed := 0

	for _, v := range r.buckets {
		total += v.Total
		failed += v.Failed
	}

	if float64(failed)/float64(total) > r.failedThreshold && total > r.reqThreshold {
		return true
	}
	return false
}

func (r *RollingWindow) Monitor() {
	go func() {
		for {
			if r.broken {
				if r.PassedBrokenInterval() {
					r.Lock()
					r.broken = false
					r.Unlock()
					continue
				}

				if r.BreakJudge() {
					r.Lock()
					r.broken = true
					r.lastBreakTime = time.Now()
					r.Unlock()
				}
			}
		}
	}()
}

func (r *RollingWindow) PassedBrokenInterval() bool {
	return time.Since(r.lastBreakTime) > r.brokeRecoverInterval
}

func (r *RollingWindow) ShowStatus() {
	go func() {
		for {
			log.Println("--broken status: ", r.broken)
			time.Sleep(time.Second)
		}
	}()
}

func (r *RollingWindow) Broken() bool {
	return r.broken
}
