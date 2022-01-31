package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dghubble/oauth1"
)

const url = "https://api.twitter.com/2"

type Client struct {
	credentials *Credentials
	http        *http.Client
}

type Credentials struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	AccessToken    string `json:"accessToken"`
	TokenSecret    string `json:"accessTokenSecret"`
	BearerToken    string `json:"bearerToken"`
}

type CreateTwitterRequest struct {
	Text string `json:"text"`
}

type CreateTwitterResponse struct {
	Data *CreateTwitterDataResponse `json:"data"`
}

type CreateTwitterDataResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func New(creds *Credentials) *Client {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	client := config.Client(oauth1.NoContext, oauth1.NewToken(creds.AccessToken, creds.TokenSecret))

	return &Client{
		credentials: creds,
		http:        client,
	}
}

func (c *Client) CreateTwitter(r *CreateTwitterRequest) (*CreateTwitterResponse, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(b))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tweets", url), payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 201 {
		return nil, errors.New(string(body))
	}

	var resp *CreateTwitterResponse

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
