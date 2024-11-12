package provider

import (
	"net/http"
	"net/http/httptest"
	"testing"

	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/tangr/pulumi-apolloconfig/provider/pkg/internal/pulumiapi"
	"github.com/tangr/pulumi-apolloconfig/provider/pkg/internal/serde"
)

func TestStackTagsUpdate(t *testing.T) {
	t.Run("Calls to Update return an error", func(t *testing.T) {
		var gotMethods []string

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotMethods = append(gotMethods, r.Method)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		apiClient, err := pulumiapi.NewClient(server.Client(), "", server.URL)
		if err != nil {
			t.Fatal(err)
		}

		st := &PulumiServiceStackTagResource{
			client: apiClient,
		}

		input := PulumiServiceStackTagInput{
			Organization: "org",
			Project:      "project",
			Stack:        "stack",
			Name:         "tag",
			Value:        "tag_value",
		}

		properties, err := serde.ToProperties(input, "pulumi")

		if err != nil {
			t.Fatal(err)
		}

		upReq := pulumirpc.UpdateRequest{
			Olds: properties,
			News: properties,
		}

		_, err = st.Update(&upReq)
		assert.ErrorContains(t, err, "unexpected call to update, expected create to be called instead")
	})

}
