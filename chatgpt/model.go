package chatgpt

import "encoding/json"

type CompletionRequest struct {
	// prompt:
	//   description: |-
	//     The prompt(s) to generate completions for, encoded as a string or array of strings.
	//     Note that <|endoftext|> is the document separator that the model sees during training, so if a prompt is not specified the model will generate as if from the beginning of a new document. Maximum allowed size of string list is 2048.
	//   oneOf:
	//     - type: string
	//       default: ''
	//       example: This is a test.
	//       nullable: true
	//     - type: array
	//       items:
	//         type: string
	//         default: ''
	//         example: This is a test.
	//         nullable: false
	//       description: Array size minimum of 1 and maximum of 2048
	Prompts []string `json:"prompt"`

	// max_tokens:
	//   description: The token count of your prompt plus max_tokens cannot exceed the model's context length. Most models have a context length of 2048 tokens (except for the newest models, which support 4096). Has minimum of 0.
	//   type: integer
	//   default: 16
	//   example: 16
	//   nullable: true
	MaxTokens int `json:"max_tokens,omitempty"`

	// temperature:
	//   description: |-
	//     What sampling temperature to use. Higher values means the model will take more risks. Try 0.9 for more creative applications, and 0 (argmax sampling) for ones with a well-defined answer.
	//     We generally recommend altering this or top_p but not both.
	//   type: number
	//   default: 1
	//   example: 1
	//   nullable: true
	Temperature float64 `json:"temperature,omitempty"`

	// top_p:
	//   description: |-
	//     An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	//     We generally recommend altering this or temperature but not both.
	//   type: number
	//   default: 1
	//   example: 1
	//   nullable: true
	TopP float64 `json:"top_p,omitempty"`

	// logit_bias:
	//   description: Defaults to null. Modify the likelihood of specified tokens appearing in the completion. Accepts a json object that maps tokens (specified by their token ID in the GPT tokenizer) to an associated bias value from -100 to 100. You can use this tokenizer tool (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token. As an example, you can pass {"50256" &#58; -100} to prevent the <|endoftext|> token from being generated.
	//   type: object
	//   nullable: false
	LogitBias map[string]int `json:"logit_bias,omitempty"`

	// user:
	//   description: A unique identifier representing your end-user, which can help monitoring and detecting abuse
	//   type: string
	//   nullable: false
	User string `json:"user,omitempty"`

	// 'n':
	//   description: |-
	//     How many completions to generate for each prompt. Minimum of 1 and maximum of 128 allowed.
	//     Note: Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for max_tokens and stop.
	//   type: integer
	//   default: 1
	//   example: 1
	//   nullable: true
	N int `json:"n,omitempty"`

	// stream:
	//   description: 'Whether to stream back partial progress. If set, tokens will be sent as data-only server-sent events as they become available, with the stream terminated by a data: [DONE] message.'
	//   type: boolean
	//   nullable: true
	//   default: false
	Stream bool `json:"stream,omitempty"`

	// logprobs:
	//   description: |-
	//     Include the log probabilities on the logprobs most likely tokens, as well the chosen tokens. For example, if logprobs is 5, the API will return a list of the 5 most likely tokens. The API will always return the logprob of the sampled token, so there may be up to logprobs+1 elements in the response.
	//     Minimum of 0 and maximum of 5 allowed.
	//   type: integer
	//   default: null
	//   nullable: true
	Logprobs int `json:"logprobs,omitempty"`

	// model:
	//   type: string
	//   example: davinci
	//   nullable: true
	//   description: ID of the model to use. You can use the Models_List operation to see all of your available models, or see our Models_Get overview for descriptions of them.
	Model string `json:"model,omitempty"`

	// suffix:
	//   type: string
	//   nullable: true
	//   description: The suffix that comes after a completion of inserted text.
	Suffix string `json:"suffix,omitempty"`

	// echo:
	//   description: Echo back the prompt in addition to the completion
	//   type: boolean
	//   default: false
	//   nullable: true
	Echo bool `json:"echo,omitempty"`

	// stop:
	//   description: Up to 4 sequences where the API will stop generating further tokens. The returned text will not contain the stop sequence.
	//   oneOf:
	//     - type: string
	//       default: <|endoftext|>
	//       example: |+
	//
	//       nullable: true
	//     - type: array
	//       items:
	//         type: string
	//         example:
	//           - |+
	//
	//         nullable: false
	//       description: Array minimum size of 1 and maximum of 4
	Stop []string `json:"stop,omitempty"`

	// completion_config:
	//   type: string
	//   nullable: true
	CompletionConfig string `json:"completion_config,omitempty"`

	// cache_level:
	//   description: can be used to disable any server-side caching, 0=no cache, 1=prompt prefix enabled, 2=full cache
	//   type: integer
	//   nullable: true
	CacheLevel int `json:"cache_level,omitempty"`

	// presence_penalty:
	//   description: Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
	//   type: number
	//   default: 0
	PresencePenalty float64 `json:"presence_penalty,omitempty"`

	// frequency_penalty:
	//   description: Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	//   type: number
	//   default: 0
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// best_of:
	//   description: |-
	//     Generates best_of completions server-side and returns the "best" (the one with the highest log probability per token). Results cannot be streamed.
	//     When used with n, best_of controls the number of candidate completions and n specifies how many to return – best_of must be greater than n.
	//     Note: Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for max_tokens and stop. Has maximum value of 128.
	//   type: integer
	BestOf int `json:"best_of,omitempty"`
}

