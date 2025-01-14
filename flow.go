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

import "github.com/google/generative-ai-go/genai"

// High-level workflow abstraction
type (
	Flow struct {
		AnthrophicAgents []*AnthropicAgent
		GoogleAgents     []*GoogleAgent
		OpenAIAgents     []*OpenAIAgent
		Extensions       map[*Function]func(*Config) error
		Resources        map[string]*Resource
	}

	FlowOption func(o *Flow)
)

func AnthropicAgents(agents []*AnthropicAgent) FlowOption {
	return func(o *Flow) {
		o.AnthrophicAgents = agents
	}
}

func GoogleAgents(agents []*GoogleAgent) FlowOption {
	return func(o *Flow) {
		o.GoogleAgents = agents
	}
}

func OpenAIAgents(agents []*OpenAIAgent) FlowOption {
	return func(o *Flow) {
		o.OpenAIAgents = agents
	}
}

func Extensions(extensions map[*Function]func(*Config) error) FlowOption {
	return func(o *Flow) {
		o.Extensions = extensions
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
	f.AnthrophicAgents = append(f.AnthrophicAgents, agent)
}

func (f *Flow) AddGoogleAgent(name string, agent *GoogleAgent) {
	f.GoogleAgents = append(f.GoogleAgents, agent)
}

func (f *Flow) AddOpenAIAgent(name string, agent *OpenAIAgent) {
	f.OpenAIAgents = append(f.OpenAIAgents, agent)
}

// Pass the function's name, description & parameters
// params should be formatted as map[parameterName]parameterType
func (f *Flow) AddExtension(name, desc string, params map[string]any, function func(*Config) error) {
	var p []genai.Schema
	for k, v := range params {
		var t genai.Type
		a := v.(string)
		if a[0:1] == "[]" && (a[2:5] != "rune" && a[2:5] != "byte") {
			t = genai.TypeArray
		} else {
			switch a {
			case "bool":
				t = genai.TypeBoolean
			case "string", "[]byte", "[]rune":
				t = genai.TypeString
			case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "byte", "rune":
				t = genai.TypeInteger
			case "float32", "float64", "complex64", "complex128":
				t = genai.TypeNumber
			default:
				t = genai.TypeObject
			}
		}
		p = append(p, genai.Schema{
			Type:        t,
			Description: k,
		})
	}
	f.Extensions[&Function{
		name,
		desc,
		p,
	}] = function
}

func (f *Flow) AddResource(name string, resource *Resource) {
	f.Resources[name] = resource
}
