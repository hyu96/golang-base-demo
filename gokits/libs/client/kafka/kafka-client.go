package ckafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/huydq/gokits/libs/env"
	"github.com/huydq/gokits/libs/ilog"
)

type ClientKafka struct {
	config            *ConfigKafka
	producer          sarama.SyncProducer
	kafkaClientConfig *sarama.Config
}

var (
	clientKafka *ClientKafka
)

// default value env key is "Kafka";
// if configKeys was set, key env will be first value (not empty) of this;
func InstallKafkaClient(configKeys ...string) *ClientKafka {
	if clientKafka != nil {
		return clientKafka
	}

	getKafkaConfigFromEnv(configKeys...)
	if kafkaConfig == nil {
		panic(fmt.Errorf("need config for kafka client first"))
	}

	conf := sarama.NewConfig()
	conf.Consumer.Return.Errors = true
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Retry.Max = producerMaxRetry
	conf.Producer.Return.Successes = true
	conf.Producer.Compression = sarama.CompressionSnappy
	conf.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	conf.Producer.Flush.Frequency = producerFlushFrequency
	conf.ClientID = env.Config().PodName

	producer, err := sarama.NewSyncProducer(kafkaConfig.Addrs, conf)
	if err != nil {
		panic(err)
	}

	clientKafka = &ClientKafka{
		kafkaClientConfig: conf,
		config:            kafkaConfig,
		producer:          producer,
	}

	go clientKafka.CreateTopic()

	return clientKafka
}

func GetCKafka() *ClientKafka {
	if clientKafka == nil {
		return InstallKafkaClient()
	}

	return clientKafka
}

func checkExistTopic(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func (c *ClientKafka) CreateTopic() {
	if c == nil || c.config == nil {
		ilog.Error("ClientKafka::CreateTopic - Need InstallKafkaClient first")
		return
	}

	if len(c.config.ProducerTopics) == 0 {
		ilog.Error("ClientKafka::CreateTopic - Need least one topic")

		return
	}

	topicDetail := &sarama.TopicDetail{}
	topicDetail.NumPartitions = c.config.NumPartitions
	topicDetail.ReplicationFactor = c.config.ReplicationFactor
	topicDetail.ConfigEntries = make(map[string]*string)

	topicDetails := make(map[string]*sarama.TopicDetail)

	consumer, err := sarama.NewConsumer(c.config.Addrs, c.kafkaClientConfig)
	if err != nil {
		ilog.Errorf("ClientKafka::CreateTopic - Error when create new consumer %+v", err)

		return
	}

	listTopic, _ := consumer.Topics()
	// ilog.Debugf("get topics available in kafka: %s", ijson.ToJsonString(listTopic))

	listTopicNotExisted := make([]string, 0)
	for _, topicName := range c.config.ProducerTopics {
		if _, found := checkExistTopic(listTopic, topicName); !found {
			listTopicNotExisted = append(listTopicNotExisted, topicName)
		}
	}

	if len(listTopicNotExisted) == 0 {
		ilog.Infof("ClientKafka::CreateTopic - All of %+v was existed", c.config.ProducerTopics)

		return
	}

	for _, topicName := range listTopicNotExisted {
		topicDetails[topicName] = topicDetail
	}

	request := sarama.CreateTopicsRequest{
		Timeout:      createRequestTimeout,
		TopicDetails: topicDetails,
	}

	// Send request to Broker
	broker := sarama.NewBroker(c.config.Addrs[0])
	broker.Open(c.kafkaClientConfig)

	response, err := broker.CreateTopics(&request)
	if err != nil {
		ilog.Errorf("ClientKafka::CreateTopic - CreateTopics Error: %+v", err)

		return
	}

	t := response.TopicErrors
	for key, val := range t {
		if val.Err != sarama.ErrNoError {
			ilog.Errorf("😡😡😡 ClientKafka::CreateTopic - Create topic key: %s - Error: %s at pod %s in host %s", key, val.Err.Error(), env.Config().PodName, env.Config().HostName)
		}
	}
}
