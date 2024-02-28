package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func MiddlewareWithTimeout(millis int) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			withTimeout, err := context.WithTimeout(ctx, time.Millisecond*time.Duration(millis))
			if err != nil {
				w.WriteHeader(http.StatusRequestTimeout)
				w.Write([]byte("timeout"))
				return
			}
			r = r.WithContext(withTimeout)
			h.ServeHTTP(w, r)
		})
	}
}

func part2() {
	sum := 0
	iterations := 0
	ctx := context.Background()
	parent, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
loop:
	for {
		iterations++
		num := rand.Intn(100_000_000)
		sum += num
		if num == 1234 {
			fmt.Println("hit the magic number")
			break loop
		}
		select {
		case <-parent.Done():
			break loop
		default:
		}
	}
	fmt.Println("sum", sum, "iterations", iterations)
}

type Level string

const Debug Level = "debug"
const Info Level = "info"

type levelKey struct{}

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	inLevel, ok := LogLevelFromContext(ctx)
	if !ok {
		return
	}
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

func ContextWithLogLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, levelKey{}, level)
}

func LogLevelFromContext(ctx context.Context) (Level, bool) {
	level, ok := ctx.Value(levelKey{}).(Level)
	return level, ok
}

func MiddlewareWithLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		level := Level(r.URL.Query().Get("log_level"))
		ctx = ContextWithLogLevel(ctx, level)
		withTimeout, err := context.WithTimeout(ctx, time.Millisecond*time.Duration(millis))
		if err != nil {
			w.WriteHeader(http.StatusRequestTimeout)
			w.Write([]byte("timeout"))
			return
		}
		r = r.WithContext(withTimeout)
		h.ServeHTTP(w, r)
	})
}
