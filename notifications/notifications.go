package notifications

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var notificationLimit = 0

type nKey string

// Context pulls out the notification handler from the context
func Context(ctx context.Context) *Handler {
	return ctx.Value("notifications").(*Handler)

}

// Handler handles the configuration for notifications. Please update the config based off
// of the pushover.net configuration settings
type Handler struct {
	Token string
	User  string
}

type notifications string

var notifier = notifications("notifications")

// Init initializes the notification handler with user information and adds it to the context
func Init(ctx context.Context, conf map[string]string) context.Context {
	h := &Handler{
		Token: conf["token"],
		User:  conf["user"],
	}

	return context.WithValue(ctx, notifier, h)

}

// Send sends a notification to your phone via the app PushOver.
// more info here: https://pushover.net/
//
// pushover has a limit of 7,500 notifications per month
// as a result, this function is limited to 100 notifications per day to avoid hitting the cap
func (h *Handler) Send(message string) {
	if notificationLimit > 100 {
		log.Println("failed to send notification, limit reached: ", message)
	}
	notificationLimit++
	log.Println("sending notification: ", message)
	form := url.Values{}
	form.Add("token", h.Token)
	form.Add("message", message)
	form.Add("user", h.User)

	client := &http.Client{}
	r, err := http.NewRequest("POST", "https://api.pushover.net/1/messages.json", strings.NewReader(form.Encode()))
	if err != nil {
		log.Println(err)
		return
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.Status)

}
