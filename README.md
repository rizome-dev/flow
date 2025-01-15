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

flow aims to solve two pain points of existing agent frameworks:
1. both agents & workflows have instructions, and with the use of tools, it tends to get messy
2. tools belong to agents, which means they have to get passed around

we propose a higher-level of abstraction; the `Flow`. all `Agents` have access to the flow's `Tools` & `Resources`; hooks to datastores such as AlloyDB.

the various `Agent` classes are just wrappers for a [maragu.dev/llm](https://maragu.dev/llm) client; which itself is simply an slog instance for logging & a native client.

all tools take a `Config` struct, which is a wrapper for all of the common information you would need in a typical tool call, and return an error.

theoretically, nearly every tool will take in the previous message, do something with it, and append the results either to the messages, or to the appropriate `Resource`, which we then direct the llm to.

simple, right?

planned features:
- resources
- parallel tool calls

## examples
### agent context switch
