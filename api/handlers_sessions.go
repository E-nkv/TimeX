package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"timex/types"

	"github.com/go-chi/chi/v5"
)

const timeoutDur = time.Second * 5

func (app *app) HandleCreateSession(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeoutDur)
	defer cancel()
	var inp types.InputSession

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&inp); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := validInputSession(inp); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := app.Service.InsertSession(ctx, &inp); err != nil {
		switch err {
		case ErrDeadlineExceeded:
			writeError(w, http.StatusRequestTimeout, "request timeout")
			return
		case ErrInvalidArguments:
			writeBadRequestError(w)
			return
		default:
			writeServerError(w)
			return
		}
	}

	writeResp(w, http.StatusOK, "success creating the session", "success")
}

func (app *app) HandleGetSession(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeoutDur)
	defer cancel()

	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		writeBadRequestError(w)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		writeBadRequestError(w)
		return
	}
	s, err := app.Service.GetSession(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			writeBadRequestError(w)
			return
		}
		fmt.Println(err)
		writeServerError(w)
		return
	}
	writeResp(w, http.StatusOK, s, "session")
}

func (app *app) HandleDeleteSession(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeoutDur)
	defer cancel()

	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		writeBadRequestError(w)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		writeBadRequestError(w)
		return
	}
	err = app.Service.DeleteSession(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			writeBadRequestError(w)
			return
		}
		writeServerError(w)
		return
	}
	writeResp(w, http.StatusOK, "session deleted", "success")
}
