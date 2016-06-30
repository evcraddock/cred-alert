package metrics

import "cred-alert/datadog"

//go:generate counterfeiter . Emitter

type Emitter interface {
	Counter(name string) Counter
}

func BuildEmitter(apiKey string, environment string) Emitter {
	if apiKey == "" {
		return &nullEmitter{
			environment: environment,
		}
	}

	client := datadog.NewClient(apiKey)

	return NewEmitter(client, environment)
}

func NewEmitter(dataDogClient datadog.Client, environment string) Emitter {
	return &emitter{
		client:      dataDogClient,
		environment: environment,
	}
}

type emitter struct {
	client      datadog.Client
	environment string
}

func (emitter *emitter) Counter(name string) Counter {
	return &counter{
		name:    name,
		emitter: emitter,
	}
}

type nullEmitter struct {
	name        string
	environment string
}

func (e *nullEmitter) Counter(name string) Counter {
	return &nullCounter{
		name:        e.name,
		environment: e.environment,
	}
}