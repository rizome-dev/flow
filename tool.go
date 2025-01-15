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
	"github.com/google/generative-ai-go/genai"
)

type (
	Function struct {
		Name        string
		Description string
		Parameters  []genai.Schema
	}

	// Used as parameter for tool calls; covers ~95% of use cases
	// If it doesn't; let us know:
	// github.com/rizome-dev/flow/issues
	Config struct {
		Name        string
		Messages    []*Message
		LastMessage *Message
		Role        string
	}

	Message struct {
		Role    string `json:"omitempty"`
		Content string
	}
)

func (c *Config) AddLastMessage(msg *Message) error {
	c.LastMessage = msg
	c.Messages = append(c.Messages, msg)
	return nil
}

func (f *Flow) CallTool(conf *Config) error {
	for x, y := range f.Tools {
		if x.Name == conf.Name {
			return y(conf)
		}
	}
	return conf.AddLastMessage(&Message{
		conf.Role,
		"tool not found.",
	})
}
