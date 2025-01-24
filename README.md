# flow

<img src="/internal/logo.png" alt="flow - rizome labs" width="250" align="right">

[![GoDoc](https://pkg.go.dev/badge/github.com/rizome-dev/flow)](https://pkg.go.dev/github.com/rizome-dev/flow)
[![Go Report Card](https://goreportcard.com/badge/github.com/rizome-dev/flow)](https://goreportcard.com/report/github.com/rizome-dev/flow)

⚠️ **This is an internal library belonging to rizome labs, and is currently in development**. Things will probably break, but existing functionality is usable. ⚠️

```shell
go get github.com/rizome-dev/flow
```

built by: [rizome labs](https://rizome.dev)

contact us: [hi (at) rizome.dev](mailto:hi@rizome.dev)

## what flow is not

while we are incorporating some level of agentic behavior (WIP), the point of this library is to provide a framework for building workflows; NOT agents for building agents. there are many good options for building ReAct, MRKL, etc. agents - and this is not one of them.

that being said, if your definition of an "agent" is just an llm that has access to tools - then this is for you!

## what flow is

flow is a WORKFLOW framework for building multi-client workflows, that incorporate some level of agentic behavior (WIP). the primary goal is to create efficient, modular structures to manage workflows.

all `Agents` have access to the flow's `Tools` & `Resources`; hooks to datastores such as AlloyDB.

the OpenAI, Anthropic & Google `Agent` classes are just wrappers for a [maragu.dev/llm](https://maragu.dev/llm) client; which itself is simply an slog instance for logging & the native client.

the `DeepseekAgent` is a wrapper for a [samjtro/go-dsr](https://github.com/samjtro/go-dsr) client.

all tools take a `Config` struct, which is a wrapper for all of the common information you would need in a typical tool call, and return an error.

theoretically, nearly every tool will take in the previous message, do something with it, and append the results either to the messages, or to the appropriate `Resource`, which we then direct the llm to.

simple, right?

planned features:
- resources
- parallel tool calls

## examples
### agent context switch
