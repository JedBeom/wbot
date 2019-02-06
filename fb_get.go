package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	fb "github.com/huandu/facebook"
	"github.com/pkg/errors"
)

var (
	posts        []BasicCard
	postsErr     error
	viewMorePost = BasicCard{
		Title:       "페이스북에서 게시물 더 보기",
		Description: "학생회 페이스북 페이지를 팔로해 새 소식을 받아보세요!",
		Thumbnail: &Thumbnail{
			ImgURL: "https://raw.githubusercontent.com/JedBeom/wbot_new/master/img/view_more_fb.jpg",
		},
		Buttons: newButton("https://facebook.com/wangunstudents", "페이스북에서 보기"),
	}
)

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
		"fields":       "link,full_picture,message,story,created_time,likes.limit(1).summary(true),comments.limit(1).summary(true),shares",
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
			post.Thumbnail = &Thumbnail{ImgURL: "https://raw.githubusercontent.com/JedBeom/wbot_new/master/img/no_picture.jpg"}
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

		/*
			likes, _ := resp.Get(key + "likes.summary.total_count").(json.Number)
			comments, _ := resp.Get(key + "comments.summary.total_count").(json.Number)
			shares, _ := resp.Get(key + "shares.count").(json.Number)

			post.Social = newSocial(StoI(likes), StoI(shares), StoI(comments))
		*/

		tmpPosts = append(tmpPosts, post)
	}

	tmpPosts = append(tmpPosts, viewMorePost)

	postsErr = nil
	posts = tmpPosts
}

// Create Button object
func newButton(link, label string) (buttons []*Button) {
	buttons = append(buttons, &Button{
		Label:  label,
		URL:    link,
		Action: "webLink",
	})
	return
}

func fbLink(id string) (buttons []*Button) {
	link := "https://facebook.com/" + id
	buttons = newButton(link, "자세히 보기")
	return
}

func newSocial(likes, shares, comments int) (social *Social) {
	social = &Social{
		Like:    likes,
		Comment: comments,
		Share:   shares,
	}
	return
}

func StoI(a json.Number) int {
	b, _ := strconv.Atoi(string(a))
	return b
}
