# Store

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#Order">Order</a>

## Orders

Methods:

- <code title="get /store/order/{orderId}">client.Store.Orders.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#StoreOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /store/order/{orderId}">client.Store.Orders.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#StoreOrderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# User

Params Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#UserParam">UserParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#User">User</a>

Methods:

- <code title="post /user">client.User.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#UserNewParams">UserNewParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#UserUpdateParams">UserUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /user/createWithList">client.User.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#UserService.NewList">NewList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#UserNewListParams">UserNewListParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/login">client.User.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go">sam</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#UserLoginParams">UserLoginParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/logout">client.User.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go#UserService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
