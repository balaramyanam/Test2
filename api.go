package controllers

import (
	"encoding/json"
	"inter/model"
	"net/http"

	"github.com/gorilla/mux"
)

func Api(res http.ResponseWriter, req *http.Request) {

	ds := mux.Vars(req)

	copy := ds["name"]

	if copy != "" {

	}

	pradeep := model.Idlibatch{1, 25, "bed rest with ajai"}

	//db.Save(&pradeep)

	json.NewEncoder(res).Encode(pradeep)

}
