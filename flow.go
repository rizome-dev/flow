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
	return &f
}

func (f *Flow) AddAnthropicAgent(name string, agent *AnthropicAgent) {
	f.AnthropicAgents[name] = agent
}

func (f *Flow) AddGoogleAgent(name string, agent *GoogleAgent) {
	f.GoogleAgents[name] = agent
}

func (f *Flow) AddOpenAIAgent(name string, agent *OpenAIAgent) {
	f.OpenAIAgents[name] = agent
}

func (f *Flow) AddDeepseekAgent(name string, agent *DeepseekAgent) {
	f.DeepseekAgents[name] = agent
}

func (f *Flow) AddTool(tool *Tool, function func(*Config) error) {
	f.Tools[tool] = function
}

func (f *Flow) AddResource(name string, resource *Resource) {
	f.Resources[name] = resource
}
