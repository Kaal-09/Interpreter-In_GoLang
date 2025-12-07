package evaluator

import (
	"github.com/Kaal-09/Interpreter-In_GoLang/02/src/monkey/ast"
	"github.com/Kaal-09/Interpreter-In_GoLang/03/src/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}
