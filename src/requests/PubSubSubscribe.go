package requests

import (
   "cloud.google.com/go/pubsub"
   "context"
   "google.golang.org/api/option"
   "log"
   "fmt"
)

func PubSubSubscribe(projectId, topicId, subscriptionId, serviceAccount string) {

    ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectId, option.WithCredentialsJSON([]byte(serviceAccount)))

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Receive messages on the subscription

	subscription := client.Subscription(subscriptionId)
    if err != nil {
        log.Fatalf("Failed to create subscription: %v", err)
    }

	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Println("Received message:", string(msg.Data))
		msg.Ack()

	})
	if err != nil {
		log.Fatalf("Failed to receive message: %v", err)
	}

}
