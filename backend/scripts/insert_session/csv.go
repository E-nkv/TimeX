package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
	"timex/api"
	"timex/database"
	"timex/types"

	"github.com/jackc/pgx/v5"
)

func main() {
	pc, fi, l, ok := runtime.Caller(0)
	fmt.Println(pc, fi, l, ok)
	db, err := database.OpenDB()
	if err != nil {
		panic(err)
	}
	svc := api.Service{Repo: api.Repo{DB: db}}
	ctx, c := context.WithTimeout(context.Background(), time.Second*10)
	defer c()
	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err)
	}
	base := "./scripts/insert_session"
	filename := "file1.csv"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	p := filepath.Join(base, filename)
	f, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(f)
	for i := 0; ; i++ {

		rc, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("Csv bulk insert finished succesfully")
				if err := tx.Commit(ctx); err != nil {
					panic(err)
				}

				return
			}
		}
		if rc[0] == "\n" {
			continue
		}
		if len(rc) != 4 {
			tx.Rollback(ctx)
			panic("invalid format for the session insertions")
		}
		var s types.InputSession

		f, err := strconv.Atoi(rc[2])
		if err != nil {
			tx.Rollback(ctx)
			panic("not an int the focus")
		}

		ctg, err := strconv.Atoi(rc[3])
		if err != nil {
			tx.Rollback(ctx)
			panic("not an int the catg id")
		}

		s.CategoryID = ctg
		s.Focus = f
		s.Start = rc[0]
		s.End = rc[1]

		if err := svc.InsertSession(ctx, &s); err != nil {
			fmt.Println("erroing iteration: ", i)
			fmt.Println(rc)
			tx.Rollback(ctx)
			panic(err)
		}
	}
}
