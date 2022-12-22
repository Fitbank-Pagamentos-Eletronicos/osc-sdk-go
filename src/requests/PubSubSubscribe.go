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
	projectID := projectId
	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsJSON([]byte(serviceAccount)))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Receive messages on the subscription
	subscriptionID := subscriptionId
	subscription := client.Subscription(subscriptionID)
	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Println("Received message:", string(msg.Data))
		msg.Ack()

	})
	if err != nil {
		log.Fatalf("Failed to receive message: %v", err)
	}



//    ctx := context.Background()
//
//    pubsubClient, err := pubsub.NewClient(ctx, projectId, option.WithCredentialsJSON([]byte(serviceAccount)))
//
//    if err != nil {
//       log.Fatal(err)
//    }
//
//    topic, err := pubsubClient.CreateTopic(ctx, topicId)
//
//    res := topic.Publish(ctx, &pubsub.Message{
//       Data: []byte("Hello World"),
//    })
//
//    res.Get(ctx)
//
//    topic.Stop()
//
//    sub, err := pubsubClient.CreateSubscription(ctx, subscriptionId, pubsub.SubscriptionConfig{
//       Topic: topic,
//    })
//
//    err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
//       log.Printf("Got message: %q", string(msg.Data))
//       msg.Ack()
//    })
//
//    if err != nil {
//       log.Fatal(err)
//    }

}
