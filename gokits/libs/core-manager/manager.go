// core manager
// Quản lý các service core business
package coremanager

import (
	"errors"
	"sync"
)

type CoreManager struct {
	cores []ICore
}

// implement by cores
type ICore interface {
	RegisterCallback(cb interface{}) // callback to other service
}

// for test or auto run like kafka, rapid mq: this will install after all cores create done
type IPartsAfterCoresInstalled interface {
	RegisterJobs()
}

func New() *CoreManager {
	return &CoreManager{
		cores: make([]ICore, 0),
	}
}

func (c *CoreManager) GetCores() []ICore {
	return c.cores
}

func (c *CoreManager) AddCore(core ICore) {
	if core == nil {
		panic(errors.New("core is nil"))
	}
	for _, existingCore := range c.cores {
		if existingCore == core {
			return
		}
	}
	c.cores = append(c.cores, core)
}

func (c *CoreManager) InstallCoresDependencies() {
	var wg sync.WaitGroup
	for _, core1 := range c.cores {
		for _, core2 := range c.cores {
			wg.Add(1)
			go func(core1 ICore, core2 ICore) {
				defer wg.Done()
				core1.RegisterCallback(core2)
			}(core1, core2)
		}
	}
	wg.Wait()

	for _, core := range c.cores {
		wg.Add(1)
		go func(core ICore) {
			defer wg.Done()
			if jobs, ok := core.(IPartsAfterCoresInstalled); ok {
				jobs.RegisterJobs()
			}
		}(core)
	}
	wg.Wait()
}
