package model

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type Checker struct {
	endpoint string
	grace    int
	status   int
}

func NewChecker(endpoint string, grace int) *Checker {
	return &Checker{
		endpoint: endpoint,
		grace:    grace,
		status:   0,
	}
}

func (c *Checker) Check() error {
	if err := c.access(); err != nil && c.status == 0 {
		c.status = 1
		return err
	}
	c.status = 0
	return nil
}

func (c *Checker) access() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.grace)*time.Second)
	defer cancel()
	req, err := http.NewRequest("GET", c.endpoint, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("status code is not 200")
	}
	return nil
}
