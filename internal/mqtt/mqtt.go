package mqtt

import (
	"fmt"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func Connect() {
	const TOPIC = "mytopic/test"
	var broker = "replace" //broker.emqx.io"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("replace")
	opts.SetPassword("replace")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)

	if token := client.Subscribe(TOPIC, 0, func(client mqtt.Client, msg mqtt.Message) {
		if string(msg.Payload()) != "mymessage" {
			//t.Fatalf("want mymessage, got %s", msg.Payload())
		}
		wg.Done()
	}); token.Wait() && token.Error() != nil {
		//t.Fatal(token.Error())
	}

	if token := client.Publish(TOPIC, 0, false, "mymessage"); token.Wait() && token.Error() != nil {
		//t.Fatal(token.Error())
	}
	wg.Wait()
}
