package gitlabapi

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Variables []Variable

// IsVarPresent checks if a variable is present in a list of variables
func IsVarPresent(vars []Variable, key string) bool {
	for i := range vars {
		if vars[i].Key == key {
			return true
		}
	}
	return false
}

// GetIndexOfVar returns the index of a variable in a list of variables
func GetIndexOfVar(vars []Variable, key string) int {
	for i := range vars {
		if vars[i].Key == key {
			return i
		}
	}
	return -1
}

// RemoveVar removes a variable from a list of variables
func RemoveVar(vars []Variable, idx int) (res []Variable) {
	for i := range vars {
		if i != idx {
			res = append(res, vars[i])
		}
	}
	return res
}

// GetVarOfScope returns the variable of a scope
func (v Variables) GetVarOfScope(key string, scope string) (Variable, error) {
	for i := range v {
		if re, err := regexp.MatchString(v[i].EnvironmentScope, scope); err == nil && re && v[i].Key == key {
			return v[i], nil
		} else if err != nil {
			if v[i].Key == key {
				return v[i], nil
			}
		}
	}
	return Variable{}, fmt.Errorf("Variable %s not found", key)
}

// GetVarValue returns the value of a variable
func (v Variables) GetVarValue(key string) (string, error) {
	for i := range v {
		if v[i].Key == key {
			return v[i].Value, nil
		}
	}
	return "", fmt.Errorf("Variable %s not found", key)
}

// function EscapeBashCharacters espace bash characters
func EscapeBashCharacters(s string) string {
	s = strings.Replace(s, "\\", "\\\\", -1)
	s = strings.Replace(s, "\"", "\\\"", -1)
	s = strings.Replace(s, "`", "\\`", -1)
	s = strings.Replace(s, "$", "\\$", -1)
	// replace newline with \n
	// s = strings.Replace(s, "\n", "\\n", -1)
	return s
}

// IsVarPartOfScope is checking if the environment is part of the scope
// for example if the environment is "staging" and the scope is "staging" it returns true
// if the environment is "staging" and the scope is "staging*" it returns true
// if the environment is "staging/mtrg" and the scope is "staging*" it returns true
// if the environment is "staging/mtrg" and the scope is "production" it returns false
func IsVarPartOfScope(environment string, varScope string) bool {
	if environment == "" {
		return true
	}

	if varScope == "*" || varScope == "" {
		return true
	}

	if environment == varScope {
		return true
	}
	if strings.HasPrefix(varScope, "*") {
		return strings.HasSuffix(environment, varScope[1:])
	}
	if strings.HasSuffix(varScope, "*") {
		return strings.HasPrefix(environment, varScope[:len(varScope)-1])
	}
	if strings.Contains(varScope, "*") {
		return strings.HasPrefix(environment, varScope[:strings.Index(varScope, "*")]) && strings.HasSuffix(environment, varScope[strings.Index(varScope, "*")+1:])
	}
	return false
}

// MergeVars merges two lists of variables
// if a variable is present in both lists, the variable from the child list is used
func MergeVars(parentVars []Variable, childVars []Variable) []Variable {
	var res []Variable
	res = make([]Variable, len(parentVars))
	copy(res, parentVars)
	for i := range childVars {
		idx := GetIndexOfVar(res, childVars[i].Key)
		if idx != -1 {
			res = RemoveVar(res, idx)
			res = append(res, childVars[i])
		} else {
			res = append(res, childVars[i])
		}
	}
	return res
}

// FilterVars filters the variables based on the scope
func FilterVars(vars Variables, scope string) Variables {
	var res []Variable
	for i := range vars {
		if IsVarPartOfScope(scope, vars[i].EnvironmentScope) {
			res = append(res, vars[i])
		}
	}
	return res
}

// ExpandAndPrintVars expands the variables and prints them to stdout
func ExpandAndPrintVars(vars Variables) {
	for i := range vars {
		if vars[i].Raw {
			os.Setenv(vars[i].Key, vars[i].Value)
			fmt.Printf("export %s=\"%s\"\n", vars[i].Key, EscapeBashCharacters(vars[i].Value))
		}
	}
	for i := range vars {
		if !vars[i].Raw {
			expandedVar := ExpandEnv(vars[i].Value)
			fmt.Printf("export %s=\"%s\"\n", vars[i].Key, expandedVar)
			os.Setenv(vars[i].Key, vars[i].Value)
		}
	}
}

// ExpandEnv replaces ${var} or $var in the string based on the values of the
// current environment variables. The replacement is case-sensitive. References
// to undefined variables are replaced by the empty string. A default value can
// be given by using the form ${var:-default value}. The default value is used
// only if var is unset or empty. A different value can be given by using the
// form ${var:default value}. The default value is used if var is unset.
// References to other variables are expanded as the string is processed.
// Recursive references are not allowed. If there is an error in the syntax of
// the variable reference, the reference is replaced by the empty string.
func ExpandEnv(s string) string {
	return os.Expand(s, func(v string) string {
		return os.Getenv(v)
	})
}
