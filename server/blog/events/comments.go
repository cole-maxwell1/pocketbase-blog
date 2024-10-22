package events

import (
	"net/mail"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/mailer"
)

func notifyPosterOnComment(pb *pocketbase.PocketBase) {
	// add a custom event handler to send an email to the user when a comment is created on their post
	pb.OnModelAfterCreate("comments").Add(func(e *core.ModelEvent) error {
		commentId := e.Model.GetId()

		// Get user's email from the comment's userId
		emailQuery := pb.Dao().
			DB().
			Select("users.email", "users.verified", "comments.content").
			From("users").
			InnerJoin("posts", dbx.NewExp("posts.id = comments.`postId`")).
			InnerJoin("comments", dbx.NewExp("comments.`postId` = posts.id AND users.id = posts.`userID`")).
			Where(dbx.NewExp("comments.id = {:commentId}", dbx.Params{"commentId": commentId}))
		var email string
		var content string
		var isVerified bool
		err := emailQuery.Row(&email, &isVerified, &content)

		if err != nil {
			return nil
		}

		// Only send email if user is verified
		if !isVerified {
			return nil
		}


		message := &mailer.Message{
			From: mail.Address{
				Address: pb.Settings().Meta.SenderAddress,
				Name:    pb.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: email}},
			Subject: "There's a new comment on your post!",
			HTML:    content,
		}

		return pb.NewMailClient().Send(message)
	})
}
