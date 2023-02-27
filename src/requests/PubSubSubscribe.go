package requests

import (
	"cloud.google.com/go/pubsub"
	"context"
	json2 "encoding/json"
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"google.golang.org/api/option"
	"log"
)

func PubSubSubscribe(projectId string, topicId string, subscriptionId string, serviceAccount string, listeningFunction func(domains.Pipeline, bool)) bool {

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectId, option.WithCredentialsJSON([]byte(serviceAccount)))

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return false
	}

	// Receive messages on the subscription
	subscription := client.Subscription(subscriptionId)
	if err != nil {
		log.Fatalf("Failed to create subscription: %v", err)
		return false
	}

	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var pipeline domains.Pipeline
		error := json2.Unmarshal(msg.Data, &pipeline)
		if error == nil {
			go listeningFunction(pipeline, true)
			msg.Ack()
		} else {
			go listeningFunction(domains.Pipeline{}, false)
			msg.Nack()
		}
	})

	if err != nil {
		log.Fatalf("Failed to receive message: %v", err)
		return false
	}
	return true
}
