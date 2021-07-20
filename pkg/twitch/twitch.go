package twitch

import (
	"fmt"
	"github.com/nicklaw5/helix"
)

const maxChunkSize = 100

type Client struct {
	helixClient *helix.Client
}

func NewTwitchClient(clientId, clientSecret string) (*Client, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:     clientId,
		ClientSecret: clientSecret,
	})
	if err != nil {
		return nil, fmt.Errorf("create new client: %w", err)
	}
	token, err := client.RequestAppAccessToken(nil)
	if err != nil {
		return nil, fmt.Errorf("request app access token: %w", err)
	}
	client.SetAppAccessToken(token.Data.AccessToken)
	return &Client{client}, nil
}

func (c *Client) FindUsers(streamers []string) ([]helix.User, error) {
	var users []helix.User
	for i := 0; i < len(streamers); i += maxChunkSize {
		end := i + maxChunkSize
		if end > len(streamers) {
			end = len(streamers)
		}
		resultChunk, err := c.helixClient.GetUsers(&helix.UsersParams{Logins: streamers[i:end]})
		if err != nil {
			return nil, fmt.Errorf("get users: %w", err)
		}
		users = append(users, resultChunk.Data.Users...)
	}
	return users, nil
}
