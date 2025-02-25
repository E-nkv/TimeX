package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"
	"timex/types"
)

type Service struct {
	Repo Repo
}

// ******************** session related stuff ********************
func (s Service) InsertSession(ctx context.Context, session *types.InputSession) error {
	return s.Repo.InsertSession(ctx, session)
}
func (s Service) DeleteSession(ctx context.Context, id int64) error {
	return s.Repo.DeleteSession(ctx, id)
}

func (s Service) GetSession(ctx context.Context, id int64) (*types.Session, error) {
	sess, err := s.Repo.GetSession(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return sess, nil
}

// ******************** history related stuff ********************
func (s Service) GraphDay(w http.ResponseWriter, timeHorizon string, category int) {
	date, err := computeDate(timeHorizon, "day")
	if err != nil {
		//technically shouldnt happen since validation works well at the handler-level. but just in case...
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	ss, err := s.Repo.GraphDay(date, category)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	var tot time.Duration
	for _, s := range ss {
		tot += s.Delta
	}
	res := map[string]any{
		"super_total": tot,
		"sessions":    ss,
	}
	writeResp(w, http.StatusOK, res, "data")
}

func (s Service) GraphMonth(w http.ResponseWriter, timeHorizon string, category int) {
	date, err := computeDate(timeHorizon, "month")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	ds, err := s.Repo.GraphMonth(date, category)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := map[string]any{
		"super_total": ds[0].Total,
		"days":        ds[1:],
	}
	writeResp(w, http.StatusOK, res, "data")
}
func (s Service) GraphYear(w http.ResponseWriter, timeHorizon string, category int) {
	date, err := computeDate(timeHorizon, "year")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	ms, err := s.Repo.GraphYear(date, category)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := map[string]any{
		"super_total": ms[0].Total,
		"months":      ms[1:],
	}
	writeResp(w, http.StatusOK, res, "data")
}

func computeDate(timeHorizon, timeMode string) (time.Time, error) {
	var date time.Time
	var y, m, d int

	switch timeMode {
	case "year":
		y--
	case "month":
		m--
	case "day":
		d--
	}

	if timeHorizon == "current" {
		date = time.Now().UTC().Truncate(time.Hour * 24)
	} else if timeHorizon == "previous" {
		date = time.Now().UTC().Truncate(time.Hour*24).AddDate(y, m, d)
	} else {
		dateP, err := time.Parse("2006-01-02", timeHorizon)
		if err != nil {
			return date, err
		}
		date = dateP
	}
	return date, nil
}

func (svc Service) PieDistribution(start, end time.Time, category int) (PieResult, error) {
	q := `
		select start, "end", category_id, focus, delta, category_line, c.name 
		from sessions s 
		JOIN categories c ON c.id = s.category_id
		WHERE start BETWEEN $1 AND $2 AND $3 = ANY(category_line)
		
	`
	rows, err := svc.Repo.DB.Query(context.Background(), q, start, end, category)
	if err != nil {
		return PieResult{}, err
	}
	res := PieResult{}

	groupingMap := map[int]time.Duration{}
	categoryNamesMap := map[int]string{}
	var s types.Session
	var CategoryName string
	for rows.Next() {
		if err := rows.Scan(&s.Start, &s.End, &s.CategoryID, &s.Focus, &s.Delta, &s.CategoryLine, &CategoryName); err != nil {
			return PieResult{}, err
		}
		var keyCatId int
		if s.CategoryID == category {
			keyCatId = category
		} else {
			i := slices.Index(s.CategoryLine, category)
			if i == len(s.CategoryLine)-1 {
				//should happen, just in case to theoretically avoid out of bounds err
				return PieResult{}, fmt.Errorf("logic of the multiverse has been twisted. 'i' should never be the last index, logically.")
			}
			keyCatId = s.CategoryLine[i+1]
		}
		groupingMap[keyCatId] += s.Delta
		categoryNamesMap[keyCatId] = CategoryName

		res.SuperTotal += s.Delta
	}
	for k, v := range groupingMap {
		var catName string
		if k == category {
			catName = "other"
		} else {
			catName = categoryNamesMap[k]
		}
		p := Partition{
			CategoryID:   k,
			CategoryName: catName,
			Total:        v,
			Percent:      float32(v * 100 / res.SuperTotal),
		}
		res.Partitions = append(res.Partitions, p)
	}

	return res, nil
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
func (s Service) Pie(w http.ResponseWriter, timeMode, timeHorizon string, category int) {
	start, end, err := determineDates(timeMode, timeHorizon)
	if err != nil {
		fmt.Println(err)
		writeBadRequestError(w)
		return
	}
	res, err := s.PieDistribution(start, end, category)
	if err != nil {
		fmt.Println(err)
		writeBadRequestError(w)
		return
	}
	writeResp(w, http.StatusOK, res, "data")
}

func (svc Service) AddCategory(name string, parent_id int) error {
	ctx := context.Background()
	q := `INSERT INTO categories (parent_category_id, name) VALUES ((NULLIF($1, -1)),$2)`
	if _, err := svc.Repo.DB.Exec(ctx, q, parent_id, name); err != nil {
		return err
	}
	return nil
}
