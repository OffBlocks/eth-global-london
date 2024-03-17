package workflows

import (
	"github.com/google/uuid"
	"github.com/offblocks/offblocks-common/types"
)

func NewPaymentWorkflowId() types.UUID {
	return types.UUID{UUID: uuid.New()}
}
