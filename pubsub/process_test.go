package pubsub

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

func topicFunc(args ...interface{}) {
	topic := args[0].(string)
	subKey := args[1]
	log.Printf("topic: %s, subKey: %v\n", topic, subKey)
	for _, v := range args[2:] {
		log.Println(v)
	}
}

func TestUnSafetyPubSub(t *testing.T) {
	obs, err := NewObServer(false)
	if err != nil {
		log.Printf("TestUnSafetyPubSub.NewObserver failed, err: %v\n", err)
		return
	}

	topic := "test"
	subKey := 1
	if err := obs.Subscribe(topic, subKey, topicFunc, topic, subKey, 100, "acc"); err != nil {
		log.Printf("TestUnSafetyPubSub.Subscribe failed, err: %v\n", err)
		return
	}

	if err := obs.Publish(topic, false, "ddd", 200); err != nil {
		log.Printf("TestUnSafetyPubSub.Publish failed, err: %v\n", err)
		return
	}

	if err := obs.Unsubscribe(topic, subKey); err != nil {
		log.Printf("TestUnSafetyPubSub.Unsubscribe failed, err: %v\n", err)
		return
	}
	log.Println("TestUnSafetyPubSub succeeded")
}

func TestSafetyPubSub(t *testing.T) {
	obs, err := NewObServer(true)
	if err != nil {
		log.Printf("TestSafetyPubSub failed, err: %v\n", err)
		return
	}

	count := 5
	topicKey := "co_topic_test"

	var waitGroup sync.WaitGroup
	// subscribe test
	waitGroup.Add(count * 2)
	for i := 0; i < count; i++ {
		go func(idx int, wg *sync.WaitGroup) {
			if err := obs.Subscribe(topicKey, idx, topicFunc, topicKey, idx); err != nil {
				log.Printf("TestSafetyPubSub.Subscribe failed, err: %v\n", err)
			}
			wg.Done()
		}(i, &waitGroup)

		go func(idx int, wg *sync.WaitGroup) {
			coTopicKey := fmt.Sprintf("%s_%d", topicKey, idx)
			if err := obs.Subscribe(coTopicKey, idx, topicFunc); err != nil {
				log.Printf("TestSafetyPubSub.Subscribe failed, err: %v\n", err)
			}
			wg.Done()
		}(i, &waitGroup)
	}
	waitGroup.Wait()

	// publish test
	if err := obs.Publish(topicKey, false); err != nil {
		log.Printf("TestSafetyPubSub.Publish failed, err: %v", err)
		return
	}

	waitGroup.Add(count)
	for i := 0; i < count; i++ {
		go func(idx int, wg *sync.WaitGroup) {
			coTopicKey := fmt.Sprintf("%s_%d", topicKey, idx)
			if err := obs.Publish(coTopicKey, false, coTopicKey, idx); err != nil {
				log.Printf("TestSafetyPubSub.Publish failed, err: %v\n", err)
			}
			wg.Done()
		}(i, &waitGroup)

	}
	waitGroup.Wait()

	// unsubscribe test
	waitGroup.Add(count * 2)
	for i := 0; i < count; i++ {
		go func(idx int, wg *sync.WaitGroup) {
			if err := obs.Unsubscribe(topicKey, idx); err != nil {
				log.Printf("TestUnSafetyPubSub.Unsubscribe failed, err: %v\n", err)
			}
			wg.Done()
		}(i, &waitGroup)
		go func(idx int, wg *sync.WaitGroup) {
			coTopicKey := fmt.Sprintf("%s_%d", topicKey, idx)
			if err := obs.Unsubscribe(coTopicKey, idx); err != nil {
				log.Printf("TestUnSafetyPubSub.Unsubscribe failed, err: %v\n", err)
			}
			wg.Done()
		}(i, &waitGroup)
	}
	waitGroup.Wait()

	log.Println("TestSafetyPubSub succeeded")
}
