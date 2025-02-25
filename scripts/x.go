package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func OpenDB() (*sql.DB, error) {
	dsn := "postgresql://postgres:admin@localhost:5432/TimeX?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

func determineDates(timeMode, timeHorizon string) (start, end time.Time, err error) {
	switch {
	case timeMode == "day" && timeHorizon == "current":
		n := time.Now()
		start = time.Date(n.Year(), n.Month(), n.Day(), 0, 0, -1, 0, time.UTC)
		end = start.AddDate(0, 0, 1)
	case timeMode == "day" && timeHorizon == "previous":
		n := time.Now()
		start = time.Date(n.Year(), n.Month(), n.Day()-1, 0, 0, -1, 0, time.UTC)
		end = start.AddDate(0, 0, 1)
	case timeMode == "day":
		dt, err := time.Parse("2006-01-02", timeHorizon)
		if err != nil {
			return start, end, err
		}
		start = dt.Add(-time.Second)
		end = start.AddDate(0, 0, 1)
	case timeMode == "month" && timeHorizon == "current":
		n := time.Now()
		start = time.Date(n.Year(), n.Month(), 1, 0, 0, -1, 0, time.UTC)
		end = time.Date(n.Year(), n.Month()+1, 1, 0, 0, -1, 0, time.UTC)

	case timeMode == "month" && timeHorizon == "previous":
		n := time.Now()
		start = time.Date(n.Year(), n.Month()-1, 1, 0, 0, -1, 0, time.UTC)
		end = time.Date(n.Year(), n.Month(), 1, 0, 0, -1, 0, time.UTC)

	case timeMode == "month":
		n, err := time.Parse("2006-01-02", timeHorizon)
		if err != nil {
			return start, end, err
		}
		start = time.Date(n.Year(), n.Month(), 1, 0, 0, -1, 0, time.UTC)
		end = time.Date(n.Year(), n.Month()+1, 1, 0, 0, -1, 0, time.UTC)
	case timeMode == "year" && timeHorizon == "current":
		n := time.Now()
		start = time.Date(n.Year(), 1, 1, 0, 0, -1, 0, time.UTC)
		end = time.Date(n.Year()+1, 1, 1, 0, 0, -1, 0, time.UTC)

	case timeMode == "year" && timeHorizon == "previous":
		n := time.Now()
		start = time.Date(n.Year()-1, n.Month()-1, 1, 0, 0, -1, 0, time.UTC)
		end = time.Date(n.Year(), 1, 1, 0, 0, -1, 0, time.UTC)

	case timeMode == "year":
		n, err := time.Parse("2006-01-02", timeHorizon)
		if err != nil {
			return start, end, err
		}
		start = time.Date(n.Year(), 1, 1, 0, 0, -1, 0, time.UTC)
		end = time.Date(n.Year()+1, 1, 1, 0, 0, -1, 0, time.UTC)
	}
	return start, end, nil
}
func main() {
	m := map[int]int{}
	m[1] += 2
	fmt.Println(m)

}
