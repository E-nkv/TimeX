package api

import (
	"fmt"
	"net/http"
	"strconv"
)

/*
3. GET /sessions/history/pie?timeMode=year&timeHorizon=current&category=2
visualMode = pie | graph
timeMode = year | month | week | day
timeHorizon = 'current' | 'previous' | Date in 'yyyy-mm-dd' format
category = int
*/

func queryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
func (app app) HandleGetHistory(w http.ResponseWriter, r *http.Request) {
	visualMode := queryParam(r, "visualMode")
	timeMode := queryParam(r, "timeMode")
	timeHorizon := queryParam(r, "timeHorizon")
	category_id := queryParam(r, "category_id")
	if !validVisualMode(visualMode) || !validTimeMode(timeMode) || !validTimeHorizon(timeHorizon) {
		fmt.Println("something was invalid. aborting")
		writeBadRequestError(w)
		return
	}

	app.ProcessHistory(w, visualMode, timeMode, timeHorizon, category_id)
}

func (app app) ProcessHistory(w http.ResponseWriter, visualMode, timeMode, timeHorizon, category_id string) {
	var catg int
	if category_id == "" {
		catg = -1
	} else {
		v, err := strconv.Atoi(category_id)
		if err != nil {
			catg = -1
		} else {
			catg = v
		}
	}

	if visualMode == "graph" {
		if timeMode == "day" {
			app.Service.GraphDay(w, timeHorizon, catg)
			return
		}
		if timeMode == "month" {
			app.Service.GraphMonth(w, timeHorizon, catg)
			return
		}
		if timeMode == "year" {
			app.Service.GraphYear(w, timeHorizon, catg)
			return
		}
	} else if visualMode == "pie" {
		if catg == -1 {
			writeBadRequestError(w)
			return
		}
		app.Service.Pie(w, timeMode, timeHorizon, catg)
		return
	}
	fmt.Println("got to the end of ProcessHistory and no handler could process the req")
	writeServerError(w)
}
