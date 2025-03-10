# Store

## Orders

# User

# Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageNewResponse">MessageNewResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageCountTokensResponse">MessageCountTokensResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageCountTokensBetaResponse">MessageCountTokensBetaResponse</a>

Methods:

- <code title="post /v1/messages">client.Messages.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageNewParams">MessageNewParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageNewResponse">MessageNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messages/count_tokens">client.Messages.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageService.CountTokens">CountTokens</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageCountTokensParams">MessageCountTokensParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageCountTokensResponse">MessageCountTokensResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messages/count_tokens?beta=true">client.Messages.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageService.CountTokensBeta">CountTokensBeta</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageCountTokensBetaParams">MessageCountTokensBetaParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageCountTokensBetaResponse">MessageCountTokensBetaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Batches

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchNewResponse">MessageBatchNewResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchGetResponse">MessageBatchGetResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchListResponse">MessageBatchListResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchDeleteResponse">MessageBatchDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchCancelResponse">MessageBatchCancelResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchCancelBetaResponse">MessageBatchCancelBetaResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchResultsResponse">MessageBatchResultsResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchResultsBetaResponse">MessageBatchResultsBetaResponse</a>

Methods:

- <code title="post /v1/messages/batches">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchNewParams">MessageBatchNewParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchNewResponse">MessageBatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messages/batches/{message_batch_id}">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchGetParams">MessageBatchGetParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchGetResponse">MessageBatchGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messages/batches">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchListParams">MessageBatchListParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchListResponse">MessageBatchListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messages/batches/{message_batch_id}">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchDeleteParams">MessageBatchDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchDeleteResponse">MessageBatchDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messages/batches/{message_batch_id}/cancel">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchCancelParams">MessageBatchCancelParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchCancelResponse">MessageBatchCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messages/batches/{message_batch_id}/cancel?beta=true">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchService.CancelBeta">CancelBeta</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchCancelBetaParams">MessageBatchCancelBetaParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchCancelBetaResponse">MessageBatchCancelBetaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messages/batches/{message_batch_id}/results">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchService.Results">Results</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchResultsParams">MessageBatchResultsParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchResultsResponse">MessageBatchResultsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messages/batches/{message_batch_id}/results?beta=true">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchService.ResultsBeta">ResultsBeta</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchResultsBetaParams">MessageBatchResultsBetaParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchResultsBetaResponse">MessageBatchResultsBetaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### BetaTrue

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchBetaTrueGetResponse">MessageBatchBetaTrueGetResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchBetaTrueDeleteResponse">MessageBatchBetaTrueDeleteResponse</a>

Methods:

- <code title="get /v1/messages/batches/{message_batch_id}?beta=true">client.Messages.Batches.BetaTrue.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchBetaTrueService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchBetaTrueGetParams">MessageBatchBetaTrueGetParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchBetaTrueGetResponse">MessageBatchBetaTrueGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messages/batches/{message_batch_id}?beta=true">client.Messages.Batches.BetaTrue.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchBetaTrueService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchBetaTrueDeleteParams">MessageBatchBetaTrueDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchBetaTrueDeleteResponse">MessageBatchBetaTrueDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## BatchesBetaTrue

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchesBetaTrueNewResponse">MessageBatchesBetaTrueNewResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchesBetaTrueListResponse">MessageBatchesBetaTrueListResponse</a>

Methods:

- <code title="post /v1/messages/batches?beta=true">client.Messages.BatchesBetaTrue.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchesBetaTrueService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchesBetaTrueNewParams">MessageBatchesBetaTrueNewParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchesBetaTrueNewResponse">MessageBatchesBetaTrueNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messages/batches?beta=true">client.Messages.BatchesBetaTrue.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchesBetaTrueService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchesBetaTrueListParams">MessageBatchesBetaTrueListParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessageBatchesBetaTrueListResponse">MessageBatchesBetaTrueListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Complete

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#CompleteNewResponse">CompleteNewResponse</a>

Methods:

- <code title="post /v1/complete">client.Complete.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#CompleteService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#CompleteNewParams">CompleteNewParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#CompleteNewResponse">CompleteNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelGetResponse">ModelGetResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelListResponse">ModelListResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelGetBetaResponse">ModelGetBetaResponse</a>

Methods:

- <code title="get /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelGetParams">ModelGetParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelGetResponse">ModelGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelListParams">ModelListParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models/{model_id}?beta=true">client.Models.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelService.GetBeta">GetBeta</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelGetBetaParams">ModelGetBetaParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelGetBetaResponse">ModelGetBetaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# MessagesBetaTrue

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessagesBetaTrueNewResponse">MessagesBetaTrueNewResponse</a>

Methods:

- <code title="post /v1/messages?beta=true">client.MessagesBetaTrue.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessagesBetaTrueService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessagesBetaTrueNewParams">MessagesBetaTrueNewParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#MessagesBetaTrueNewResponse">MessagesBetaTrueNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ModelsBetaTrue

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelsBetaTrueListResponse">ModelsBetaTrueListResponse</a>

Methods:

- <code title="get /v1/models?beta=true">client.ModelsBetaTrue.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelsBetaTrueService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelsBetaTrueListParams">ModelsBetaTrueListParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#ModelsBetaTrueListResponse">ModelsBetaTrueListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
