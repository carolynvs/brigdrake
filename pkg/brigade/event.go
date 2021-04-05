package brigade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/brigadecore/brigade/sdk/v2/core"
	"github.com/pkg/errors"
)

// Event represents a Brigade event.
type Event struct {
	/** A unique identifier for the event. */
	ID string
	/** The project that registered the handler being called for the event. */
	Project Project
	/** The unique identifier of the gateway which created the event. */
	Source string
	/** The type of event. Values and meanings are source-specific. */
	Type string
	/** A short title for the event, suitable for display in space-limited UI such as lists. */
	ShortTitle string
	/** A detailed title for the event. */
	LongTitle string
	/** The content of the event. This is source- and type-specific. */
	Payload string
	/** The Brigade worker assigned to handle the event. */
	Worker Worker
}

type Project struct {
	/** The unique identifier of the project. */
	ID string
	/** A map of secrets defined in the Brigade project. */
	Secrets map[string]string
}

type Worker struct {
	/** The address of the Brigade API server. */
	ApiAddress string
	/**
	 * A token which can be used to authenticate to the API server.
	 * The token is specific to the current event and allows you to create
	 * jobs for that event. It has no other permissions.
	 */
	ApiToken string
	/**
	 * The directory where the worker stores configuration files,
	 * including event handler code files such as brigade.js and brigade.json.
	 */
	ConfigFilesDirectory string
	/**
	 * The default values to use for any configuration files that are not present.
	 */
	DefaultConfigFiles map[string]string
	/**
	 * The desired granularity of worker logs. Worker logs are distinct from job
	 * logs - the containers in a job will emit logs according to their own
	 * configuration.
	 */
	LogLevel string
	/**
	 * If applicable, contains git-specific Worker details.
	 */
	Git core.GitConfig
}

// Revision represents VCS-related details.
type Revision struct {
	// Commit is the VCS commit ID (e.g. the Git commit)
	Commit string `envconfig:"BRIGADE_COMMIT_ID"`
	// Ref is the VCS full reference, defaults to `refs/heads/master`
	Ref string `envconfig:"BRIGADE_COMMIT_REF"`
}

// GetEventPayload returns an EventPayload object with values derived from
// /var/event/event.json
func GetEventPayload() (Event, error) {
	payloadPath := "/var/event/event.json"
	contents, err := ioutil.ReadFile(payloadPath)
	if err != nil {
		return Event{}, fmt.Errorf("error reading %s", payloadPath)
	}

	evt := Event{}
	err = json.Unmarshal(contents, &evt)
	return evt, errors.Wrap(err, "error loading event payload json")
}