type CompletionResponse struct {
	//  id:
	//    type: string
	ID string `json:"id,omitempty"`

	//  object:
	//    type: string
	Object string `json:"object,omitempty"`

	//  created:
	//    type: integer
	Created int `json:"created,omitempty"`

	//  model:
	//    type: string
	Model string `json:"model,omitempty"`

	//  choices:
	//    type: array
	//    items:
	//      type: CompletionChoice
	Choices []CompletionChoice `json:"choices,omitempty"`
}

type CompletionChoice struct {
	// text:
	//   type: string
	Text string `json:"text,omitempty"`

	// index:
	//   type: integer
	Index int `json:"index,omitempty"`

	// logprobs:
	//   type: Logprobes
	Logprobs Logprobs `json:"logprobs,omitempty"`

	// finish_reason:
	//   type: string
	FinishReason string `json:"finish_reason,omitempty"`
}

type Logprobs struct {
	// tokens:
	//   type: array
	//   items:
	//     type: string
	Tokens []string `json:"tokens,omitempty"`

	// token_logprobs:
	//   type: array
	//   items:
	//     type: number
	TokenLogprobs []float64 `json:"token_logprobs,omitempty"`

	// top_logprobs:
	//   type: array
	//   items:
	//     type: object
	//     additionalProperties:
	//       type: number
	TopLogprobs []map[string]float64 `json:"top_logprobs,omitempty"`

	// text_offset:
	//   type: array
	//   items:
	//     type: integer
	TextOffset []int `json:"text_offset,omitempty"`
}

type EmbeddingRequest struct {
	// input:
	//   description: |-
	//    Inputs text to get embeddings for, encoded as a string. To get embeddings for multiple inputs in a single
	//    request, pass an array of strings. Each input must not exceed 2048 tokens in length.
	//    Unless you are embedding code, we suggest replacing newlines (\n) in your input with a single space, as we
	//    have observed inferior results when newlines are present.
	//   oneOf:
	//     - type: string
	//       default: ''
	//       example: This is a test.
	//       nullable: true
	//     - type: array
	//       minItems: 1
	//       maxItems: 2048
	//       items:
	//         type: string
	//         minLength: 1
	//         example: This is a test.
	//         nullable: false
	Inputs []string `json:"input,omitempty"`

	// user:
	//   description: A unique identifier representing your end-user, which can help monitoring and detecting abuse.
	//   type: string
	//   nullable: false
	User string `json:"user,omitempty"`

	// input_type:
	//   description: input type of embedding search to use
	//   type: string
	//   example: query
	InputType string `json:"input_type,omitempty"`

	// model:
	//   type: string
	//   description:
	//  	ID of the model to use. You can use the Models_List operation to see all of your available models, or see
	// 		our Models_Get overview for descriptions of them.
	//   nullable: false
	Model string `json:"model,omitempty"`

	AdditionalProp1 map[string]interface{} `json:"additionalProp1,omitempty"`
}

type EmbeddingResponse struct {
	// object:
	//   type: string
	Object string `json:"object,omitempty"`

	// model:
	//   type: string
	Model string `json:"model,omitempty"`

	// data:
	//   type: []Data
	Data []EmbeddingData `json:"data,omitempty"`

	// usage:
	//   type: Usage
	Usage Usage `json:"usage,omitempty"`
}

type EmbeddingData struct {
	// index:
	//   type: integer
	Index int `json:"index,omitempty"`

	// object:
	//   type: string
	Object string `json:"object,omitempty"`

	// embedding:
	//   type: array
	//   items:
	//     type: number
	Embedding []float64 `json:"embedding,omitempty"`
}

type Usage struct {
	// prompt_tokens:
	//   type: integer
	PromptTokens int `json:"prompt_tokens,omitempty"`

	// completion_tokens:
	//   type: integer
	CompletionTokens int `json:"completion_tokens,omitempty"`

	// total_tokens:
	//   type: integer
	TotalTokens int `json:"total_tokens,omitempty"`
}

