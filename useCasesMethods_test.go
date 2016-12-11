package useCases

import (
	"reflect"
	"testing"
)

func TestEmptyContext(t *testing.T) {
	tests := []struct {
		name string
		want *Context
	}{
		{
			"initializing an empty context",
			&Context{
				Internal: make(map[string]interface{}),
				Errors:   make(map[string]string),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EmptyContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmptyContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitContext(t *testing.T) {
	params := map[string]interface{}{
		"id":      1,
		"name":    "Ricardo",
		"visited": true,
	}

	type args struct {
		values *map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"initializing the context with a map of keys and values",
			args{
				&params,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InitContext(tt.args.values)

			for key, value := range got.Internal {
				if expected := value; !reflect.DeepEqual(expected, params[key]) {
					t.Errorf("Unmatched set param %v, got %v, want %v", key, value, params[key])
				}
			}
		})
	}
}

func TestContext_GettersSetters(t *testing.T) {
	type complexExample struct {
		Name string
	}

	examplePointer := &complexExample{"Complex Pointer"}

	type fields struct {
		Internal map[string]interface{}
		Errors   map[string]string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			"getting values from the context when they are simple",
			fields{
				map[string]interface{}{
					"ID":   1,
					"Name": "My Name is",
				},
				make(map[string]string),
			},
			args{
				"Name",
			},
			"My Name is",
		},
		{
			"getting values from the context when they are structs",
			fields{
				map[string]interface{}{
					"Example": complexExample{"Ricardo"},
				},
				make(map[string]string),
			},
			args{
				"Example",
			},
			complexExample{"Ricardo"},
		},
		{
			"getting values from the context when they are pointers to structs",
			fields{
				map[string]interface{}{
					"Example": examplePointer,
				},
				make(map[string]string),
			},
			args{
				"Example",
			},
			examplePointer,
		},
		// Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			context := &Context{
				Internal: tt.fields.Internal,
				Errors:   tt.fields.Errors,
			}

			context.Set(tt.args.name, tt.want)

			if got := context.Get(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Context.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_Depends(t *testing.T) {
	type fields struct {
		Internal map[string]interface{}
		Errors   map[string]string
	}
	type args struct {
		dependencies []UseCase
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			context := &Context{
				Internal: tt.fields.Internal,
				Errors:   tt.fields.Errors,
			}
			context.Depends(tt.args.dependencies...)
		})
	}
}

func Test_getDependsContext(t *testing.T) {
	type args struct {
		useCase UseCase
	}
	tests := []struct {
		name string
		args args
		want *Context
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDependsContext(tt.args.useCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDependsContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setDependsContext(t *testing.T) {
	type args struct {
		useCase UseCase
		context *Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setDependsContext(tt.args.useCase, tt.args.context)
		})
	}
}
