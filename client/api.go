package client

import (
	"github.com/geovanisouza92/lambdac/types"
)

type API interface {
	FunctionList() (types.Functions, error)
	FunctionCreate(function types.Function) (types.Function, error)
	FunctionInfo(id string) (types.Function, error)
	FunctionConfig(id string, function types.Function) error
	FunctionDestroy(id string, force bool) error
	FunctionEnv(id string) ([]string, error)
	FunctionEnvSet(id string, vars []string) error
	FunctionEnvUnset(id string, vars []string) error
	FunctionPull(id string) (string, error)
	FunctionPush(id, code string) error
	FunctionPs(id string) ([]string, error)    // TODO Confirm return type
	FunctionLogs(id string) ([]string, error)  // TODO Confirm return type
	FunctionStats(id string) ([]string, error) // TODO Confirm return type
	FunctionInvoke(id string) error            // TODO Add arguments

	RuntimeList() (types.Runtimes, error)
	RuntimeCreate(runtime types.Runtime) (types.Runtime, error)
	RuntimeInfo(id string) (types.Runtime, error)
	RuntimeDestroy(id string, force bool) error
}
