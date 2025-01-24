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
	"github.com/openai/openai-go"
)

type (
	Tool struct {
		Name        string
		Description string
		Parameters  []*Parameter
	}

	Parameter struct {
		Name        string
		Description string
		Type        string
	}

	// Used as parameter for tool calls; covers ~95% of use cases
	// If it doesn't; let us know:
	// github.com/rizome-dev/flow/issues
	Config struct {
		LastMessage *Message
		Messages    []*Message
	}

	Message struct {
		Name    string
		Role    string
		Content string
	}
)

func (c *Config) AddLastMessage(msg *Message) error {
	c.LastMessage = msg
	c.Messages = append(c.Messages, msg)
	return nil
}

func MarshalTool(func(*Config) error) *Tool {
	return &Tool{}
}

// Pass the function's name, description & parameters
// params should be formatted as map[parameterName]parameterType
func CreateTool(name, desc string, params []*Parameter) *Tool {
	return &Tool{
		name,
		desc,
		params,
	}
}

func (t *Tool) MarshalOpenAIChatCompletionTool() *openai.ChatCompletionToolParam {
	props := map[string]interface{}{}
	for _, x := range t.Parameters {
		props[x.Name] = map[string]string{
			"type": x.Type,
		}
	}
	params := openai.FunctionParameters{
		"type":       "object",
		"properties": props,
		"required":   []string{},
	}
	return &openai.ChatCompletionToolParam{
		Type: openai.F(openai.ChatCompletionToolTypeFunction),
		Function: openai.F(openai.FunctionDefinitionParam{
			Name:        openai.String(t.Name),
			Description: openai.String(t.Description),
			Parameters:  openai.F(params),
		}),
	}
}

func (t *Tool) MarshalGeminiChatCompletionTool() *genai.Tool {
	var arr []string
	for _, x := range t.Parameters {
		arr = append(arr, x.Name)
	}
	p := &genai.Schema{
		Type:       genai.TypeObject,
		Properties: map[string]*genai.Schema{},
		Required:   arr,
	}
	for _, x := range t.Parameters {
		var t genai.Type
		switch x.Type {
		case "string":
			t = genai.TypeString
		case "number":
			t = genai.TypeNumber
		case "integer":
			t = genai.TypeInteger
		case "boolean":
			t = genai.TypeBoolean
		case "array":
			t = genai.TypeArray
		case "object":
			t = genai.TypeObject
		}
		p.Properties[x.Name] = &genai.Schema{
			Type:        t,
			Description: p.Description,
		}
	}
	return &genai.Tool{
		FunctionDeclarations: []*genai.FunctionDeclaration{{
			t.Name,
			t.Description,
			p,
		}},
	}
}
