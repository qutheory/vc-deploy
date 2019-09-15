package api

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"sync"
	"time"
)

func Listen(activityID string) {
	//time.Sleep(4 * time.Second)
	u := "wss://api.v2.vapor.cloud/v2/activity/activities/" + activityID + "/channel"

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	defaultCloseHandler := c.CloseHandler()
	c.SetCloseHandler(func(code int, text string) error {
		result := defaultCloseHandler(code, text)
		fmt.Println("Disconnected from server ", result)

		os.Exit(0)

		return result
	})

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go readPump(&waitgroup, c)
	waitgroup.Wait()
}

func readPump(waitgroup *sync.WaitGroup, c *websocket.Conn) {

	c.SetReadDeadline(time.Now().Add(60 * time.Second))

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(string(message))
	}

	waitgroup.Done()
}