package ekafka

import (
	"fmt"

	"github.com/gotomicro/ego/core/elog"
)

const PackageName = "component.ekafka"

// Component kafka 组件，包含Client、Producers、Consumers
type Component struct {
	Config    *Config
	logger    *elog.Component
	client    *Client
	consumers map[string]*Consumer
	producers map[string]*Producer
}

// Consumer 返回指定名称的kafka Consumer
func (cmp Component) Consumer(name string) *Consumer {
	fmt.Printf("cmp.consumers--------------->"+"%+v\n", cmp.consumers)
	return cmp.consumers[name]
}

// Consumer 返回指定名称的kafka Producer
func (cmp Component) Producer(name string) *Producer {
	fmt.Printf("cmp.producers--------------->"+"%+v\n", cmp.producers)
	return cmp.producers[name]
}

// Client 返回kafka Client
func (cmp Component) Client() *Client {
	return cmp.client
}
