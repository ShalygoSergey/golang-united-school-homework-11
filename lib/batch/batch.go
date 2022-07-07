package batch

import (
	"context"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var mutex sync.Mutex
	errG, _ := errgroup.WithContext(context.Background())
	errG.SetLimit(int(pool))
	for i := 0; i < int(n); i++ {
		id := i
		errG.Go(func() error {
			one := getOne(int64(id))
			mutex.Lock()
			res = append(res, one)
			mutex.Unlock()
			return nil
		})
	}

	err := errG.Wait()
	if err != nil {
		panic(err)
	}

	return res
}
