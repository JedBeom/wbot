package main

import (
	"context"
	"log"
	"net/http"

	"github.com/JedBeom/wbot_new/model"
)

//MiddlewareHistory
func MiddlewareHistory(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h, err := ParseHistory(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		var user model.User
		if user, err = model.GetUserByID(db, h.UserID); err != nil || user.ID == "" {
			user.ID = h.UserID
			err = user.Create(db)

			if err != nil {
				log.Println(err)
				w.WriteHeader(500)
				return
			}
		}

		h.User = &user
		err = h.Create(db)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		ctx := context.WithValue(r.Context(), "history", h)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
