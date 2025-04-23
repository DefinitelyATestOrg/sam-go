# Store

## Orders

# User

# Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageNewResponse">MessageNewResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageCountTokensResponse">MessageCountTokensResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageCountTokensBetaResponse">MessageCountTokensBetaResponse</a>

Methods:

- <code title="post /v1/messages">client.Messages.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageNewParams">MessageNewParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageNewResponse">MessageNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messages/count_tokens">client.Messages.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageService.CountTokens">CountTokens</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageCountTokensParams">MessageCountTokensParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageCountTokensResponse">MessageCountTokensResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messages/count_tokens?beta=true">client.Messages.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageService.CountTokensBeta">CountTokensBeta</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageCountTokensBetaParams">MessageCountTokensBetaParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageCountTokensBetaResponse">MessageCountTokensBetaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Batches

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchNewResponse">MessageBatchNewResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchGetResponse">MessageBatchGetResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchListResponse">MessageBatchListResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchDeleteResponse">MessageBatchDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchCancelResponse">MessageBatchCancelResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchCancelBetaResponse">MessageBatchCancelBetaResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchResultsResponse">MessageBatchResultsResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchResultsBetaResponse">MessageBatchResultsBetaResponse</a>

Methods:

- <code title="post /v1/messages/batches">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchNewParams">MessageBatchNewParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchNewResponse">MessageBatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messages/batches/{message_batch_id}">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchGetParams">MessageBatchGetParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchGetResponse">MessageBatchGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messages/batches">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchListParams">MessageBatchListParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchListResponse">MessageBatchListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messages/batches/{message_batch_id}">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchDeleteParams">MessageBatchDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchDeleteResponse">MessageBatchDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messages/batches/{message_batch_id}/cancel">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchCancelParams">MessageBatchCancelParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchCancelResponse">MessageBatchCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messages/batches/{message_batch_id}/cancel?beta=true">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchService.CancelBeta">CancelBeta</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchCancelBetaParams">MessageBatchCancelBetaParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchCancelBetaResponse">MessageBatchCancelBetaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messages/batches/{message_batch_id}/results">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchService.Results">Results</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchResultsParams">MessageBatchResultsParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchResultsResponse">MessageBatchResultsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messages/batches/{message_batch_id}/results?beta=true">client.Messages.Batches.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchService.ResultsBeta">ResultsBeta</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchResultsBetaParams">MessageBatchResultsBetaParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchResultsBetaResponse">MessageBatchResultsBetaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### BetaTrue

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchBetaTrueGetResponse">MessageBatchBetaTrueGetResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchBetaTrueDeleteResponse">MessageBatchBetaTrueDeleteResponse</a>

Methods:

- <code title="get /v1/messages/batches/{message_batch_id}?beta=true">client.Messages.Batches.BetaTrue.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchBetaTrueService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchBetaTrueGetParams">MessageBatchBetaTrueGetParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchBetaTrueGetResponse">MessageBatchBetaTrueGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messages/batches/{message_batch_id}?beta=true">client.Messages.Batches.BetaTrue.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchBetaTrueService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageBatchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchBetaTrueDeleteParams">MessageBatchBetaTrueDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchBetaTrueDeleteResponse">MessageBatchBetaTrueDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## BatchesBetaTrue

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchesBetaTrueNewResponse">MessageBatchesBetaTrueNewResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchesBetaTrueListResponse">MessageBatchesBetaTrueListResponse</a>

Methods:

- <code title="post /v1/messages/batches?beta=true">client.Messages.BatchesBetaTrue.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchesBetaTrueService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchesBetaTrueNewParams">MessageBatchesBetaTrueNewParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchesBetaTrueNewResponse">MessageBatchesBetaTrueNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messages/batches?beta=true">client.Messages.BatchesBetaTrue.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchesBetaTrueService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchesBetaTrueListParams">MessageBatchesBetaTrueListParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessageBatchesBetaTrueListResponse">MessageBatchesBetaTrueListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Complete

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#CompleteNewResponse">CompleteNewResponse</a>

Methods:

- <code title="post /v1/complete">client.Complete.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#CompleteService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#CompleteNewParams">CompleteNewParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#CompleteNewResponse">CompleteNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelGetResponse">ModelGetResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelListResponse">ModelListResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelGetBetaResponse">ModelGetBetaResponse</a>

Methods:

- <code title="get /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelGetParams">ModelGetParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelGetResponse">ModelGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelListParams">ModelListParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models/{model_id}?beta=true">client.Models.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelService.GetBeta">GetBeta</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelGetBetaParams">ModelGetBetaParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelGetBetaResponse">ModelGetBetaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# MessagesBetaTrue

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessagesBetaTrueNewResponse">MessagesBetaTrueNewResponse</a>

Methods:

- <code title="post /v1/messages?beta=true">client.MessagesBetaTrue.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessagesBetaTrueService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessagesBetaTrueNewParams">MessagesBetaTrueNewParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#MessagesBetaTrueNewResponse">MessagesBetaTrueNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ModelsBetaTrue

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelsBetaTrueListResponse">ModelsBetaTrueListResponse</a>

Methods:

- <code title="get /v1/models?beta=true">client.ModelsBetaTrue.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelsBetaTrueService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelsBetaTrueListParams">ModelsBetaTrueListParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v4#ModelsBetaTrueListResponse">ModelsBetaTrueListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SamPlopPlop
