package notifications

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var notificationLimit = 0

// Send ...
func Send(message string) {
	if notificationLimit > 100 {
		fmt.Println("failed to send notification, limit reached: ", message)
	}
	notificationLimit++
	fmt.Println("sending notification: ", message)
	form := url.Values{}
	form.Add("token", "a53g47tydze4fm4wnndjtutkdk8fs6")
	form.Add("message", message)
	form.Add("user", "ukr9b4nitkkff3t6e6su4x36vamngx")

	client := &http.Client{}
	r, err := http.NewRequest("POST", "https://api.pushover.net/1/messages.json", strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Status)

}
