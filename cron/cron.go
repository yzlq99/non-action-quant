package cron

import (
	"github.com/robfig/cron/v3"
)

type cronTab struct {
	c *cron.Cron
}

var CronTab *cronTab

func InitCron() {
	CronTab = &cronTab{
		c: cron.New(),
	}
}

type JobInterface interface {
	Spec() string
	Run()
}

func (c *cronTab) AddJob(job JobInterface) error {
	_, err := c.c.AddFunc(job.Spec(), job.Run)
	return err
}

func (c *cronTab) Run() {
	c.c.Run()
}
