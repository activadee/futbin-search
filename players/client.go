package players

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type Client interface {
	Get(opt *Options) ([]FilteredPlayer, error)
	TOTW() ([]FilteredPlayer, error)
	Popular() ([]FilteredPlayer, error)
	Latest() ([]FilteredPlayer, error)
	GetByName(name string) ([]NamePlayer, error)
	GetById(id int, platform string) ([]Player, error)
}

func DefaultClient() Client {
	return NewClient(http.DefaultClient)
}

func NewClient(c *http.Client) Client {
	return &client{client: c}
}

type client struct {
	client *http.Client
}

func (c client) Get(opt *Options) ([]FilteredPlayer, error) {
	u, _ := url.Parse("https://www.futbin.org/futbin/api/23/getFilteredPlayers")
	pq := newPlayerQuery(opt)
	q, err := query.Values(pq)
	if err != nil {
		return nil, err
	}
	u.RawQuery = q.Encode()
	return c.get(u.String())
}

func (c client) TOTW() ([]FilteredPlayer, error) {
	return c.get("https://www.futbin.org/futbin/api/23/currentTOTW")
}

func (c client) Popular() ([]FilteredPlayer, error) {
	return c.get("https://www.futbin.org/futbin/api/23/getPopularPlayers")
}
func (c client) GetByName(name string) ([]NamePlayer, error) {
	return c.getByName(fmt.Sprintf("https://www.futbin.org/futbin/api/searchPlayersByName?playername=%s&year=23", name))
}
func (c client) GetById(id int, platform string) ([]Player, error) {
	return c.getById(fmt.Sprintf("https://www.futbin.org/futbin/api/23/fetchPlayerInformation?ID=%d&ap=1.0.1&platform=%s", id, platform))
}
func (c client) Latest() ([]FilteredPlayer, error) {
	return c.get("https://www.futbin.org/futbin/api/23/newPlayers")
}

func (c client) get(url string) ([]FilteredPlayer, error) {
	r, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	var data playersDataFilter
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Players, err
}
func (c client) getByName(url string) ([]NamePlayer, error) {
	r, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	var data playersDataName
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Players, err
}
func (c client) getById(url string) ([]Player, error) {
	r, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	var data playersData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Players, err
}
