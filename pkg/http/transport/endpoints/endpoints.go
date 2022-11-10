package endpoints

import (
	"context"

	service "github.com/Antoha2/tlgrmBot1/service/windService"
	"github.com/go-kit/kit/endpoint"
)

func MakeHistoryEndpoints(s service.Service) *Endpoints {
	return &Endpoints{
		History: MakeHistoryEndpoint(s),
	}
}

type Endpoints struct {
	History endpoint.Endpoint
}

type DownloadRequest struct {
	UserId int `json:"user_id"`
	//	Count int    `json:"count"`
}

type DownloadResponse struct {
	History []Message `json:"response"`
}

type Message struct {
	MessageId int `json:"id" gorm:"column:id"`
	//UserName  string `json:"user_name"`
	//UserId   int    `json:"user_id"`
	ChatId   int64  `json:"chat"`
	Text     string `json:"text"`
	Response string `json:"bot_response"`
}

func MakeHistoryEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		webData := request.(DownloadRequest)

		servMessage, err := s.GetHistory(ctx, webData.UserId)
		httpData := make([]Message, len(servMessage))
		for index, data := range servMessage {
			httpData[index].MessageId = data.MessageId
			httpData[index].ChatId = data.ChatId
			httpData[index].Text = data.Text
			httpData[index].Response = data.Response
		}
		if err != nil {
			return DownloadResponse{History: httpData}, err
		}
		return DownloadResponse{History: httpData}, nil
	}
}
