/*
Copyright 2025 rizome labs llc, hi@rizome.dev

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package flow

import (
	"os"

	"github.com/samjtro/go-dsr"
	"maragu.dev/llm"
)

type (
	Flow struct {
		AnthropicAgents map[string]*AnthropicAgent
		GoogleAgents    map[string]*GoogleAgent
		OpenAIAgents    map[string]*OpenAIAgent
		DeepseekAgents  map[string]*DeepseekAgent
		Tools           map[*Tool]func(*Config) error
		Resources       map[string]*Resource
	}

	FlowOption func(o *Flow)
)

func AnthropicAgents(agents map[string]*AnthropicAgent) FlowOption {
	return func(o *Flow) {
		o.AnthropicAgents = agents
	}
}

func GoogleAgents(agents map[string]*GoogleAgent) FlowOption {
	return func(o *Flow) {
		o.GoogleAgents = agents
	}
}

func OpenAIAgents(agents map[string]*OpenAIAgent) FlowOption {
	return func(o *Flow) {
		o.OpenAIAgents = agents
	}
}

func DeepseekAgents(agents map[string]*DeepseekAgent) FlowOption {
	return func(o *Flow) {
		o.DeepseekAgents = agents
	}
}

func Tools(tools map[*Tool]func(*Config) error) FlowOption {
	return func(o *Flow) {
		o.Tools = tools
	}
}

func Resources(resources map[string]*Resource) FlowOption {
	return func(o *Flow) {
		o.Resources = resources
	}
}

// Create a new Flow
func NewFlow(opts ...FlowOption) *Flow {
	f := Flow{}
	for _, o := range opts {
		o(&f)
	}
	if f.AnthropicAgents == nil {
		f.AnthropicAgents = make(map[string]*AnthropicAgent)
	}
	if f.GoogleAgents == nil {
		f.GoogleAgents = make(map[string]*GoogleAgent)
	}
	if f.OpenAIAgents == nil {
		f.OpenAIAgents = make(map[string]*OpenAIAgent)
	}
	if f.DeepseekAgents == nil {
		f.DeepseekAgents = make(map[string]*DeepseekAgent)
	}
	return &f
}

func (f *Flow) AddAnthropicAgent(name string, opts ...AgentOption) {
	var opt AgentOptions
	for _, o := range opts {
		o(&opt)
	}
	ANTHROPIC_API_KEY := os.Getenv("ANTHROPIC_API_KEY")
	client := llm.NewAnthropicClient(llm.NewAnthropicClientOptions{
		Key: ANTHROPIC_API_KEY,
	})
	f.AnthropicAgents[name] = &AnthropicAgent{
		client,
		opt.Instruction,
		opt.Role,
	}
}

func (f *Flow) AddDeepseekAgent(name string, opts ...AgentOption) {
	var opt AgentOptions
	for _, o := range opts {
		o(&opt)
	}
	client := dsr.NewChatClient()
	f.DeepseekAgents[name] = &DeepseekAgent{
		client,
		opt.Instruction,
		opt.Role,
	}
}

func (f *Flow) AddGoogleAgent(name string, opts ...AgentOption) {
	var opt AgentOptions
	for _, o := range opts {
		o(&opt)
	}
	client := llm.NewGoogleClient(llm.NewGoogleClientOptions{
		Key: os.Getenv("GOOGLE_API_KEY"),
	})
	f.GoogleAgents[name] = &GoogleAgent{
		client,
		opt.Instruction,
		opt.Role,
	}
}

func (f *Flow) AddOpenAIAgent(name string, opts ...AgentOption) {
	var opt AgentOptions
	for _, o := range opts {
		o(&opt)
	}
	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")
	client := llm.NewOpenAIClient(llm.NewOpenAIClientOptions{
		Key: OPENAI_API_KEY,
	})
	f.OpenAIAgents[name] = &OpenAIAgent{
		client,
		opt.Instruction,
		opt.Role,
	}
}

func (f *Flow) AddTool(tool *Tool, function func(*Config) error) {
	f.Tools[tool] = function
}

func (f *Flow) AddResource(name string, resource *Resource) {
	f.Resources[name] = resource
}
