# Shared Params Types

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2/shared#OrderParam">OrderParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2/shared#Order">Order</a>

# Pets

Params Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetParam">PetParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#APIResponse">APIResponse</a>
- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#Pet">Pet</a>

Methods:

- <code title="post /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetNewParams">PetNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /pet">client.Pets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetUpdateParams">PetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetDeleteParams">PetDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /pet/findByStatus">client.Pets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetService.FindByStatus">FindByStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetFindByStatusParams">PetFindByStatusParams</a>) ([]<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/findByTags">client.Pets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetService.FindByTags">FindByTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetFindByTagsParams">PetFindByTagsParams</a>) ([]<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /pet/{petId}/uploadImage">client.Pets.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetService.UploadImage">UploadImage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#PetUploadImageParams">PetUploadImageParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#APIResponse">APIResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Stores

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#StoreInventoryResponse">StoreInventoryResponse</a>

Methods:

- <code title="get /store/order/{orderId}">client.Stores.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#StoreService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /store/order/{orderId}">client.Stores.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#StoreService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /store/order">client.Stores.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#StoreService.NewOrder">NewOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#StoreNewOrderParams">StoreNewOrderParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /store/inventory">client.Stores.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#StoreService.Inventory">Inventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#StoreInventoryResponse">StoreInventoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orders

# Users

Params Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#UserParam">UserParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#User">User</a>

Methods:

- <code title="post /user">client.Users.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#UserNewParams">UserNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /user/{username}">client.Users.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /user/{username}">client.Users.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#UserUpdateParams">UserUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /user/{username}">client.Users.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /user/createWithList">client.Users.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#UserService.NewWithList">NewWithList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#UserNewWithListParams">UserNewWithListParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/login">client.Users.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2">samgo</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#UserLoginParams">UserLoginParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/logout">client.Users.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/sam-go/v2#UserService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
