package hackberry

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	OPERATOR_EQ string = "="
	OPERATOR_NE string = "!="
	OPERATOR_LT string = "<"
	OPERATOR_LE string = "<="
	OPERATOR_GT string = ">"
	OPERATOR_GE string = ">="
)

// default condition evaluator
type defaultConditionEvaluator struct{
	
}

func NewDefaultConditionEvaluator() defaultConditionEvaluator{
	return defaultConditionEvaluator{}
}

func (ce *defaultConditionEvaluator) IsSatisfied(condition string, context *Context) bool{
	op := getOperator(condition)
	cs := strings.Split(condition, op)
	name := strings.TrimSpace(cs[0])
	value := strings.TrimSpace(cs[1])
	
	attrValue := context.GetAttribute(name);
	if attrValue == nil { return false }
	
	switch v := attrValue.(type){
		case bool:
			return compareBool(v, parseBool(value), op)
		case int8:
			return compareInt64(int64(v), parseInt(value), op)
		case int16:
			return compareInt64(int64(v), parseInt(value), op)
		case int32:
			return compareInt64(int64(v), parseInt(value), op)
		case int64:
			return compareInt64(v, parseInt(value), op)
		case int:
			return compareInt64(int64(v), parseInt(value), op)
		case uint8:
			return compareUint64(uint64(v), parseUint(value), op)
		case uint16:
			return compareUint64(uint64(v), parseUint(value), op)
		case uint32:
			return compareUint64(uint64(v), parseUint(value), op)
		case uint64:
			return compareUint64(v, parseUint(value), op)
		case uint:
			return compareUint64(uint64(v), parseUint(value), op)
		case float32:
			return compareFloat64(float64(v), parseFloat(value), op)
		case float64:
			return compareFloat64(v, parseFloat(value), op)
		case string:
			return compareString(v, value, op)
		default:
			msg := fmt.Sprintf("Unsupported value type [%T] for condition [%s].", v, condition)
			panic(&IllegalConditionError{msg})	
	}
}

func getOperator(condition string) string{
	operators := []string{OPERATOR_NE,
			OPERATOR_LE, OPERATOR_LT,
			OPERATOR_GE, OPERATOR_GT, OPERATOR_EQ}

	for _, op := range operators {
		index := strings.Index(condition, op)
		// operator shouldn't be at the head or the tail
		if index > 0 && index < len(condition) - len(op) {
			return op
		}
	}
	
	panic(&IllegalConditionError{"Unsupported operator of condition [" + condition + "]."})
}

func parseBool(s string) bool{
	if b, err := strconv.ParseBool(s); err == nil{
		return b
	}else{
		panic(&IllegalConditionError{"Can't parse bool value in condition [" + s + "]."})
	}
}

func parseInt(s string) int64{
	if i, err := strconv.ParseInt(s, 10, 64); err == nil{
		return i
	}else{
		panic(&IllegalConditionError{"Can't parse int value in condition [" + s + "]."})
	}
}

func parseUint(s string) uint64{
	if u, err := strconv.ParseUint(s, 10, 64); err == nil{
		return u
	}else{
		panic(&IllegalConditionError{"Can't parse uint value in condition [" + s + "]."})
	}
}

func parseFloat(s string) float64{
	if f, err := strconv.ParseFloat(s, 64); err == nil{
		return f
	}else{
		panic(&IllegalConditionError{"Can't parse float value in condition [" + s + "]."})
	}
}

func compareBool(v1, v2 bool, op string) bool{
	switch op{
		case OPERATOR_EQ :
			return v1 == v2
		case OPERATOR_NE :
			return v1 != v2
		default:
			panic(&IllegalConditionError{"Unsupported bool operation [" + op + "]."})
	}
}

func compareInt64(v1, v2 int64, op string) bool{
	switch op{
		case OPERATOR_EQ :
			return v1 == v2
		case OPERATOR_NE :
			return v1 != v2
		case OPERATOR_LT :
			return v1 < v2
		case OPERATOR_LE :
			return v1 <= v2
		case OPERATOR_GT :
			return v1 > v2
		case OPERATOR_GE :
			return v1 >= v2
		default:
			return false
	}
}

func compareUint64(v1, v2 uint64, op string) bool{
	switch op{
		case OPERATOR_EQ :
			return v1 == v2
		case OPERATOR_NE :
			return v1 != v2
		case OPERATOR_LT :
			return v1 < v2
		case OPERATOR_LE :
			return v1 <= v2
		case OPERATOR_GT :
			return v1 > v2
		case OPERATOR_GE :
			return v1 >= v2
		default:
			return false
	}
}

func compareFloat64(v1, v2 float64, op string) bool{
	switch op{
		case OPERATOR_EQ :
			return v1 == v2
		case OPERATOR_NE :
			return v1 != v2
		case OPERATOR_LT :
			return v1 < v2
		case OPERATOR_LE :
			return v1 <= v2
		case OPERATOR_GT :
			return v1 > v2
		case OPERATOR_GE :
			return v1 >= v2
		default:
			return false
	}
}

func compareString(v1, v2, op string) bool{
	switch op{
		case OPERATOR_EQ :
			return v1 == v2
		case OPERATOR_NE :
			return v1 != v2
		case OPERATOR_LT :
			return v1 < v2
		case OPERATOR_LE :
			return v1 <= v2
		case OPERATOR_GT :
			return v1 > v2
		case OPERATOR_GE :
			return v1 >= v2
		default:
			return false
	}
}