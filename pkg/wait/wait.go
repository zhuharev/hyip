package wait

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

var (
	mu sync.RWMutex
	m  = map[int64]chan string{}

	DefaultTimeout time.Duration = time.Second * 20
)

func NeedNext(chatID int64, text string) bool {
	color.Green("check")
	mu.Lock()
	defer mu.Unlock()
	if ch, ok := m[chatID]; ok {
		ch <- text
		delete(m, chatID)
		return false
	}
	return true
}

// TODO: wait document photo video audio etcs
func Wait(chatID int64) (string, bool) {
	color.Green("start wait")
	mu.Lock()
	m[chatID] = make(chan string, 1)
	mu.Unlock()
	color.Green("unlock")

	select {
	case <-time.After(DefaultTimeout):
		color.Red("return with timeout")
		return "", false
	case text := <-m[chatID]:
		color.Green("return done")
		return text, true
	}
}