type ChatRequest struct {
	// messages:
	//   description: The messages to generate chat completions for, in the chat format.
	//   type: []ChatMessage
	//   minItems: 1
	//   items:
	//     type: object
	Messages []ChatMessage `json:"messages,omitempty"`

	// temperature:
	//   description: |-
	//     What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random,
	//    while lower values like 0.2 will make it more focused and deterministic.
	//     We generally recommend altering this or `top_p` but not both.
	//   type: number
	//   minimum: 0
	//   maximum: 2
	//   default: 1
	//   example: 1
	//   nullable: true
	Temperature float64 `json:"temperature,omitempty"`

	// top_p:
	//   description: |-
	//     An alternative to sampling with temperature, called nucleus sampling, where the model considers the results
	//    of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability
	//    mass are considered.
	//     We generally recommend altering this or `temperature` but not both.
	//   type: number
	//   minimum: 0
	//   maximum: 1
	//   default: 1
	//   example: 1
	//   nullable: true
	TopP float64 `json:"top_p,omitempty"`

	// 'n':
	//   description: How many chat completion choices to generate for each input message.
	//   type: integer
	//   minimum: 1
	//   maximum: 128
	//   default: 1
	//   example: 1
	//   nullable: true
	N int `json:"n,omitempty"`

	// stream:
	//   description:
	//  	'If set, partial message deltas will be sent, like in ChatGPT. Tokens will be sent as data-only server-sent
	// 		events as they become available, with the stream terminated by a `data: [DONE]` message.'
	//   type: boolean
	//   nullable: true
	//   default: false
	Stream bool `json:"stream,omitempty"`

	// stop:
	//   description: Up to 4 sequences where the API will stop generating further tokens.
	//   oneOf:
	//     - type: string
	//       nullable: true
	//     - type: array
	//       items:
	//         type: string
	//         nullable: false
	//       minItems: 1
	//       maxItems: 4
	//       description: Array minimum size of 1 and maximum of 4
	//   default: null
	Stop []string `json:"stop,omitempty"`

	// max_tokens:
	//   description:
	//  	The maximum number of tokens allowed for the generated answer. By default, the number of tokens the model
	// 		can return will be (4096 - prompt tokens).
	//   type: integer
	//   default: inf
	MaxTokens int `json:"max_tokens,omitempty"`

	// presence_penalty:
	//   description:
	//  	Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text
	// 		so far, increasing the model's likelihood to talk about new topics.
	//   type: number
	//   default: 0
	//   minimum: -2
	//   maximum: 2
	PresencePenalty float64 `json:"presence_penalty,omitempty"`

	// frequency_penalty:
	//   description: Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	//   type: number
	//   default: 0
	//   minimum: -2
	//   maximum: 2
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// logit_bias:
	//   description:
	//  	Modify the likelihood of specified tokens appearing in the completion. Accepts a json object that maps
	// 		tokens (specified by their token ID in the tokenizer) to an associated bias value from -100 to 100.
	//		Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect
	//		will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values
	//		like -100 or 100 should result in a ban or exclusive selection of the relevant token.
	//		like -100 or 100 should result in a ban or exclusive selection of the relevant token.
	//   type: object
	//   nullable: true
	LogitBias map[string]float64 `json:"logit_bias,omitempty"`

	// user:
	//   description: A unique identifier representing your end-user, which can help Azure OpenAI to monitor and detect abuse.
	//   type: string
	//   example: user-1234
	//   nullable: false
	User string `json:"user,omitempty"`
}

type ChatResponse struct {
	// id:
	//   type: string
	ID string `json:"id,omitempty"`

	// object:
	//   type: string
	Object string `json:"object,omitempty"`

	// created:
	//   type: integer
	//   format: unixtime
	Created int `json:"created,omitempty"`

	// model:
	//   type: string
	Model string `json:"model,omitempty"`

	// choices:
	//   type: []ChatChoice
	Choices []ChatChoice `json:"choices,omitempty"`

	// usage:
	//   type: Usage
	Usage Usage `json:"usage,omitempty"`
}

// ChatChoice
// Example of a Chunk response:
//
//	{
//	 "id": "chatcmpl-6zzeMna3lgVJNguBvgncWuAcVL8f2",
//	 "object": "chat.completion.chunk",
//	 "created": 1680233010,
//	 "model": "gpt-35-turbo",
//	 "choices": [
//	   {
//	     "index": 0,
//	     "finish_reason": null,
//	     "delta": {
//	       "content": "the"
//	     }
//	   }
//	 ],
//	 "usage": null
//	}
type ChatChoice struct {
	// index:
	//   type: integer
	Index int `json:"index,omitempty"`

	// message:
	//   type: ChatMessage
	Message ChatMessage `json:"message,omitempty"`

	// delta:
	//   type: ChatMessage
	Delta ChatMessage `json:"delta,omitempty"`

	// finish_reason:
	//   type: string
	FinishReason string `json:"finish_reason,omitempty"`
}

type ChatMessage struct {
	// role:
	//   type: string
	//   enum:
	//     - system
	//     - user
	//     - assistant
	//   description: The role of the author of this message.
	Role string `json:"role,omitempty"`

	// content:
	//   type: string
	//   description: The contents of the message
	Content string `json:"content,omitempty"`

	// name:
	//   type: string
	//   description: The name of the user in a multi-user chat
	Name string `json:"name,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Param   string `json:"param"`
	Type    string `json:"type"`
}

func (e *Error) Error() string {
	m, _ := json.Marshal(e)
	return string(m)
}

type ErrorResponse struct {
	Error Error `json:"error"`
}
