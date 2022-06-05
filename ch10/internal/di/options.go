package di

type DependencyOption interface {
	configureDependency(*dependencyInfo)
}

func (f DependencyCleanupFunc) configureDependency(i *dependencyInfo) {
	i.cleanup = f
}
