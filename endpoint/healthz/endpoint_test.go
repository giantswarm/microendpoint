package healthz

import (
	"context"
	"net/http"
	"testing"

	"github.com/giantswarm/micrologger/microloggertest"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/giantswarm/microendpoint/service/healthz"
)

func Test_Endpoint_ServiceFailed_True(t *testing.T) {
	testCases := []struct {
		Failed             bool
		ExpectedStatusCode int
	}{
		{
			Failed:             false,
			ExpectedStatusCode: 0,
		},
		{
			Failed:             true,
			ExpectedStatusCode: http.StatusInternalServerError,
		},
	}

	for i, tc := range testCases {
		var encoder kithttp.EncodeResponseFunc
		var endpoint kitendpoint.Endpoint
		{
			endpointConfig := DefaultConfig()
			endpointConfig.Logger = microloggertest.New()
			endpointConfig.Services = []healthz.Service{
				&testService{Failed: tc.Failed},
			}
			newEndpoint, err := New(endpointConfig)
			if err != nil {
				t.Fatalf("test", i+1, "expected", nil, "got", err)
			}
			encoder = newEndpoint.Encoder()
			endpoint = newEndpoint.Endpoint()
		}

		res, err := endpoint(context.TODO(), nil)
		if err != nil {
			t.Fatalf("test", i+1, "expected", nil, "got", err)
		}

		w := &testWriter{}

		err = encoder(context.TODO(), w, res)
		if err != nil {
			t.Fatalf("test", i+1, "expected", nil, "got", err)
		}

		if w.StatusCode() != tc.ExpectedStatusCode {
			t.Fatalf("test", i+1, "expected", tc.ExpectedStatusCode, "got", w.StatusCode())
		}
	}
}

type testService struct {
	Failed bool
}

func (s *testService) GetHealthz(ctx context.Context) (healthz.Response, error) {
	response := healthz.Response{
		Failed: s.Failed,
	}

	return response, nil
}

type testWriter struct {
	// Settings.
	statusCode int
}

func (rw *testWriter) Header() http.Header {
	return http.Header{}
}

func (rw *testWriter) StatusCode() int {
	return rw.statusCode
}

func (rw *testWriter) Write(b []byte) (int, error) {
	return 0, nil
}

func (rw *testWriter) WriteHeader(c int) {
	rw.statusCode = c
}
