package telegram

import (
	"context"
	"encoding/json"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/weAutomateEverything/go2hal/gokit"
	"github.com/weAutomateEverything/go2hal/machineLearning"
	"net/http"
)

/*
MakeHandler returns a HTTP Restul endpoint to handle user requests
*/
func MakeHandler(service Service, logger kitlog.Logger, ml machineLearning.Service) http.Handler {
	opts := gokit.GetServerOpts(logger, ml)

	//requestAuth := kithttp.NewServer(gokitjwt.NewParser(gokit.GetJWTKeys(),jwt.SigningMethodHS256,gokitjwt.StandardClaimsFactory)( makeTelegramAuthRequestEndpoint(service)),
	//gokit.DecodeString, gokit.EncodeResponse, opts...)
	requestAuth := kithttp.NewServer(makeTelegramAuthRequestEndpoint(service),
		decodeAuthRequest, gokit.EncodeResponse, opts...)

	authpoll := kithttp.NewServer(makeTelegramAuthPollEndpoint(service),
		decodeAuthPoll, encodeAuthResoinse)

	r := mux.NewRouter()

	// swagger:operation POST /api/telegram/auth auth RequestToken
	//
	// Initialises a reuqest for a token for a telegram room.
	//
	// someone in the room will need to approve the request.
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: message
	//   in: body
	//   description: request
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/authRequestObject"
	// responses:
	//   '200':
	//     description: Message deleted successfully
	//     schema:
	//       "$ref": "#/definitions/authResponse"
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/definitions/errorResponse"
	r.Handle("/api/telegram/auth", requestAuth).Methods("POST")

	// swagger:operation GET /api/telegram/auth/{id} auth PollToken
	//
	// Polls to see if the access request has been approved in the telegram group.
	//
	// Token will be in the AUTHORIZATION header
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: Key
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: Message deleted successfully
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/definitions/errorResponse"
	r.Handle("/api/telegram/auth/{id}", authpoll).Methods("GET")

	return r

}

func decodeAuthRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//jwt := ctx.Value(gokitjwt.JWTClaimsContextKey).(CustomClaims)
	v := authRequestObject{}
	err := json.NewDecoder(r.Body).Decode(&v)
	return v, err
}

func decodeAuthPoll(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return vars["id"], nil

}

type errorer interface {
	error() error
}

func encodeAuthResoinse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if response == nil {
		return nil
	}
	if e, ok := response.(errorer); ok && e.error() != nil {
		gokit.EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Add("Authorization", "Bearer "+response.(string))
	return nil
}
