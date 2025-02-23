package elastic

import (
	"fmt"
	"log"

	elastic "github.com/elastic/go-elasticsearch/v7"
)

func ConnectElasticSearch() (*elastic.Client, error){
	fmt.Printf("Creating elasticsearch client...")
	LoadGrpcClientConfig()
	cfgE := elastic.Config{
		Addresses: []string{
			fmt.Sprintf("http://%v:%v@%v:%v", Configs.ElasUser, Configs.ElasPass, Configs.ElasHost, Configs.ElasPort),
		},
	}
	es, err := elastic.NewClient(cfgE)
	if err != nil {
		log.Printf("Error creating elastic client: %v", err)
	}
	_, err = es.Ping()
	if err!= nil {
		log.Printf("Error Pinging es: %v", err)
	}

	return es, err
}
