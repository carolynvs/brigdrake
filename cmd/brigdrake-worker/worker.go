package main

import (
	"log"

	"github.com/lovethedrake/brigdrake/pkg/brigade"
	"github.com/lovethedrake/brigdrake/pkg/brigade/executor"
	"github.com/lovethedrake/brigdrake/pkg/signals"
	"github.com/lovethedrake/brigdrake/pkg/version"
	"github.com/lovethedrake/drakecore/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {

	log.Printf(
		"Starting BrigDrake worker -- version %s -- commit %s -- supports "+
			"DrakeSpec %s",
		version.Version(),
		version.Commit(),
		config.SupportedSpecVersions,
	)

	clientConfig, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}
	kubeClient, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	workerConfig, err := brigade.GetWorkerConfigFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	project, err := brigade.GetProjectFromEnvironmentAndSecret(kubeClient)
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
		kubeClient,
	); err != nil {
		log.Fatal(err)
	}
}
