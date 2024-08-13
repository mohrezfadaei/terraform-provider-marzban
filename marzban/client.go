package marzban

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
		Token:      token,
	}
}

func (c *Client) CreateNode(name, address string, port, apiPort int, usageCoefficient float64, addAsNewHost bool) (int, error) {
	reqBody, err := json.Marshal(map[string]interface{}{
		"name":              name,
		"address":           address,
		"port":              port,
		"api_port":          apiPort,
		"usage_coefficient": usageCoefficient,
		"add_as_new_host":   addAsNewHost,
	})
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/api/node", bytes.NewBuffer(reqBody))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var respData struct {
		NodeID int `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return 0, err
	}

	return respData.NodeID, nil
}

func (c *Client) GetNode(nodeID string) (*Node, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/node/%s", c.BaseURL, nodeID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var node Node
	if err := json.NewDecoder(resp.Body).Decode(&node); err != nil {
		return nil, err
	}

	return &node, nil
}

func (c *Client) UpdateNode(nodeID, name, address string, port, apiPort int, usageCoefficient float64, addAsNewHost bool) error {
	reqBody, err := json.Marshal(map[string]interface{}{
		"name":              name,
		"address":           address,
		"port":              port,
		"api_port":          apiPort,
		"usage_coefficient": usageCoefficient,
		"add_as_new_host":   addAsNewHost,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/node/%s", c.BaseURL, nodeID), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) DeleteNode(nodeID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/node/%s", c.BaseURL, nodeID), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

type Node struct {
	Name             string  `json:"name"`
	Address          string  `json:"address"`
	Port             int     `json:"port"`
	ApiPort          int     `json:"api_port"`
	UsageCoefficient float64 `json:"usage_coefficient"`
	AddAsNewHost     bool    `json:"add_as_new_host"`
}
