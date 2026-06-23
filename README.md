# mockapi

A self-hosted **mockmock server**.

## The problem

When building against a service that isn't ready yet — a third-party API, a teammate's
unfinished endpoint, a flaky staging environment — you need *something* that answers.
The usual options are awkward:

## Layout

```
cmd/api/         entry point — wires the dependency graph
internal/
├── config/      environment configuration
├── database/    PostgreSQL connection + migrations
├── server/      router assembly
├── web/         embedded web UI
└── mockmock/    the feature (domain, service, handler, repository)
pkg/
├── apperr/      reusable typed application errors
└── httputil/    reusable JSON response helpers + RequestID middleware
db/migrations/   SQL migrations
deployments/     Dockerfile + docker-compose
```# mockmock
# mockmockk
