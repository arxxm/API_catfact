package catfacts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	URLstring  = "https://catfact.ninja/"
	listBreeds = "breeds"
	listFacts  = "facts"
	randomFact = "fact"
)

type httpGetClient interface {
	Get(url string) (*http.Response, error)
}

type Client struct {
	baseURL    *url.URL
	httpClient httpGetClient
	pageSize   int
}

func NewClientWithStringURL(URLstring string) (*Client, error) {
	baseUrl, err := url.Parse(URLstring)
	if err != nil {
		return nil, err
	}
	return NewClientWithURL(baseUrl), nil
}

func NewClientWithURL(baseURL *url.URL) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
		pageSize:   10,
	}
}

func NewClient() *Client {
	client, err := NewClientWithStringURL(URLstring)
	if err != nil {
		panic(err)
	}
	return client
}

func (c *Client) WithPageSize(pageSize int) *Client {
	c.pageSize = pageSize
	return c
}

func (c *Client) GetRandomFact() (*CatFact, error) {
	u := fmt.Sprintf("%s%s", c.baseURL, randomFact)
	res, err := c.httpClient.Get(u)
	if err != nil {
		return nil, err
	}

	var parsed CatFact
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&parsed); err != nil {
		return nil, err
	}

	return &parsed, nil
}

type pageHandler func(page *json.Decoder) (*pagination, error)

func (c *Client) paginationGet(urlStr string, handle pageHandler) error {
	currentPage := 0
	lastPage := 1

	for currentPage != lastPage {
		currentPage++
		res, err := c.httpClient.Get(fmt.Sprintf("%s?limit=%d&page=%d", urlStr, c.pageSize, currentPage))
		if err != nil {
			return err
		}
		decoder := json.NewDecoder(res.Body)
		pagination, err := handle(decoder)
		if err != nil {
			return err
		}
		currentPage = pagination.CurrentPage
		lastPage = pagination.LastPage
	}

	return nil
}

func (c *Client) ListAllFacts() ([]CatFact, error) {
	u := fmt.Sprintf("%s%s", c.baseURL.String(), listFacts)
	var allFacts []CatFact

	err := c.paginationGet(u, func(page *json.Decoder) (*pagination, error) {
		var parsed facts
		err := page.Decode(&parsed)
		if err != nil {
			return nil, err
		}
		allFacts = append(allFacts, parsed.Data...)
		return &parsed.pagination, nil
	})
	if err != nil {
		return nil, err
	}

	return allFacts, nil
}
