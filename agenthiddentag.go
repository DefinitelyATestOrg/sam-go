// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DefinitelyATestOrg/sam-go/v3/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

// AgentHiddenTagService contains methods and other services that help with
// interacting with the sam API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAgentHiddenTagService] method instead.
type AgentHiddenTagService struct {
	Options []option.RequestOption
}

// NewAgentHiddenTagService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentHiddenTagService(opts ...option.RequestOption) (r *AgentHiddenTagService) {
	r = &AgentHiddenTagService{}
	r.Options = opts
	return
}

func (r *AgentHiddenTagService) Update(ctx context.Context, id string, body AgentHiddenTagUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/agents/%s/hiddenTags", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AgentHiddenTagUpdateParams struct {
	Body []string `json:"body,required"`
}

func (r AgentHiddenTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
