package main

type AppConfig struct {
	Capacity int
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Capacity: 10,
	}
}

type RouterConfig struct {
	CoffeeController *CoffeeController
}

func NewRouterConfig(config *AppConfig) *RouterConfig {
	coffeeController := &CoffeeController{
		Table: NewTable(config.Capacity),
	}

	return &RouterConfig{
		CoffeeController: coffeeController,
	}
}
