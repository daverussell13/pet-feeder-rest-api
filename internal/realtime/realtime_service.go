package realtime

import (
	"context"
	"github.com/daverussell13/Pet_Feeder_Rest_API/infrastructures/mqtt"
	"strconv"
	"time"
)

type service struct {
	mqtt        *mqtt.Mqtt
	feedTimeout time.Duration
}

func NewService(mqtt *mqtt.Mqtt) Service {
	return &service{
		mqtt:        mqtt,
		feedTimeout: time.Duration(3) * time.Second,
	}
}

func (s *service) RealtimeFeed(c context.Context, request *FeedRequest) (*FeedResponse, error) {
	feedAmount := request.FeedAmount
	mqttClient := s.mqtt.GetClient()

	topic := s.mqtt.GetTopic().FeedTopic + "/" + request.DeviceID
	token := mqttClient.Publish(topic, 2, false, strconv.Itoa(feedAmount))
	token.WaitTimeout(s.feedTimeout)

	if token.Error() != nil {
		return nil, token.Error()
	}

	return &FeedResponse{
		DeviceID:   request.DeviceID,
		FeedAmount: feedAmount,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}