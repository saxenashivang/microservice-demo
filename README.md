# Microservice demo

Simple microservice using go-micr

## Getting Started

[Download][2] and install **Go**. Version `1.16` or higher is required.

Installation is done by using the [`go install`][3] command.

```bash
go install github.com/go-micro/cli/cmd/go-micro@latest
```

Let's create a new service using the `new` command.

```bash
go-micro new service helloworld
```

Follow the on-screen instructions. Next, we can run the program.

```bash
cd helloworld
make proto tidy
go-micro run
```

Finally we can call the service.

```bash
go-micro call helloworld Helloworld.Call '{"name": "John"}'
```

That's all you need to know to get started. Refer to the [Go Micro][1]
documentation for more info on developing services.

## Frmaeworks and library used

### Jaeger

### gRPC Server/Client

### Tern - Postgres Migrations

Tern can be used to create and manage Postgres migrations. Go-micro can set the
service up to use Tern SQL migrations.

### sqlc - SQL Code Generation

[Sqlc](14) can compile SQL queries into boilerplate Go code that allows you to easily
create and manage your database layer. Go-micro can set your service up for use
with sqlc, and used Postgres as a default backend. Sqlc works well in combination
with [Tern](#tern---postgres-migrations).

### Docker BuildKit

[Docker BuildKit](11) is a new container build engine that provides new useful
features, such as the ability to cache specific directories across builds. This
can prevent Go from having to re-download modules every build.

### Kubernetes

### Kustomize - Kubernetes Resource Management

[Kustomize](15) can be used to manage more complex Kubernetes manifests for various
deployments, such as a development and production environment.

### gRPC Health Protocol - Kubernetes Probes

Since Kubernetes [1.24](12), probes can make use of the [gRPC Health Protocol](13).
This allows you to directly probe the go-micro service in a Kubernetes container
if it implements the health protocol.

### Tilt - Kubernetes Deployment

[Tilt][16] can be used to set up a local Kubernetes deployment pipeline.

### Skaffold - Kubernetes Deployment

Skaffold can be used to locally deploy a service.
