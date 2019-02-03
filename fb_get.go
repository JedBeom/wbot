package main

import (
	"fmt"
	"log"

	fb "github.com/huandu/facebook"
)

var (
	posts    []BasicCard
	postsErr error
)

func getFBPosts() {

	limit := 5
	postsErr = nil

	resp, err := fb.Get("/v3.2/wangunstudents/posts", fb.Params{
		"access_token": config.FBKey,
		"limit":        limit,
	})
	if err != nil {
		log.Println(err)
		postsErr = err
		return
	}

	var tmpPosts []BasicCard
	for i := 0; i < limit; i++ {

		key := fmt.Sprintf("data.%d.", i)
		var post BasicCard

		if id := resp.Get(key + "id"); id != nil {
			post.Buttons = fbLink(id.(string))

			resp, err := fb.Get("/v3.2/"+id.(string), fb.Params{
				"access_token": config.FBKey,
				"fields":       "picture",
			})
			if err != nil {
				log.Println("Error while getting a facebook post:", err)
				return
			}

			if picture := resp.Get("picture"); picture != nil {
				post.Thumbnail.ImgURL = picture.(string)
			}

		} else {
			log.Println("Error while getting facebook posts")
			postsErr = err
			return
		}

		if msg := resp.Get(key + "message"); msg != nil {
			post.Title = msg.(string)
		} else {
			story := resp.Get(key + "story")
			if story == nil {
				continue
			}

			post.Title = story.(string)
		}

		if timeStr := resp.Get(key + "created_time"); timeStr != nil {

			post.Description = "게시 날짜: " + timeStr.(string)[:10]

		}

		tmpPosts = append(tmpPosts, post)

	}

	postsErr = nil
	posts = tmpPosts
}

func fbLink(id string) (buttons []Button) {
	link := "https://facebook.com/" + id
	buttons = append(buttons, Button{
		Label:  "자세히 보기",
		URL:    link,
		Action: "webLink",
	})
	return
}
