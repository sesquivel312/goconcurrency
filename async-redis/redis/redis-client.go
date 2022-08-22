package rclient

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
)

type Device struct {
	Id     string
	Active bool // this is the thing to check for "in progress", true = yes
}

func SaveDevice(d Device, c *redis.Client) error {
	ctx := context.Background()

	dJson, err := json.Marshal(d)
	if err != nil {
		return err
	}

	c.Set(ctx, d.Id, dJson, 0) // !!! No result/success/error check?

	return nil
}

func RetrieveDevice(id string, c *redis.Client) Device {
	ctx := context.Background()

	j, _ := c.Get(ctx, id).Result()

	var d Device
	json.Unmarshal([]byte(j), &d)

	return d
}
