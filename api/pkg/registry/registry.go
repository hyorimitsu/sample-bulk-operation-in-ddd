package registry

import (
	"gorm.io/gorm"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/queryservice"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/usecase"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/domainservice"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/infra/db/repository"
)

type Registry struct {
	TaskQueryService  queryservice.TaskQueryService
	TaskDomainService domainservice.TaskDomainService
	TaskUseCaser      usecase.TaskUseCaser
}

func NewRegistry(db *gorm.DB) *Registry {
	reg := &Registry{}
	reg.TaskQueryService = repository.NewTaskQueryRepository(db)
	reg.TaskDomainService = repository.NewTaskCommandRepository(db)
	reg.TaskUseCaser = usecase.NewTaskUseCase(
		reg.TaskQueryService,
		reg.TaskDomainService,
	)
	return reg
}
