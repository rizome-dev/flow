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
	"context"
	"os"

	"maragu.dev/llm"
)

type (
	AnthropicAgent struct{ *llm.AnthropicClient }

	GoogleAgent struct{ *llm.GoogleClient }

	OpenAIAgent struct{ *llm.OpenAIClient }
)

func (f *Flow) NewAnthropicAgent(ctx context.Context, name string) {
	ANTHROPIC_API_KEY := os.Getenv("ANTHROPIC_API_KEY")
	client := llm.NewAnthropicClient(llm.NewAnthropicClientOptions{
		Key: ANTHROPIC_API_KEY,
	})
	f.AnthropicAgents[name] = &AnthropicAgent{client}
}

func (f *Flow) NewGoogleAgent(ctx context.Context, name string) {
	GOOGLE_API_KEY := os.Getenv("GOOGLE_API_KEY")
	client := llm.NewGoogleClient(llm.NewGoogleClientOptions{
		Key: GOOGLE_API_KEY,
	})
	f.GoogleAgents[name] = &GoogleAgent{client}
}

func (f *Flow) NewOpenAIAgent(ctx context.Context, name string) {
	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")
	client := llm.NewOpenAIClient(llm.NewOpenAIClientOptions{
		Key: OPENAI_API_KEY,
	})
	f.OpenAIAgents[name] = &OpenAIAgent{client}
}
