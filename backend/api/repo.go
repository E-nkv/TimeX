package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
	"timex/types"

	"github.com/jackc/pgx/v5"
)

type Repo struct {
	DB *pgx.Conn
}

func (r *Repo) InsertSession(ctx context.Context, s *types.InputSession) error {
	query := `INSERT INTO sessions (start, "end", focus, category_id) VALUES($1, $2, $3, $4)`
	startT := time.Unix(s.Start, 0)
	endT := time.Unix(s.End, 0)
	args := []any{startT.UTC().Truncate(time.Minute), endT.UTC().Truncate(time.Minute), s.Focus, s.CategoryID}
	if _, err := r.DB.Exec(ctx, query, args...); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return ErrDeadlineExceeded
		}
		return fmt.Errorf("%w, %+v", ErrInvalidArguments, err)

	}
	return nil
}

func (r *Repo) GetSession(ctx context.Context, id int64) (*types.Session, error) {
	q := `SELECT id, start, "end", category_id, focus, category_line FROM sessions WHERE id = $1`
	row := r.DB.QueryRow(ctx, q, id)

	var s types.Session

	args := []any{&s.ID, &s.Start, &s.End, &s.CategoryID, &s.Focus, &s.CategoryLine}
	if err := row.Scan(args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		//if the err aint nil, and its not ErrNoRows, but the category line is nil, it means that the error was sql: Scan error on column index 5, name "category_line": unsupported Scan, storing driver.Value type []uint8 into type *[]int, so we ignore it
		if s.CategoryLine != nil {
			fmt.Println("NOT SUPPOSED TO HAPPEN ERROR: ", err)
			return nil, err
		}
		//manually insert the category line (happens only when its a parent-less category, like 'study')
		//since lib/pq cannot convert a pgArray with len <= 1 into a goSlice.
		s.CategoryLine = []int{s.CategoryID}
	}

	return &s, nil
}

func (r *Repo) DeleteSession(ctx context.Context, id int64) error {
	q := "DELETE FROM sessions WHERE id = $1"
	rows, err := r.DB.Exec(ctx, q, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}

	if rows.RowsAffected() != 1 {
		return ErrNotFound
	}

	return nil
}

func (r *Repo) GraphDay(date time.Time, category int) ([]types.Session, error) {
	query := `select id, start, "end", category_id, focus, delta, category_line from sessions WHERE 
    start > ($1::DATE - '1 second' :: INTERVAL) AND 
    start < ($1::DATE + '23 hours 59 minutes'::INTERVAL) AND
    ($2 = -1 OR $2 = ANY(category_line))
    ORDER BY start;`
	res, err := r.DB.Query(context.Background(), query, date, category)
	if err != nil {
		return nil, err
	}
	var ss []types.Session
	for res.Next() {
		var s types.Session

		if err := res.Scan(&s.ID, &s.Start, &s.End, &s.CategoryID, &s.Focus, &s.Delta, &s.CategoryLine); err != nil {
			if strings.Contains(err.Error(), "category_line") {
				fmt.Println("hehe we are here puta")
				fmt.Println(err)
				continue
			}
			return nil, err
		}

		ss = append(ss, s)
	}
	return ss, nil

}
func (r *Repo) GraphMonth(date time.Time, category int) ([]types.Day, error) {
	query := `
		WITH 
			DaySeries AS (
				SELECT GENERATE_SERIES(1, (SELECT EXTRACT(DAY FROM (date_trunc('month', $1::DATE) + INTERVAL '1 month - 1 day')))) AS day
			),
			SessionTotals AS (
				SELECT EXTRACT(DAY FROM DATE_TRUNC('day', start)) AS day, SUM(delta) AS tot 
				FROM sessions 
				WHERE  
					start >= DATE_TRUNC('month', $1::DATE) - INTERVAL '1 day' AND 
					"end" <= DATE_TRUNC('month', $1::DATE + INTERVAL '1 month') - INTERVAL '1 day' AND
					($2 = -1 OR $2 = ANY(category_line))
				GROUP BY day
			),
			SuperTotal AS (
				SELECT SUM(tot) AS super_tot FROM SessionTotals
			)
		SELECT 
			NULL AS day, 
			COALESCE(st.super_tot, INTERVAL '0 seconds') AS tot
		FROM SuperTotal st
		UNION ALL
		SELECT 
			ds.day, 
			COALESCE(st.tot, INTERVAL '0 seconds') AS tot
		FROM DaySeries ds
		LEFT JOIN SessionTotals st ON ds.day = st.day
		ORDER BY day NULLS FIRST;
	`
	res, err := r.DB.Query(context.Background(), query, date, category)
	if err != nil {
		return nil, err
	}
	var ds []types.Day

	for res.Next() {
		var d types.Day

		if err := res.Scan(&d.DayNum, &d.Total); err != nil {
			if strings.Contains(err.Error(), "category_line") {
				fmt.Println("hehe we are here puta")
				fmt.Println(err)
				continue
			}
			return nil, err
		}

		ds = append(ds, d)
	}
	return ds, nil

}
func (r *Repo) GraphYear(date time.Time, category int) ([]types.Month, error) {
	query := `
		WITH MonthSeries AS (
			SELECT generate_series(1, 12) AS month
		),
		SessionTotals AS (
			SELECT 
				EXTRACT(MONTH FROM DATE_TRUNC('month', start)) AS month, 
				SUM(delta) AS tot 
			FROM sessions 
			WHERE 
				start > MAKE_TIMESTAMPTZ((EXTRACT (YEAR FROM $1::DATE) :: INT ), 1, 1, 0, 0, 0) - INTERVAL '1 day 1 second' AND
				"end" < MAKE_TIMESTAMPTZ((EXTRACT (YEAR FROM $1::DATE) :: INT + 1), 1, 1, 0, 0, 0) - INTERVAL '1 day 1 second' AND
				($2 = -1 OR $2 = ANY(category_line))
			GROUP BY month
		),
		SuperTotal AS (
			SELECT SUM(tot) AS super_tot FROM SessionTotals
		)
		SELECT 
			NULL AS month, 
			COALESCE(st.super_tot, INTERVAL '0 seconds') AS tot
		FROM SuperTotal st

		UNION ALL

		SELECT 
			ms.month, 
			COALESCE(st.tot, INTERVAL '0 seconds') AS tot
		FROM MonthSeries ms
		LEFT JOIN SessionTotals st ON ms.month = st.month
		ORDER BY month NULLS FIRST;
	`
	res, err := r.DB.Query(context.Background(), query, date, category)
	if err != nil {
		return nil, err
	}
	var ms []types.Month

	for res.Next() {
		var m types.Month

		if err := res.Scan(&m.MonthNum, &m.Total); err != nil {
			if strings.Contains(err.Error(), "category_line") {
				fmt.Println("hehe we are here puta")
				fmt.Println(err)
				continue
			}
			return nil, err
		}

		ms = append(ms, m)
	}
	return ms, nil

}

type Partition struct {
	Percent      float32       `json:"percent"`
	Total        time.Duration `json:"partition_total"`
	CategoryID   int           `json:"category_id"`
	CategoryName string        `json:"category_name"`
}

type PieResult struct {
	SuperTotal time.Duration `json:"super_total"`
	Partitions []Partition   `json:"pie_partitions"`
}

func (r *Repo) GetCategories() ([]types.CategoryS, error) {

	q := "SELECT id, name FROM categories"
	rs, err := r.DB.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}
	var res []types.CategoryS
	for rs.Next() {
		var c types.CategoryS
		if err := rs.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		res = append(res, c)
	}
	fmt.Println("res is: ", res)
	return res, nil
}
