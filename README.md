# flow

<img src="/logo.png" alt="flow - rizome labs" width="250" align="right">

[![GoDoc](https://pkg.go.dev/badge/github.com/rizome-dev/flow)](https://pkg.go.dev/github.com/rizome-dev/flow)
[![Go Report Card](https://goreportcard.com/badge/github.com/rizome-dev/flow)](https://goreportcard.com/report/github.com/rizome-dev/flow)

⚠️ **This library is in development**. Things will probably break, but existing functionality is usable. ⚠️

```shell
go get github.com/rizome-dev/flow
```

built by: [rizome labs](https://rizome.dev)

contact us: [hi (at) rizome.dev](mailto:hi@rizome.dev)

## what is flow?

pre-existing agent workflow libraries are great, but they have one major conceptual problem; agents can take actions, groups of agents can take actions, and the abstractions for both tend to get muddled.

a larger, more glaring flaw is that tools belong to agents - this creates a situation where tools have to be handed off, which impacts your ability to write high-performant multi-agent workflows

that's the entire point of a workflow library! we propose a higher-level of abstraction; the `Flow`.

all `Agents` have access to the flow's `Extensions`, a fancy word for tools, & `Resources`, hooks to datastores such as AlloyDB.

the various `Agent` classes have names & clients; the clients are just wrappers for the llm provider vis-a-vis [maragu.dev/llm](https://maragu.dev/llm).

extension functions take a `Config` struct, which is a wrapper for all of the common information you would need in a typical tool call, and return an error.

theoretically, nearly every extension function will take in the previous message, do something with it, and append the results either to the messages, or to the appropriate `Resource`, which we then direct the llm to.

simple, right?

planned features:
- resources
- parallel tool calls

## examples
### agent context switch
