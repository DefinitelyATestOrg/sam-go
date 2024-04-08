# ReferenceSets

Methods:

- <code title="get /api/v1/referencesets/{id}">client.ReferenceSets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ReferenceSetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/v1/referencesets/{id}">client.ReferenceSets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ReferenceSetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ReferenceSetUpdateParams">ReferenceSetUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/referencesets/{id}">client.ReferenceSets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ReferenceSetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# ReferenceSessions

Methods:

- <code title="get /api/v1/referencesessions/{id}">client.ReferenceSessions.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ReferenceSessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/v1/referencesessions/{id}">client.ReferenceSessions.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ReferenceSessionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ReferenceSessionUpdateParams">ReferenceSessionUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/referencesessions/{id}">client.ReferenceSessions.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ReferenceSessionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Organizations

Methods:

- <code title="put /api/v1/organizations">client.Organizations.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#OrganizationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#OrganizationUpdateParams">OrganizationUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Members

Methods:

- <code title="put /api/v1/members/{memberId}">client.Members.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#MemberService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, memberID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#MemberUpdateParams">MemberUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/members/{memberId}">client.Members.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#MemberService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, memberID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Feedbacks

Methods:

- <code title="put /api/v1/feedbacks/{feedback_id}">client.Feedbacks.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#FeedbackService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, feedbackID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#FeedbackUpdateParams">FeedbackUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Documents

# Corpora

Methods:

- <code title="get /api/v1/corpora/{corpus_id}">client.Corpora.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#CorporaService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, corpusID <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/v1/corpora/{corpus_id}">client.Corpora.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#CorporaService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, corpusID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#CorporaUpdateParams">CorporaUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/corpora/{corpus_id}">client.Corpora.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#CorporaService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, corpusID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Agents

Methods:

- <code title="get /api/v1/agents/{id}">client.Agents.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#AgentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/v1/agents/{id}">client.Agents.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#AgentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#AgentUpdateParams">AgentUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/agents/{id}">client.Agents.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#AgentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## HiddenTags

## Configs

Methods:

- <code title="get /api/v1/agents/{agentId}/configs/{integration}">client.Agents.Configs.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#AgentConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, integration <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#AgentConfigGetParamsIntegration">AgentConfigGetParamsIntegration</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/v1/agents/{agentId}/configs/{integration}">client.Agents.Configs.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#AgentConfigService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, integration <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#AgentConfigUpdateParamsIntegration">AgentConfigUpdateParamsIntegration</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#AgentConfigUpdateParams">AgentConfigUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Chat

Methods:

- <code title="get /api/v1/agents/{agentId}/configs/chat">client.Agents.Configs.Chat.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#AgentConfigChatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/v1/agents/{agentId}/configs/chat">client.Agents.Configs.Chat.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#AgentConfigChatService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#AgentConfigChatUpdateParams">AgentConfigChatUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ActionSets

Methods:

- <code title="get /api/v1/actionsets/{id}">client.ActionSets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ActionSetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/v1/actionsets/{id}">client.ActionSets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ActionSetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ActionSetUpdateParams">ActionSetUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/actionsets/{id}">client.ActionSets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ActionSetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Actions

Methods:

- <code title="get /api/v1/actions/{actionId}">client.Actions.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ActionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, actionID <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/v1/actions/{actionId}">client.Actions.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ActionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, actionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v3#ActionUpdateParams">ActionUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
