package main

import (
	"log"

	"github.com/brigadecore/brigade/sdk/v2/core"
	"github.com/brigadecore/brigade/sdk/v2/restmachinery"
	"github.com/lovethedrake/brigdrake/pkg/brigade"
	"github.com/lovethedrake/brigdrake/pkg/brigade/executor"
	"github.com/lovethedrake/brigdrake/pkg/signals"
	"github.com/lovethedrake/brigdrake/pkg/version"
	"github.com/lovethedrake/drakecore/config"
)

func main() {

	log.Printf(
		"Starting BrigDrake worker -- version %s -- commit %s -- supports "+
			"DrakeSpec %s",
		version.Version(),
		version.Commit(),
		config.SupportedSpecVersions,
	)

	var endpoint, token string
	var allowInsecure bool = true
	apiClient := core.NewAPIClient(endpoint, token, &restmachinery.APIClientOptions{AllowInsecureConnections: allowInsecure})

	workerConfig, err := brigade.GetWorkerConfigFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	project, err := brigade.GetProjectFromEnvironmentAndSecret(apiClient.Projects())
	if err != nil {
		log.Fatal(err)
	}

	event, err := brigade.GetEventFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	ctx := signals.Context()
	if err = executor.ExecuteBuild(
		ctx,
		project,
		event,
		workerConfig,
		apiClient,
	); err != nil {
		log.Fatal(err)
	}
}
