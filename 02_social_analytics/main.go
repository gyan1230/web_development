package main

// Initialize the client with your api key and secret.
// Error checks are removed for brevity, but always check err != nil
import (
	"fmt"
	"log"

	"gopkg.in/GetStream/stream-go2.v1"
)

func main() {

	//Creating a client

	client, err := stream.NewClient(
		"er5ksmzxe39e",
		"kpu7wskyvmzq8uekqjumpxywezap8xt5nx89sq5qmjrmn54qdgu5cvhtc4ddd6jj",
	)

	if err != nil {
		log.Println(err)
	}

	client2, err := stream.NewClient(
		"mgksm5aup63t",
		"mx2bs7s6hj36uvf83j4ydbp39nn35c48zuk4yfqey2d6ucsaj63ahnbjj26mh4av",
	)

	if err != nil {
		log.Println(err)
	}

	flat := client.FlatFeed("user", "123")

	resp, err := flat.GetActivities()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Duration:", resp.Duration)
	fmt.Println("Next:", resp.Next)
	fmt.Println("Activities:")
	for _, activity := range resp.Results {
		fmt.Println(activity)
	}

	aggregated := client.AggregatedFeed("New_Aggr_Feed", "gyan1230")

	resp2, err := aggregated.GetActivities()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Duration:", resp2.Duration)
	fmt.Println("Next:", resp2.Next)
	fmt.Println("Groups:")
	for _, group := range resp2.Results {
		//	fmt.Println("Group:", group.Name, "ID:", group.ID, "Verb:", group.Verb)
		fmt.Println("Activities:", group.ActivityCount, "Actors:", group.ActorCount)
		for _, activity := range group.Activities {
			fmt.Println(activity)
		}
	}

	notification := client2.NotificationFeed("notification_feed", "gyan12")

	resp3, err := notification.GetActivities()
	if err != nil {
		// ...
	}

	fmt.Println("Duration:", resp3.Duration)
	fmt.Println("Next:", resp3.Next)
	fmt.Println("Unseen:", resp3.Unseen, "Unread:", resp3.Unread)
	fmt.Println("Groups:")
	for _, group := range resp3.Results {
		fmt.Println("Group:", group.Group, "ID:", group.ID, "Verb:", group.Verb)
		fmt.Println("Seen:", group.IsSeen, "Read:", group.IsRead)
		fmt.Println("Activities:", group.ActivityCount, "Actors:", group.ActorCount)
		for _, activity := range group.Activities {
			fmt.Println(activity)
		}
	}

	// Get the feed object
	gyanFeed := client.FlatFeed("New_Aggr_Feed", "gyan1230")
	// Add the activity to the feed
	gyanFeed.AddActivity(stream.Activity{
		Actor:  gyanFeed.ID(),
		Verb:   "post",
		Object: "1",
		Extra: map[string]interface{}{
			"post": "Hello world",
		},
	})
	fmt.Println(gyanFeed.ID())

}
