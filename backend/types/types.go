package types

import (
	"time"
)

/*
CREATE TABLE sessions (

	    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	    start TIMESTAMPTZ,
	    "end" TIMESTAMPTZ,
	    category_id INT REFERENCES categories(id),
	    focus SMALLINT,
		delta INTERVAL GENERATED ALWAYS AS ("end" - start ) STORED,
	    category_line INT[] GENERATED ALWAYS AS (category_line(category_id)) STORED

);
*/
type Session struct {
	ID           int64         `json:"id"`
	Start        time.Time     `json:"start"`
	End          time.Time     `json:"end"`
	CategoryID   int           `json:"category_id"`
	Focus        int           `json:"focus"`
	Delta        time.Duration `json:"delta"`
	CategoryLine []int         `json:"category_line"`
}

type Category struct {
	ID               int64  `json:"id"`
	ParentCategoryID int    `json:"parent_category_id"`
	Name             string `json:"name"`
}

type InputSession struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	CategoryID int    `json:"category_id"`
	Focus      int    `json:"focus"`
}

type Day struct {
	DayNum any           `json:"day"`
	Total  time.Duration `json:"total"`
}

type Month struct {
	MonthNum any           `json:"month"`
	Total    time.Duration `json:"total"`
}
