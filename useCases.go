package goUseCases

import "reflect"

// UseCase defines the contract for the use case implementation
type UseCase interface {
	Perform()
}

// Base All UseCases implement type Base
type Base struct {
	*Context
}

// Context keeps the state of the invoked use cases
type Context struct {
	Internal map[string]interface{}
	Errors   map[string]string
	Status   int // uses http status codes
}

func initContext() *Context {
	return &Context{
		Internal: make(map[string]interface{}),
		Errors:   make(map[string]string),
	}
}

// EmptyContext initializes a context with a map
func EmptyContext() *Context {
	return initContext()
}

// InitContext initializes a context with a values map
func InitContext(params *map[string]interface{}) *Context {
	return &Context{
		Internal: *params,
		Errors:   make(map[string]string),
	}
}

// Get function gets a property set on the context by its name
func (context *Context) Get(name string) interface{} {
	value := context.Internal[name]

	if value == nil {
		return nil
	}

	return value
}

// Set function sets a property on the context by its name
func (context *Context) Set(name string, value interface{}) {
	context.Internal[name] = value
}

// Depends function executes the use case dependencies
func (context *Context) Depends(dependencies ...UseCase) {
	if len(dependencies) == 0 {
		return
	}

	//initial context to the useCase dependencies
	lastContext := context

	for _, depends := range dependencies {
		setDependsContext(depends, lastContext)
		(depends).Perform()
		lastContext = getDependsContext(depends)

		if lastContext.HasFailed() {
			return
		}
	}
}

func getDependsContext(useCase UseCase) *Context {
	return reflect.ValueOf(useCase).Elem().FieldByName("Context").Interface().(*Context)
}

func setDependsContext(useCase UseCase, context *Context) {
	reflect.ValueOf(useCase).Elem().FieldByName("Context").Set(reflect.ValueOf(context))
}

// Failure signals that a UseCase Failed, assigning the status result and error message
func (context *Context) Failure(status string, message string) {
	switch status {
	case "unprocessable_entity":
		context.Status = 422
	case "not_found":
		context.Status = 404
	default:
		context.Status = 500
	}

	context.Errors[status] = message
}

// SucHasFailedcess Returns true if there were no errors invoking the UseCase
func (context *Context) HasFailed() bool {
	return context.Status != 200
}
