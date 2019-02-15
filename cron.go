package main

import cron "gopkg.in/robfig/cron.v2"

func midnightJob() {
	go getMeals()
	go getEvents()
}

func workInit() {

	c := cron.New()

	// every 12 am
	if _, err := c.AddFunc("@midnight", midnightJob); err != nil {
		panic(err)
	}

	// Every xx:14
	if _, err := c.AddFunc("0 10 * * * *", getAirqDefault); err != nil {
		panic(err)
	}

	if _, err := c.AddFunc("0 0/30 * * * *", getFBPosts); err != nil {
		panic(err)
	}

	// init
	go midnightJob()
	go getFBPosts()

	setAirqKey()
	getAirq("연향동")

	go c.Start()
}

func getAirqDefault() {
	getAirq("연향동")
}
