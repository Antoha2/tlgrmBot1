package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Antoha2/tlgrmBot1/pkg/http/transport/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func (httpImpl *httpImpl) StartHTTP() error {

	Options := []httptransport.ServerOption{}

	HistoryHandler := httptransport.NewServer(
		endpoints.MakeHistoryEndpoint(httpImpl.service),
		decodeMakeHistoryRequest,
		encodeResponse,
		Options...,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path("/history").Handler(HistoryHandler)

	httpImpl.server = &http.Server{Addr: httpImpl.config.HTTP.HostAddr}              //:8180
	log.Printf(" Запуск HTTP-сервера на http://127.0.0.1%s\n", httpImpl.server.Addr) //:8180

	if err := http.ListenAndServe(httpImpl.server.Addr, r); err != nil {
		log.Println(err)
	}
	return nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func decodeMakeHistoryRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.DownloadRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func (httpImpl *httpImpl) Stop() {

	if err := httpImpl.server.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
}
