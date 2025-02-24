package ckafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/huydq/gokits/libs/ilog"
)

func (c *ClientKafka) ProducerPushMessage(topic string, message string) (partition int32, offset int64, err error) {
	if c == nil || c.producer == nil {
		return 0, 0, fmt.Errorf("ClientKafka::ProducerPushMessage - Not found any producer")
	}

	ilog.Infof("ClientKafka::ProducerPushMessage - Push to topic: %s object data: %s", topic, message)

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	return c.producer.SendMessage(msg)
}

func (c *ClientKafka) ProducerPushMessageWithKey(topic, key string, message string) (partition int32, offset int64, err error) {
	if c == nil || c.producer == nil {
		return 0, 0, fmt.Errorf("ClientKafka::ProducerPushMessage - Not found any producer")
	}

	ilog.Infof("ClientKafka::ProducerPushMessage - Push to topic: %s, key: %s, object data: %s", topic, key, message)

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(message),
	}

	return c.producer.SendMessage(msg)
}
