package shared

import "errors"

var (
	DoesNotExistsException       = errors.New("reqired domain entity does not exists")
	NotAllowedException          = errors.New("current behavior now allowed")
	InvalidArgumentException     = errors.New("invalid argument(s)")
	UnknownException             = errors.New("unknown nature of exception")
	InvalidDataException         = errors.New("error during work with data. Probably some of them invalid type or values")
	InfrastructureException      = errors.New("Something Wrong inside infrastructure side")
	ScenarioAlreadyDoneException = errors.New("Current call already done")
)
