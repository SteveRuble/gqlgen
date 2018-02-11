// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT

package todo

import (
	context "context"
	fmt "fmt"
	io "io"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
	sync "sync"
	time "time"

	mapstructure "github.com/mitchellh/mapstructure"
	jsonw "github.com/vektah/gqlgen/jsonw"
	errors "github.com/vektah/gqlgen/neelance/errors"
	introspection "github.com/vektah/gqlgen/neelance/introspection"
	query "github.com/vektah/gqlgen/neelance/query"
	schema "github.com/vektah/gqlgen/neelance/schema"
	validation "github.com/vektah/gqlgen/neelance/validation"
)

type Resolvers interface {
	MyMutation_createTodo(ctx context.Context, text string) (Todo, error)
	MyMutation_updateTodo(ctx context.Context, id int, changes map[string]interface{}) (*Todo, error)
	MyQuery_todo(ctx context.Context, id int) (*Todo, error)
	MyQuery_lastTodo(ctx context.Context) (*Todo, error)
	MyQuery_todos(ctx context.Context) ([]Todo, error)
}

func NewExecutor(resolvers Resolvers) func(context.Context, string, string, map[string]interface{}, io.Writer) []*errors.QueryError {
	return func(ctx context.Context, document string, operationName string, variables map[string]interface{}, w io.Writer) []*errors.QueryError {
		doc, qErr := query.Parse(document)
		if qErr != nil {
			return []*errors.QueryError{qErr}
		}

		errs := validation.Validate(parsedSchema, doc)
		if len(errs) != 0 {
			return errs
		}

		op, err := doc.GetOperation(operationName)
		if err != nil {
			return []*errors.QueryError{errors.Errorf("%s", err)}
		}

		c := executionContext{
			resolvers: resolvers,
			variables: variables,
			doc:       doc,
			ctx:       ctx,
		}

		var result jsonw.JsonWriter
		if op.Type == query.Query {
			result = c._myQuery(op.Selections, nil)
		} else if op.Type == query.Mutation {
			result = c._myMutation(op.Selections, nil)
		} else {
			return []*errors.QueryError{errors.Errorf("unsupported operation type")}
		}

		c.wg.Wait()

		writer := jsonw.New(w)
		writer.BeginObject()

		writer.ObjectKey("data")
		result.WriteJson(writer)

		if len(c.Errors) > 0 {
			writer.ObjectKey("errors")
			errors.WriteErrors(w, c.Errors)
		}

		writer.EndObject()
		return nil
	}
}

type executionContext struct {
	errors.Builder
	resolvers Resolvers
	variables map[string]interface{}
	doc       *query.Document
	ctx       context.Context
	wg        sync.WaitGroup
}

type _MyMutationNode struct {
	_fields    []collectedField
	CreateTodo jsonw.JsonWriter
	UpdateTodo jsonw.JsonWriter
}

