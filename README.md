Sample Bulk Operation in DDD
---

This is sample of bulk operation in DDD.


## Description

This is a sample implementation for bulk operation in DDD using the Specification patterns and Command patterns.

NOTE: This is just a sample, and details such as error handling, logging, environment variables usage, and other fine-grained processing have been simplified or omitted.


## Structure

### Used language, tools and other components

| language / tools                     | description                                               |
|--------------------------------------|-----------------------------------------------------------|
| [Go](https://github.com/golang/go)   | programming language                                      |
| [Kubernetes](https://kubernetes.io/) | container orchestrator                                    |
| [Skaffold](https://skaffold.dev/)    | tool for building, pushing and deploying your application |

### Used architectures

- [DDD](https://en.wikipedia.org/wiki/Domain-driven_design)
- [CQRS Pattern](https://en.wikipedia.org/wiki/Command%E2%80%93query_separation#Command_Query_Responsibility_Segregation)
- [Specification Pattern](https://en.wikipedia.org/wiki/Specification_pattern)
- [Command Pattern](https://en.wikipedia.org/wiki/Command_pattern)

### Directories

```
.
├── .k8s                          # => Kubernetes manifests
├── api                           # => API implementation
│    ├── cmd                      # => Entrypoint
│    ├── config                   # => Configurations
│    ├── pkg
│    │   ├── application          # => Application layer (Use Case layer)
│    │   ├── domain               # => Domain layer
│    │   │    ├── cmd             # => Domain Specification
│    │   │    ├── spec            # => Domain Command
│    │   │    └── (some omitted)
│    │   ├── infra                # => Infrastructure layer
│    │   ├── interface            # => Presentation layer
│    │   └── (some omitted)
│    └── (some omitted)
├── database                      # => Database migrations and seeds
├── skaffold.yaml
└── (some omitted)
```


## Architecture

In this sample, the Specification Pattern and Command Pattern are used to realize bulk operation in DDD.  
The advantages of using this architecture are explained step by step.

### Pattern of Sequential processing

As the most simple and easy-to-understand approach, there is sequential processing using a for-loop, as shown below.

```go
tasks, err := u.taskDoaminService.List(ctx, params)
...

for _, task := range tasks {
    err := task.UpdateStatus(value.TaskStatusDoneString)
    ...
}
```

This may be acceptable if the amount of data is small.
However, as the data grows during operation, it is easy to imagine that performance issues may arise, such as memory and processing speed.

### Pattern of expressing everything in SQL

Next, a pattern can be considered where everything is expressed in the SQL of the repository layer, as shown below.

```go
r.db.
    WithContext(ctx).
    Model(&model.Task{}).
    Where("due_date < ? AND status != ?", params...).
    Updates(values).
    Error
```

In this case, the performance issues mentioned in the "[Pattern of Sequential processing](https://github.com/hyorimitsu/sample-bulk-operation-in-ddd#pattern-of-sequential-processing)" are resolved.
However, it is necessary to define the application's specifications (domain knowledge) in both the domain layer and the repository layer.

### Patterns using Specification and Command

The pattern that solves the problems mentioned so far is using Specifications and Commands.

Specifications (defining which data to update) and Commands (defining how to update) are defined in the domain layer, and the repository layer uses these definitions to construct SQL queries.

In this approach, there is no need to maintain the actual entity of the data being updated, so the performance issues mentioned in the "[Pattern of Sequential processing](https://github.com/hyorimitsu/sample-bulk-operation-in-ddd#pattern-of-sequential-processing)" are resolved.  
Furthermore, since the application's specifications (domain knowledge) are aggregated as Specifications and Commands in the domain layer, the issues mentioned in the "[Pattern of expressing everything in SQL](https://github.com/hyorimitsu/sample-bulk-operation-in-ddd#pattern-of-expressing-everything-in-sql)" are also resolved.

For detailed implementation methods, please refer to the following.

- [Specification](https://github.com/hyorimitsu/sample-bulk-operation-in-ddd/blob/main/api/pkg/domain/spec/task.go)
- [Command](https://github.com/hyorimitsu/sample-bulk-operation-in-ddd/blob/main/api/pkg/domain/cmd/task.go)


## Usage

1. Run the application in minikube

      ```shell
      make run
      ```

2. Migrate database

      ```shell
      make migrate
      ```

3. Call API 

      Let's call api to see how it works.
      
      - [Routes definitions](https://github.com/hyorimitsu/sample-bulk-operation-in-ddd/blob/main/api/pkg/interface/router/router.go#L19-L28)
      - [Parameter definitions](https://github.com/hyorimitsu/sample-bulk-operation-in-ddd/tree/main/api/pkg/application/input)
      
      For example, a list of tasks can be obtained by calling the API as follows.
      
      ```
      curl http://sample-bulk-operation-in-ddd.localhost.com/api/tasks
      ```

4. Stop the application in minikube

      ```shell
      make stop
      ```


## Troubleshoot

- Q1: Cannot start with the following error output.

  ```shell
  > [sample-bulk-operation-in-ddd-api-59f647fd7-ssg8j sample-bulk-operation-in-ddd-api] 2023/04/29 23:11:36 /go/src/github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/infra/db/db.go:15
  > [sample-bulk-operation-in-ddd-api-59f647fd7-ssg8j sample-bulk-operation-in-ddd-api] [error] failed to initialize database, got error dial tcp 10.97.125.188:3306: connect: connection refused
  > [sample-bulk-operation-in-ddd-api-59f647fd7-ssg8j sample-bulk-operation-in-ddd-api] [Error] unable to new db: dial tcp 10.97.125.188:3306: connect: connection refused
  ```

  A1: This problem is due to the timing of api and database  deployment creation. Try `make run` again without `make stop`.


## References

- DDD and bulk operations https://enterprisecraftsmanship.com/posts/ddd-bulk-operations/
