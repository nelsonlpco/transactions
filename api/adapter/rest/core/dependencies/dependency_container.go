package dependencies

import (
	"github.com/nelsonlpco/transactions/infrastructure/db_manager"
)

type DependencyContainer struct {
	Datasources  *Datasources
	Repositories *Repositories
	UseCases     *UseCases
	Services     *Services
	Controllers  *Controllers
}

func NewDependencyContainer(dbManager *db_manager.DBManager) *DependencyContainer {
	datasources := NewSqlDatasources(dbManager)
	repositories := NewRepositories(datasources)
	usecases := NewUseCases(repositories)
	services := NewServices(usecases)
	controllers := NewControllers(services)

	dependencyContainer := &DependencyContainer{
		Datasources:  datasources,
		Repositories: repositories,
		UseCases:     usecases,
		Services:     services,
		Controllers:  controllers,
	}

	return dependencyContainer
}
