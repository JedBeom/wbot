package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	fb "github.com/huandu/facebook"
	"github.com/pkg/errors"
)

var (
	posts        []BasicCard
	viewMorePost = BasicCard{
		"페이스북에서 게시물 더 보기",
		"학생회 페이스북 페이지를 팔로해 새 소식을 받아보세요!",
		nil,
		nil,
		&Thumbnail{
			ImgURL: "https://raw.githubusercontent.com/JedBeom/wbot_new/master/img/view_more_fb.jpg",
		},
		[]Button{{"webLink", "페이스북에서 보기", "https://facebook.com/wangunstudents"}},
	}
)

func getFBPosts() {

	// Post limit
	limit := 5
	// Page ID
	pageID := "1557399350946249"

	// Get
	resp, err := fb.Get("/v3.2/"+pageID+"/posts", fb.Params{
		"access_token": config.FBKey,
		"limit":        limit,
		"fields":       "link,full_picture,message,story,created_time,likes.limit(1).summary(true),comments.limit(1).summary(true),shares",
	})
	// Check error
	if err != nil {
		log.Println(err)
		return
	}

	tmpPosts := make([]BasicCard, 0, 6)
	for i := 0; i < limit; i++ {
		post, err := analysisPost(i, &resp)
		if err != nil {
			log.Println("Error while parsing post:", err)
			continue
		}

		tmpPosts = append(tmpPosts, post)
	}

	if len(tmpPosts) == 0 {
		return
	}
	tmpPosts = append(tmpPosts, viewMorePost)
	posts = tmpPosts
}

func analysisPost(i int, resp *fb.Result) (post BasicCard, err error) {

	key := fmt.Sprintf("data.%d.", i)

	// get ID
	if id, ok := resp.Get(key + "id").(string); ok {
		fbLink(&post.Buttons, id)
	} else {
		err = errors.New("Can't parse id")
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
			err = errors.New("Cannot get a message nor a story")
			return
		}

		post.Title = story
	}

	link, ok := detectFirstLink(post.Title)
	if ok {
		newButton(&post.Buttons, link, "게시물 내 링크 바로가기")
	}

	likes, _ := resp.Get(key + "likes.summary.total_count").(json.Number)
	comments, _ := resp.Get(key + "comments.summary.total_count").(json.Number)
	// shares, _ := resp.Get(key + "shares.count").(json.Number)

	// get created_at
	if timeStr, ok := resp.Get(key + "created_time").(string); ok {
		format := "📅 %s\n👍 %d개, 💬 %d개"
		post.Description = fmt.Sprintf(format, timeStr[:10], StoI(likes), StoI(comments))
	}

	return
}

// Create Button object
func newButton(buttons *[]Button, link, label string) {
	*buttons = append(*buttons, Button{
		Label:  label,
		URL:    link,
		Action: "webLink",
	})
	return
}

func fbLink(buttons *[]Button, id string) {
	link := "https://facebook.com/" + id
	newButton(buttons, link, "자세히 보기")
	return
}

/*
func newSocial(likes, shares, comments int) (social *Social) {
	social = &Social{
		Like:    likes,
		Comment: comments,
		Share:   shares,
	}
	return
}
*/

func StoI(a json.Number) int {
	b, _ := strconv.Atoi(string(a))
	return b
}

func detectFirstLink(a string) (link string, ok bool) {
	index := strings.Index(a, "http")
	if index == -1 {
		return
	}

	spaceSplit := strings.Split(a[index:], " ")
	if len(spaceSplit) == 0 {
		return
	}

	link = spaceSplit[0]
	ok = true
	return
}
