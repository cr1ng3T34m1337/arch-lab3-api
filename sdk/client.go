package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cr1ng3T34m1337/arch-lab3-api/api/dataProviders"
	"github.com/cr1ng3T34m1337/arch-lab3-api/api/httpHandlers"
)

type Client struct {
	BaseUrl string
}

func (c *Client) ListBalancers() ([]dataProviders.Balancer, error) {
	res, err := http.Get(c.BaseUrl + "/balancers")
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
	defer res.Body.Close()
	var result []dataProviders.Balancer
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) SetWorkingStatus(id int, status string) error {
	reqStruct := httpHandlers.SetWorkingStatusRequest{Id: id, WorkingStatus: status}
	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(reqStruct)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, c.BaseUrl+"/machines", reqBody)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
	return nil
}
