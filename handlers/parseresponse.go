package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/saurabh2013/CRPCDemo/consts"
	"github.com/saurabh2013/CRPCDemo/log"
)

func ProcessError(w http.ResponseWriter, r *http.Request, e error) {
	if e != nil {
		parseJSON(w, r, map[string]interface{}{"Error:": e.Error()})
		log.Error(e)
		//fmt.Printf("STACK: %s", string(debug.Stack()))
	}
}

func ProcessResponse(w http.ResponseWriter, r *http.Request, d interface{}) {

	//To process string msgs
	switch txt := d.(type) {
	case string:
		d = map[string]string{"Status": txt}
		break
	}

	parseJSON(w, r, d)

}

func parseJSON(w http.ResponseWriter, r *http.Request, d interface{}) {

	w.Header().Set(consts.HEADER_CONTENTTYPE, consts.APPLICATION_JSON)
	res, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		ProcessError(w, r, err)
	}
	fmt.Fprint(w, string(res))
}
