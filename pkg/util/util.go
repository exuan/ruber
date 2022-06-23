package util

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	cLockFormat = "lock:%s"
	Continuity  = "20060102150405"

	sLock = `
if redis.call("get",KEYS[1]) == ARGV[1]
then
return redis.call("del",KEYS[1])
else
return 0
end
`
)

func Md5(str string) string {
	m := md5.New()
	_, _ = io.WriteString(m, str)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func UnixTimestamp() int64 {
	return time.Now().Unix()
}

var num int64

func OrderNo(t time.Time) string {
	s := t.Format(Continuity)
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

//对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

func Lock(ctx context.Context, r *redis.Client, name string, token interface{}, ttl, retry int64) (lockRes bool) {
	for retry > 0 && lockRes == false {
		if lockRes = r.SetNX(ctx, fmt.Sprintf(cLockFormat, name), token, time.Duration(ttl)*time.Second).Val(); lockRes == true {
			retry = 0
		}
		retry--
		if retry > 0 {
			time.Sleep(time.Second)
		}
	}

	return lockRes
}

func UnLock(ctx context.Context, r *redis.Client, name string, token interface{}) bool {
	sp := redis.NewScript(sLock)
	v := sp.Run(ctx, r, []string{fmt.Sprintf(cLockFormat, name)}, token).Val()
	if v.(int64) == 1 {
		return true
	} else {
		return false
	}
}
