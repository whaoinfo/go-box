package pubsub

import "reflect"

type Subject struct {
	topic          interface{}
	key            interface{}
	topicFunc      TopicFunc
	topicFuncValue reflect.Value
	preTopicArgs   []reflect.Value
}

func (t *Subject) getCallArgs(pubArgs ...interface{}) []reflect.Value {
	if len(pubArgs) <= 0 {
		return t.preTopicArgs
	}

	var retArgs []reflect.Value
	retArgs = append(retArgs, t.preTopicArgs...)
	for _, val := range pubArgs {
		retArgs = append(retArgs, reflect.ValueOf(val))
	}
	return retArgs
}
