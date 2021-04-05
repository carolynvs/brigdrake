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

	event, err := brigade.GetEventPayload()
	if err != nil {
		log.Fatal(err)
	}

	// TODO(carolynvs): Make this configurable
	var allowInsecure bool = true
	apiClient := core.NewAPIClient(event.Worker.ApiAddress, event.Worker.ApiToken, &restmachinery.APIClientOptions{AllowInsecureConnections: allowInsecure})

	ctx := signals.Context()
	if err = executor.ExecuteBuild(
		ctx,
		event,
		apiClient,
	); err != nil {
		log.Fatal(err)
	}
}
