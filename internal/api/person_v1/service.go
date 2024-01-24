package person_v1

import "github.com/kom1ssar/tech_em/internal/api"

var _ api.PersonV1Implementation = (*implementation)(nil)

type implementation struct {
}

func NewImplementation() api.PersonV1Implementation {
	return &implementation{}
}
