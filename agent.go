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
	"github.com/samjtro/go-dsr"
	"maragu.dev/llm"
)

type (
	AnthropicAgent struct {
		Client      *llm.AnthropicClient
		Instruction string
		Role        string
	}

	GoogleAgent struct {
		Client      *llm.GoogleClient
		Instruction string
		Role        string
	}

	OpenAIAgent struct {
		Client      *llm.OpenAIClient
		Instruction string
		Role        string
	}

	DeepseekAgent struct {
		Client      *dsr.ChatClient
		Instruction string
		Role        string
	}

	AgentOptions struct {
		Instruction string
		Role        string `default:"Assistant"`
	}

	AgentOption func(o *AgentOptions)
)

func Instruction(i string) AgentOption {
	return func(o *AgentOptions) {
		o.Instruction = i
	}
}

func Role(r string) AgentOption {
	return func(o *AgentOptions) {
		o.Role = r
	}
}