var myMutationImplementors = []string{"MyMutation"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _myMutation(sel []query.Selection, it *interface{}) jsonw.JsonWriter {
	node := _MyMutationNode{
		_fields: ec.collectFields(sel, myMutationImplementors, map[string]bool{}),
	}

	for _, field := range node._fields {
		switch field.Name {
		case "createTodo":
			var arg0 string
			if tmp, ok := field.Args["text"]; ok {
				tmp2, err := coerceString(tmp)
				if err != nil {
					ec.Error(err)
					continue
				}
				arg0 = tmp2
			}
			ec.wg.Add(1)
			go func(field collectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.MyMutation_createTodo(ec.ctx, arg0)
				if err != nil {
					ec.Error(err)
					return
				}
				node.CreateTodo = ec._todo(field.Selections, &res)
			}(field)
		case "updateTodo":
			var arg0 int
			if tmp, ok := field.Args["id"]; ok {
				tmp2, err := coerceInt(tmp)
				if err != nil {
					ec.Error(err)
					continue
				}
				arg0 = tmp2
			}
			var arg1 map[string]interface{}
			if tmp, ok := field.Args["changes"]; ok {
				arg1 = tmp.(map[string]interface{})
			}
			ec.wg.Add(1)
			go func(field collectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.MyMutation_updateTodo(ec.ctx, arg0, arg1)
				if err != nil {
					ec.Error(err)
					return
				}
				if res != nil {
					node.UpdateTodo = ec._todo(field.Selections, res)
				}
			}(field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return &node
}

func (t *_MyMutationNode) WriteJson(w *jsonw.Writer) {
	w.BeginObject()
	for _, field := range t._fields {
		switch field.Name {
		case "createTodo":
			w.ObjectKey("createTodo")
			t.CreateTodo.WriteJson(w)
		case "updateTodo":
			w.ObjectKey("updateTodo")
			if t.UpdateTodo == nil {
				w.Null()
			} else {
				t.UpdateTodo.WriteJson(w)
			}
		}
	}
	w.EndObject()
}

type _MyQueryNode struct {
	_fields  []collectedField
	Todo     jsonw.JsonWriter
	LastTodo jsonw.JsonWriter
	Todos    []jsonw.JsonWriter
	__schema jsonw.JsonWriter
	__type   jsonw.JsonWriter
}

var myQueryImplementors = []string{"MyQuery"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _myQuery(sel []query.Selection, it *interface{}) jsonw.JsonWriter {
	node := _MyQueryNode{
		_fields: ec.collectFields(sel, myQueryImplementors, map[string]bool{}),
	}

	for _, field := range node._fields {
		switch field.Name {
		case "todo":
			var arg0 int
			if tmp, ok := field.Args["id"]; ok {
				tmp2, err := coerceInt(tmp)
				if err != nil {
					ec.Error(err)
					continue
				}
				arg0 = tmp2
			}
			ec.wg.Add(1)
			go func(field collectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.MyQuery_todo(ec.ctx, arg0)
				if err != nil {
					ec.Error(err)
					return
				}
				if res != nil {
					node.Todo = ec._todo(field.Selections, res)
				}
			}(field)
		case "lastTodo":
			ec.wg.Add(1)
			go func(field collectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.MyQuery_lastTodo(ec.ctx)
				if err != nil {
					ec.Error(err)
					return
				}
				if res != nil {
					node.LastTodo = ec._todo(field.Selections, res)
				}
			}(field)
		case "todos":
			ec.wg.Add(1)
			go func(field collectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.MyQuery_todos(ec.ctx)
				if err != nil {
					ec.Error(err)
					return
				}
				if res != nil {
					for i := range res {
						node.Todos = append(node.Todos, ec._todo(field.Selections, &res[i]))
					}
				}
			}(field)
		case "__schema":
			res := ec.introspectSchema()
			if res != nil {
				node.__schema = ec.___Schema(field.Selections, res)
			}
		case "__type":
			var arg0 string
			if tmp, ok := field.Args["name"]; ok {
				tmp2, err := coerceString(tmp)
				if err != nil {
					ec.Error(err)
					continue
				}
				arg0 = tmp2
			}
			res := ec.introspectType(arg0)
			if res != nil {
				node.__type = ec.___Type(field.Selections, res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return &node
}

func (t *_MyQueryNode) WriteJson(w *jsonw.Writer) {
	w.BeginObject()
	for _, field := range t._fields {
		switch field.Name {
		case "todo":
			w.ObjectKey("todo")
			if t.Todo == nil {
				w.Null()
			} else {
				t.Todo.WriteJson(w)
			}
		case "lastTodo":
			w.ObjectKey("lastTodo")
			if t.LastTodo == nil {
				w.Null()
			} else {
				t.LastTodo.WriteJson(w)
			}
		case "todos":
			w.ObjectKey("todos")
			w.BeginArray()
			for _, val := range t.Todos {
				val.WriteJson(w)
			}
			w.EndArray()
		case "__schema":
			w.ObjectKey("__schema")
			if t.__schema == nil {
				w.Null()
			} else {
				t.__schema.WriteJson(w)
			}
		case "__type":
			w.ObjectKey("__type")
			if t.__type == nil {
				w.Null()
			} else {
				t.__type.WriteJson(w)
			}
		}
	}
	w.EndObject()
}

type _TodoNode struct {
	_fields []collectedField
	Id      int
	Text    string
	Done    bool
}

var todoImplementors = []string{"Todo"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _todo(sel []query.Selection, it *Todo) jsonw.JsonWriter {
	node := _TodoNode{
		_fields: ec.collectFields(sel, todoImplementors, map[string]bool{}),
	}

	for _, field := range node._fields {
		switch field.Name {
		case "id":
			res := it.ID
			node.Id = res
		case "text":
			res := it.Text
			node.Text = res
		case "done":
			res := it.Done
			node.Done = res
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return &node
}

func (t *_TodoNode) WriteJson(w *jsonw.Writer) {
	w.BeginObject()
	for _, field := range t._fields {
		switch field.Name {
		case "id":
			w.ObjectKey("id")
			w.Int(t.Id)
		case "text":
			w.ObjectKey("text")
			w.String(t.Text)
		case "done":
			w.ObjectKey("done")
			w.Bool(t.Done)
		}
	}
	w.EndObject()
}

type ___DirectiveNode struct {
	_fields     []collectedField
	Name        string
	Description *string
	Locations   []string
	Args        []jsonw.JsonWriter
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(sel []query.Selection, it *introspection.Directive) jsonw.JsonWriter {
	node := ___DirectiveNode{
		_fields: ec.collectFields(sel, __DirectiveImplementors, map[string]bool{}),
	}

	for _, field := range node._fields {
		switch field.Name {
		case "name":
			res := it.Name()
			node.Name = res
		case "description":
			res := it.Description()
			node.Description = res
		case "locations":
			res := it.Locations()
			node.Locations = res
		case "args":
			res := it.Args()
			if res != nil {
				for i := range res {
					node.Args = append(node.Args, ec.___InputValue(field.Selections, res[i]))
				}
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return &node
}

func (t *___DirectiveNode) WriteJson(w *jsonw.Writer) {
	w.BeginObject()
	for _, field := range t._fields {
		switch field.Name {
		case "name":
			w.ObjectKey("name")
			w.String(t.Name)
		case "description":
			w.ObjectKey("description")
			if t.Description == nil {
				w.Null()
			} else {
				w.String(*t.Description)
			}
		case "locations":
			w.ObjectKey("locations")
			w.BeginArray()
			for _, val := range t.Locations {
				w.String(val)
			}
			w.EndArray()
		case "args":
			w.ObjectKey("args")
			w.BeginArray()
			for _, val := range t.Args {
				if val == nil {
					w.Null()
				} else {
					val.WriteJson(w)
				}
			}
			w.EndArray()
		}
	}
	w.EndObject()
}

type ___EnumValueNode struct {
	_fields           []collectedField
	Name              string
	Description       *string
	IsDeprecated      bool
	DeprecationReason *string
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(sel []query.Selection, it *introspection.EnumValue) jsonw.JsonWriter {
	node := ___EnumValueNode{
		_fields: ec.collectFields(sel, __EnumValueImplementors, map[string]bool{}),
	}

	for _, field := range node._fields {
		switch field.Name {
		case "name":
			res := it.Name()
			node.Name = res
		case "description":
			res := it.Description()
			node.Description = res
		case "isDeprecated":
			res := it.IsDeprecated()
			node.IsDeprecated = res
		case "deprecationReason":
			res := it.DeprecationReason()
			node.DeprecationReason = res
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return &node
}

func (t *___EnumValueNode) WriteJson(w *jsonw.Writer) {
	w.BeginObject()
	for _, field := range t._fields {
		switch field.Name {
		case "name":
			w.ObjectKey("name")
			w.String(t.Name)
		case "description":
			w.ObjectKey("description")
			if t.Description == nil {
				w.Null()
			} else {
				w.String(*t.Description)
			}
		case "isDeprecated":
			w.ObjectKey("isDeprecated")
			w.Bool(t.IsDeprecated)
		case "deprecationReason":
			w.ObjectKey("deprecationReason")
			if t.DeprecationReason == nil {
				w.Null()
			} else {
				w.String(*t.DeprecationReason)
			}
		}
	}
	w.EndObject()
}

type ___FieldNode struct {
	_fields           []collectedField
	Name              string
	Description       *string
	Args              []jsonw.JsonWriter
	Type              jsonw.JsonWriter
	IsDeprecated      bool
	DeprecationReason *string
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(sel []query.Selection, it *introspection.Field) jsonw.JsonWriter {
	node := ___FieldNode{
		_fields: ec.collectFields(sel, __FieldImplementors, map[string]bool{}),
	}

	for _, field := range node._fields {
		switch field.Name {
		case "name":
			res := it.Name()
			node.Name = res
		case "description":
			res := it.Description()
			node.Description = res
		case "args":
			res := it.Args()
			if res != nil {
				for i := range res {
					node.Args = append(node.Args, ec.___InputValue(field.Selections, res[i]))
				}
			}
		case "type":
			res := it.Type()
			if res != nil {
				node.Type = ec.___Type(field.Selections, res)
			}
		case "isDeprecated":
			res := it.IsDeprecated()
			node.IsDeprecated = res
		case "deprecationReason":
			res := it.DeprecationReason()
			node.DeprecationReason = res
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return &node
}

func (t *___FieldNode) WriteJson(w *jsonw.Writer) {
	w.BeginObject()
	for _, field := range t._fields {
		switch field.Name {
		case "name":
			w.ObjectKey("name")
			w.String(t.Name)
		case "description":
			w.ObjectKey("description")
			if t.Description == nil {
				w.Null()
			} else {
				w.String(*t.Description)
			}
		case "args":
			w.ObjectKey("args")
			w.BeginArray()
			for _, val := range t.Args {
				if val == nil {
					w.Null()
				} else {
					val.WriteJson(w)
				}
			}
			w.EndArray()
		case "type":
			w.ObjectKey("type")
			if t.Type == nil {
				w.Null()
			} else {
				t.Type.WriteJson(w)
			}
		case "isDeprecated":
			w.ObjectKey("isDeprecated")
			w.Bool(t.IsDeprecated)
		case "deprecationReason":
			w.ObjectKey("deprecationReason")
			if t.DeprecationReason == nil {
				w.Null()
			} else {
				w.String(*t.DeprecationReason)
			}
		}
	}
	w.EndObject()
}

type ___InputValueNode struct {
	_fields      []collectedField
	Name         string
	Description  *string
	Type         jsonw.JsonWriter
	DefaultValue *string
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(sel []query.Selection, it *introspection.InputValue) jsonw.JsonWriter {
	node := ___InputValueNode{
		_fields: ec.collectFields(sel, __InputValueImplementors, map[string]bool{}),
	}

	for _, field := range node._fields {
		switch field.Name {
		case "name":
			res := it.Name()
			node.Name = res
		case "description":
			res := it.Description()
			node.Description = res
		case "type":
			res := it.Type()
			if res != nil {
				node.Type = ec.___Type(field.Selections, res)
			}
		case "defaultValue":
			res := it.DefaultValue()
			node.DefaultValue = res
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return &node
}

func (t *___InputValueNode) WriteJson(w *jsonw.Writer) {
	w.BeginObject()
	for _, field := range t._fields {
		switch field.Name {
		case "name":
			w.ObjectKey("name")
			w.String(t.Name)
		case "description":
			w.ObjectKey("description")
			if t.Description == nil {
				w.Null()
			} else {
				w.String(*t.Description)
			}
		case "type":
			w.ObjectKey("type")
			if t.Type == nil {
				w.Null()
			} else {
				t.Type.WriteJson(w)
			}
		case "defaultValue":
			w.ObjectKey("defaultValue")
			if t.DefaultValue == nil {
				w.Null()
			} else {
				w.String(*t.DefaultValue)
			}
		}
	}
	w.EndObject()
}

type ___SchemaNode struct {
	_fields          []collectedField
	Types            []jsonw.JsonWriter
	QueryType        jsonw.JsonWriter
	MutationType     jsonw.JsonWriter
	SubscriptionType jsonw.JsonWriter
	Directives       []jsonw.JsonWriter
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(sel []query.Selection, it *introspection.Schema) jsonw.JsonWriter {
	node := ___SchemaNode{
		_fields: ec.collectFields(sel, __SchemaImplementors, map[string]bool{}),
	}

	for _, field := range node._fields {
		switch field.Name {
		case "types":
			res := it.Types()
			if res != nil {
				for i := range res {
					node.Types = append(node.Types, ec.___Type(field.Selections, res[i]))
				}
			}
		case "queryType":
			res := it.QueryType()
			if res != nil {
				node.QueryType = ec.___Type(field.Selections, res)
			}
		case "mutationType":
			res := it.MutationType()
			if res != nil {
				node.MutationType = ec.___Type(field.Selections, res)
			}
		case "subscriptionType":
			res := it.SubscriptionType()
			if res != nil {
				node.SubscriptionType = ec.___Type(field.Selections, res)
			}
		case "directives":
			res := it.Directives()
			if res != nil {
				for i := range res {
					node.Directives = append(node.Directives, ec.___Directive(field.Selections, res[i]))
				}
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return &node
}

func (t *___SchemaNode) WriteJson(w *jsonw.Writer) {
	w.BeginObject()
	for _, field := range t._fields {
		switch field.Name {
		case "types":
			w.ObjectKey("types")
			w.BeginArray()
			for _, val := range t.Types {
				if val == nil {
					w.Null()
				} else {
					val.WriteJson(w)
				}
			}
			w.EndArray()
		case "queryType":
			w.ObjectKey("queryType")
			if t.QueryType == nil {
				w.Null()
			} else {
				t.QueryType.WriteJson(w)
			}
		case "mutationType":
			w.ObjectKey("mutationType")
			if t.MutationType == nil {
				w.Null()
			} else {
				t.MutationType.WriteJson(w)
			}
		case "subscriptionType":
			w.ObjectKey("subscriptionType")
			if t.SubscriptionType == nil {
				w.Null()
			} else {
				t.SubscriptionType.WriteJson(w)
			}
		case "directives":
			w.ObjectKey("directives")
			w.BeginArray()
			for _, val := range t.Directives {
				if val == nil {
					w.Null()
				} else {
					val.WriteJson(w)
				}
			}
			w.EndArray()
		}
	}
	w.EndObject()
}

type ___TypeNode struct {
	_fields       []collectedField
	Kind          string
	Name          *string
	Description   *string
	Fields        []jsonw.JsonWriter
	Interfaces    []jsonw.JsonWriter
	PossibleTypes []jsonw.JsonWriter
	EnumValues    []jsonw.JsonWriter
	InputFields   []jsonw.JsonWriter
	OfType        jsonw.JsonWriter
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(sel []query.Selection, it *introspection.Type) jsonw.JsonWriter {
	node := ___TypeNode{
		_fields: ec.collectFields(sel, __TypeImplementors, map[string]bool{}),
	}

	for _, field := range node._fields {
		switch field.Name {
		case "kind":
			res := it.Kind()
			node.Kind = res
		case "name":
			res := it.Name()
			node.Name = res
		case "description":
			res := it.Description()
			node.Description = res
		case "fields":
			var arg0 bool
			if tmp, ok := field.Args["includeDeprecated"]; ok {
				tmp2, err := coerceBool(tmp)
				if err != nil {
					ec.Error(err)
					continue
				}
				arg0 = tmp2
			}
			res := it.Fields(arg0)
			if res != nil {
				for i := range res {
					node.Fields = append(node.Fields, ec.___Field(field.Selections, res[i]))
				}
			}
		case "interfaces":
			res := it.Interfaces()
			if res != nil {
				for i := range res {
					node.Interfaces = append(node.Interfaces, ec.___Type(field.Selections, res[i]))
				}
			}
		case "possibleTypes":
			res := it.PossibleTypes()
			if res != nil {
				for i := range res {
					node.PossibleTypes = append(node.PossibleTypes, ec.___Type(field.Selections, res[i]))
				}
			}
		case "enumValues":
			var arg0 bool
			if tmp, ok := field.Args["includeDeprecated"]; ok {
				tmp2, err := coerceBool(tmp)
				if err != nil {
					ec.Error(err)
					continue
				}
				arg0 = tmp2
			}
			res := it.EnumValues(arg0)
			if res != nil {
				for i := range res {
					node.EnumValues = append(node.EnumValues, ec.___EnumValue(field.Selections, res[i]))
				}
			}
		case "inputFields":
			res := it.InputFields()
			if res != nil {
				for i := range res {
					node.InputFields = append(node.InputFields, ec.___InputValue(field.Selections, res[i]))
				}
			}
		case "ofType":
			res := it.OfType()
			if res != nil {
				node.OfType = ec.___Type(field.Selections, res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return &node
}

func (t *___TypeNode) WriteJson(w *jsonw.Writer) {
	w.BeginObject()
	for _, field := range t._fields {
		switch field.Name {
		case "kind":
			w.ObjectKey("kind")
			w.String(t.Kind)
		case "name":
			w.ObjectKey("name")
			if t.Name == nil {
				w.Null()
			} else {
				w.String(*t.Name)
			}
		case "description":
			w.ObjectKey("description")
			if t.Description == nil {
				w.Null()
			} else {
				w.String(*t.Description)
			}
		case "fields":
			w.ObjectKey("fields")
			w.BeginArray()
			for _, val := range t.Fields {
				if val == nil {
					w.Null()
				} else {
					val.WriteJson(w)
				}
			}
			w.EndArray()
		case "interfaces":
			w.ObjectKey("interfaces")
			w.BeginArray()
			for _, val := range t.Interfaces {
				if val == nil {
					w.Null()
				} else {
					val.WriteJson(w)
				}
			}
			w.EndArray()
		case "possibleTypes":
			w.ObjectKey("possibleTypes")
			w.BeginArray()
			for _, val := range t.PossibleTypes {
				if val == nil {
					w.Null()
				} else {
					val.WriteJson(w)
				}
			}
			w.EndArray()
		case "enumValues":
			w.ObjectKey("enumValues")
			w.BeginArray()
			for _, val := range t.EnumValues {
				if val == nil {
					w.Null()
				} else {
					val.WriteJson(w)
				}
			}
			w.EndArray()
		case "inputFields":
			w.ObjectKey("inputFields")
			w.BeginArray()
			for _, val := range t.InputFields {
				if val == nil {
					w.Null()
				} else {
					val.WriteJson(w)
				}
			}
			w.EndArray()
		case "ofType":
			w.ObjectKey("ofType")
			if t.OfType == nil {
				w.Null()
			} else {
				t.OfType.WriteJson(w)
			}
		}
	}
	w.EndObject()
}

var parsedSchema = schema.MustParse("schema {\n\tquery: MyQuery\n\tmutation: MyMutation\n}\n\ntype MyQuery {\n\ttodo(id: Int!): Todo\n\tlastTodo: Todo\n\ttodos: [Todo!]!\n}\n\ntype MyMutation {\n\tcreateTodo(text: String!): Todo!\n\tupdateTodo(id: Int!, changes: TodoInput!): Todo\n}\n\ntype Todo {\n\tid: Int!\n\ttext: String!\n\tdone: Boolean!\n}\n\ninput TodoInput {\n\ttext: String\n\tdone: Boolean\n}\n")

func (ec *executionContext) introspectSchema() *introspection.Schema {
	return introspection.WrapSchema(parsedSchema)
}

func (ec *executionContext) introspectType(name string) *introspection.Type {
	t := parsedSchema.Resolve(name)
	if t == nil {
		return nil
	}
	return introspection.WrapType(t)
}

func instanceOf(val string, satisfies []string) bool {
	for _, s := range satisfies {
		if val == s {
			return true
		}
	}
	return false
}

func (ec *executionContext) collectFields(selSet []query.Selection, satisfies []string, visited map[string]bool) []collectedField {
	var groupedFields []collectedField

	for _, sel := range selSet {
		switch sel := sel.(type) {
		case *query.Field:
			f := getOrCreateField(&groupedFields, sel.Name.Name, func() collectedField {
				f := collectedField{
					Alias: sel.Alias.Name,
					Name:  sel.Name.Name,
				}
				if len(sel.Arguments) > 0 {
					f.Args = map[string]interface{}{}
					for _, arg := range sel.Arguments {
						f.Args[arg.Name.Name] = arg.Value.Value(ec.variables)
					}
				}
				return f
			})

			f.Selections = append(f.Selections, sel.Selections...)
		case *query.InlineFragment:
			if !instanceOf(sel.On.Ident.Name, satisfies) {
				continue
			}

			for _, childField := range ec.collectFields(sel.Selections, satisfies, visited) {
				f := getOrCreateField(&groupedFields, childField.Name, func() collectedField { return childField })
				f.Selections = append(f.Selections, childField.Selections...)
			}

		case *query.FragmentSpread:
			fragmentName := sel.Name.Name
			if _, seen := visited[fragmentName]; seen {
				continue
			}
			visited[fragmentName] = true

			fragment := ec.doc.Fragments.Get(fragmentName)
			if fragment == nil {
				ec.Errorf("missing fragment %s", fragmentName)
				continue
			}

			if !instanceOf(fragment.On.Ident.Name, satisfies) {
				continue
			}

			for _, childField := range ec.collectFields(fragment.Selections, satisfies, visited) {
				f := getOrCreateField(&groupedFields, childField.Name, func() collectedField { return childField })
				f.Selections = append(f.Selections, childField.Selections...)
			}

		default:
			panic(fmt.Errorf("unsupported %T", sel))
		}
	}

	return groupedFields
}

type collectedField struct {
	Alias      string
	Name       string
	Args       map[string]interface{}
	Selections []query.Selection
}

func decodeHook(sourceType reflect.Type, destType reflect.Type, value interface{}) (interface{}, error) {
	if destType.PkgPath() == "time" && destType.Name() == "Time" {
		if dateStr, ok := value.(string); ok {
			return time.Parse(time.RFC3339, dateStr)
		}
		return nil, errors.Errorf("time should be an RFC3339 formatted string")
	}
	return value, nil
}

// nolint: deadcode, megacheck
func unpackComplexArg(result interface{}, data interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:     "graphql",
		ErrorUnused: true,
		Result:      result,
		DecodeHook:  decodeHook,
	})
	if err != nil {
		panic(err)
	}

	return decoder.Decode(data)
}

func getOrCreateField(c *[]collectedField, name string, creator func() collectedField) *collectedField {
	for i, cf := range *c {
		if cf.Alias == name {
			return &(*c)[i]
		}
	}

	f := creator()

	*c = append(*c, f)
	return &(*c)[len(*c)-1]
}

// nolint: deadcode, megacheck
func coerceString(v interface{}) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case int:
		return strconv.Itoa(v), nil
	case float64:
		return fmt.Sprintf("%f", v), nil
	case bool:
		if v {
			return "true", nil
		} else {
			return "false", nil
		}
	case nil:
		return "null", nil
	default:
		return "", fmt.Errorf("%T is not a string", v)
	}
}

// nolint: deadcode, megacheck
func coerceBool(v interface{}) (bool, error) {
	switch v := v.(type) {
	case string:
		return "true" == strings.ToLower(v), nil
	case int:
		return v != 0, nil
	case bool:
		return v, nil
	default:
		return false, fmt.Errorf("%T is not a bool", v)
	}
}

// nolint: deadcode, megacheck
func coerceInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case string:
		return strconv.Atoi(v)
	case int:
		return v, nil
	case float64:
		return int(v), nil
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}

// nolint: deadcode, megacheck
func coercefloat64(v interface{}) (float64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("%T is not an float", v)
	}
}
