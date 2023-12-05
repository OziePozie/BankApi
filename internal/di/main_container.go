package di

type Container struct {
	repos    *RepoContainer
	routes   *RouterContainer
	services *ServiceContainer
}

func (c *Container) Repos() *RepoContainer {
	return c.repos
}

func (c *Container) SetRepos(repos *RepoContainer) {
	c.repos = repos
}

func (c *Container) Routes() *RouterContainer {
	return c.routes
}

func (c *Container) SetRoutes(routes *RouterContainer) {
	c.routes = routes
}

func (c *Container) Services() *ServiceContainer {
	return c.services
}

func (c *Container) SetServices(services *ServiceContainer) {
	c.services = services
}

func NewContainer() *Container {

	repos := NewRepoContainer()
	services := NewServiceContainer(repos)
	routes := NewRouterContainer(services)
	return &Container{
		services: services,
		routes:   routes,
		repos:    repos,
	}
}
