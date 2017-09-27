package rest

import (
	"net/http"
	"github.com/zamedic/go2hal/database"
	"io/ioutil"
	"log"
)

func receiveAppDynamicsAlert(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	database.SaveAudit("APPDYNAMICS",string(body))
	database.ReceiveAppynamicsMessage()

}