package main

import (
	"fmt"
	fb "github.com/huandu/facebook"
	"github.com/pkg/errors"
	"log"
)

var (
	posts        []BasicCard
	postsErr     error
	viewMorePost = BasicCard{
		Title:       "페이스북에서 게시물 더 보기",
		Description: "학생회 페이스북 페이지를 팔로해 새 소식을 받아보세요!",
		Thumbnail: &Thumbnail{
			ImgURL: "https://raw.githubusercontent.com/JedBeom/wbot_new/facebook/img/view_more_fb.jpg",
		},
	}
)

func init() {
	viewMorePost.Buttons = append(viewMorePost.Buttons, &Button{
		Action: "페이스북 가기",
		Label:  "webLink",
		URL:    "https://facebook.com/wangunstudents",
	})
}

func getFBPosts() {

	// Post limit
	limit := 5
	// Page ID
	pageID := "1557399350946249"
	// Init error
	postsErr = nil

	// Get
	resp, err := fb.Get("/v3.2/"+pageID+"/posts", fb.Params{
		"access_token": config.FBKey,
		"limit":        limit,
		"fields":       "link,full_picture,message,story,created_time",
	})
	// Check error
	if err != nil {
		log.Println(err)
		postsErr = err
		return
	}

	tmpPosts := make([]BasicCard, 0, 6)
	for i := 0; i < limit; i++ {

		var post BasicCard
		key := fmt.Sprintf("data.%d.", i)

		// get ID
		if id, ok := resp.Get(key + "id").(string); ok {
			post.Buttons = fbLink(id)
		} else {
			log.Println("Error while getting facebook posts")
			postsErr = errors.New("Can't parse id")
			return
		}

		// get picture url
		if picture, ok := resp.Get(key + "full_picture").(string); ok {
			post.Thumbnail = &Thumbnail{ImgURL: picture}
		} else {
			post.Thumbnail = &Thumbnail{ImgURL: "https://raw.githubusercontent.com/JedBeom/wbot_new/facebook/img/no_picture.jpg"}
		}

		// get message
		if msg, ok := resp.Get(key + "message").(string); ok {
			post.Title = msg
			// if not
		} else {
			// get story instead
			story, ok := resp.Get(key + "story").(string)
			if !ok {
				// if there's no story and message both
				continue
			}

			post.Title = story
		}

		// get created_at
		if timeStr, ok := resp.Get(key + "created_time").(string); ok {

			post.Description = "게시 날짜: " + timeStr[:10]

		}

		tmpPosts = append(tmpPosts, post)
	}

	tmpPosts = append(tmpPosts, viewMorePost)

	postsErr = nil
	posts = tmpPosts
}

// Create Button object
func fbLink(id string) (buttons []*Button) {
	link := "https://facebook.com/" + id
	buttons = append(buttons, &Button{
		Label:  "자세히 보기",
		URL:    link,
		Action: "webLink",
	})
	return
}
