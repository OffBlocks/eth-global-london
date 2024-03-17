package worker

import (
	"github.com/offblocks/eth-global-london/api/pkg/activities"
	"github.com/offblocks/eth-global-london/api/pkg/blockchain"
	"github.com/offblocks/eth-global-london/api/pkg/repository"
	"github.com/offblocks/eth-global-london/api/pkg/service"
	"github.com/offblocks/eth-global-london/api/pkg/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func NewWorker(client client.Client, repository *repository.Repository, base *blockchain.BaseBlockchain, ethereum *blockchain.EthereumBlockchain) worker.Worker {
	// This worker hosts both Workflow and Activity functions
	w := worker.New(client, service.TaskQueue, worker.Options{})

	w.RegisterWorkflow(workflows.Payment)

	w.RegisterActivity(&activities.Activities{Repository: repository, BaseBlockchain: base, EthereumBlockchain: ethereum})

	return w
}
