package pubsub

import (
	"fmt"
	"log"
	"reflect"
	"runtime/debug"
	"sync"
)

type TopicFunc func(args ...interface{})

type SubArg struct {
	Topic     interface{}
	SubKey    interface{}
	TopicFunc TopicFunc
	PreArgs   []interface{}
}

type ObServer struct {
	enableSafeMode bool
	topicMap       map[interface{}]map[interface{}]*Subject
	mutex          sync.RWMutex
}

func (t *ObServer) init(enableSafeMode bool) {
	t.enableSafeMode = enableSafeMode
	t.topicMap = make(map[interface{}]map[interface{}]*Subject)
}

func (t *ObServer) Unsubscribe(topic interface{}, subKey interface{}) (retErr error) {
	defer func() {
		if err := recover(); err != nil {
			retErr = fmt.Errorf("unsubscribe err: %v", err)
		}
	}()

	if t.enableSafeMode {
		t.mutex.Lock()
		defer t.mutex.Unlock()
	}

	subjectMap, exist := t.topicMap[topic]
	if !exist {
		return
	}

	delete(subjectMap, subKey)
	if len(subjectMap) <= 0 {
		delete(t.topicMap, topic)
	}
	return
}

func (t *ObServer) Subscribe(topic interface{}, subKey interface{}, topicFunc TopicFunc, preArgs ...interface{}) (retErr error) {
	defer func() {
		if err := recover(); err != nil {
			retErr = fmt.Errorf("subscribe err: %v", err)
		}
	}()

	subject := &Subject{}
	subject.topicFunc = topicFunc
	subject.topicFuncValue = reflect.ValueOf(topicFunc)
	for _, val := range preArgs {
		subject.preTopicArgs = append(subject.preTopicArgs, reflect.ValueOf(val))
	}

	if t.enableSafeMode {
		t.mutex.Lock()
		defer t.mutex.Unlock()
	}

	if _, exist := t.topicMap[topic]; !exist {
		t.topicMap[topic] = make(map[interface{}]*Subject)
	}
	t.topicMap[topic][subKey] = subject
	return
}

func (t *ObServer) Publish(topic interface{}, enableCo bool, args ...interface{}) (retErr error) {
	defer func() {
		if err := recover(); err != nil {
			retErr = fmt.Errorf("publish err: %v", err)
		}
	}()

	if t.enableSafeMode {
		t.mutex.RLock()
	}

	subjectMap, exist := t.topicMap[topic]
	if !exist {
		if t.enableSafeMode {
			t.mutex.RUnlock()
		}
		return
	}

	var subjects []*Subject
	for _, subject := range subjectMap {
		subjects = append(subjects, subject)
	}

	if t.enableSafeMode {
		t.mutex.RUnlock()
	}

	for _, subject := range subjects {
		if !enableCo {
			callArgs := subject.getCallArgs(args...)
			subject.topicFuncValue.Call(callArgs)
			continue
		}

		go func(sub *Subject) {
			if r := recover(); r != nil {
				log.Printf("Call event func, event: %v, subKy: %v, recover: %v, stack: %v\n",
					r, string(debug.Stack()), sub.topic, sub.key)
			}

			callArgs := sub.getCallArgs(args...)
			sub.topicFuncValue.Call(callArgs)

		}(subject)

	}

	return
}

func NewObServer(enableSafeMode bool, subArgs ...SubArg) (*ObServer, error) {
	obs := &ObServer{}
	obs.init(enableSafeMode)
	for _, arg := range subArgs {
		if err := obs.Subscribe(arg.Topic, arg.SubKey, arg.TopicFunc, arg.PreArgs...); err != nil {
			return nil, err
		}
	}

	return obs, nil
}
