package api

import (
	"log"
	"net/http"
	"strconv"
)

func (app app) HandleGetCategories(w http.ResponseWriter, r *http.Request) {
	cats, err := app.Service.Repo.GetCategories()
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeResp(w, http.StatusOK, cats, "categories")
}
func (app app) HandleCreateCategory(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		//invalid stuff sending the category
		writeBadRequestError(w)
		return
	}
	parentIdStr := r.PostFormValue("id")
	name := r.PostFormValue("name")
	var parentId int
	if parentIdStr == "" {
		parentId = -1
	} else {
		v, err := strconv.Atoi(parentIdStr)
		if err != nil {
			writeBadRequestError(w)
			return
		}
		parentId = v
	}
	if err := app.Service.AddCategory(name, parentId); err != nil {
		log.Println(err)
		writeServerError(w)
		return
	}
	writeResp(w, http.StatusCreated, "success", "data")

}
