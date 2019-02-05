package main

import cron "gopkg.in/robfig/cron.v2"

func midnightDo() {
	go getMeals()
	go getEvents()
	go getFBPosts()
}

func workInit() {

	c := cron.New()

	// every 12 am
	if _, err := c.AddFunc("0 0 0 * * *", midnightDo); err != nil {
		panic(err)
	}

	// Every xx:14
	if _, err := c.AddFunc("0 14 * * * *", getAirqDefault); err != nil {
		panic(err)
	}

	// init
	midnightDo()

	setAirqKey()
	getAirq("연향동")

	go c.Start()
}

func getAirqDefault() {
	getAirq("연향동")
}
