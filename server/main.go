package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

const PORT = ":3000"

type RFC3339Time struct {
	time.Time
}

func (rt RFC3339Time) MarshalJSON() ([]byte, error) {
	out := rt.Time.Format(time.RFC3339)
	return []byte(`"` + out + `"`), nil
}

func (rt *RFC3339Time) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(`"`+time.RFC3339+`"`, string(b))
	if err != nil {
		return err
	}
	*rt = RFC3339Time{t}
	return nil
}

func IPLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		slog.Info("ip of request",
			"ip", r.RemoteAddr)
	})
}

type jsonTime struct {
	DayOfWeek  string `json:"day_of_week"`
	Month      string `json:"month"`
	DayOfMonth int    `json:"day_of_month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!\n"))
	})
	mux.Handle("/time", IPLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		instance := jsonTime{
			DayOfWeek:  now.Weekday().String(),
			Month:      now.Month().String(),
			DayOfMonth: int(now.Month()),
			Year:       now.Year(),
			Hour:       now.Hour(),
			Minute:     now.Minute(),
			Second:     now.Second(),
		}
		// need to either return JSON or string based on the Accept header
		if r.Header.Get("Accept") == "application/json" {
			val, err := json.Marshal(instance)
			if err != nil {
				panic(err)
			}
			w.Write(val)
		} else {
			w.Write([]byte(`"` + now.Format(time.RFC3339) + `"`))

		}
	})))

	s := http.Server{
		Addr:         PORT,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
