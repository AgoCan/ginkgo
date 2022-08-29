package health

import "ginkgo/internal/pkg/response"

type Health struct{}

func (h *Health) Status() response.Response {
	return response.Success("health")
}
