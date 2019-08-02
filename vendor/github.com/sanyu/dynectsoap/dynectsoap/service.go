package dynectsoap

import (
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	PollingInterval = 2 * time.Second
	ApiUrl          = "https://api2.dynect.net/SOAP/"
)

type dynectsoap struct {
	client *Client
}

func NewDynect(client *Client) *dynectsoap {
	return &dynectsoap{
		client: client,
	}
}

func (service *dynectsoap) GetJobRetry(request *GetJobRequestType, response *GetJobResponseType) error {
	var err error

	for {
		select {
		case <-time.After(PollingInterval):
			err = service.client.Call(ApiUrl, request, response)
			if err != nil {
				return err
			}

			realResponse := response.Data.(JobResponseInterface)
			status := realResponse.status()

			switch status {
			case "incomplete":
				//Loop around again
				log.Debugf("Retrieving Job %v again, status is incomplete", response.Job_id)
				continue
			default:
				return nil
			}
		}
	}

	return nil
}

func (service *dynectsoap) Do(request, response interface{}) error {
	err := service.client.Call(ApiUrl, request, response)
	return err
}
