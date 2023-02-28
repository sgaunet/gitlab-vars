package gitlabapi

import (
	"fmt"
	"regexp"
	"strings"
)

type Variables []Variable

func (v Variables) GetVar(key string) (Variable, error) {
	for i := range v {
		if v[i].Key == key {
			return v[i], nil
		}
	}
	return Variable{}, fmt.Errorf("Variable %s not found", key)
}

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

func (v Variables) GetVarValue(key string) (string, error) {
	for i := range v {
		if v[i].Key == key {
			return v[i].Value, nil
		}
	}
	return "", fmt.Errorf("Variable %s not found", key)
}

// func GetAllVarsOfGroup(groupId int, scope string) ([]Variable, error) {
// 	g, err := GetGroup(groupId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if g.IsRootGroup() {
// 		return GetVarsOfGroup(groupId, scope)
// 	}

// 	var v []Variable
// 	v, err = GetAllVarsOfGroup(g.GetParentId(), scope)
// 	if err != nil {
// 		return nil, err
// 	}
// 	vtmp, err := GetVarsOfGroup(groupId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	v = append(v, vtmp...)
// 	return v, nil
// }

// for v := range vars {
// 	// fmt.Printf("%s=\"%s\"\n", vars[v].Key, EscapeBashCharacters(vars[v].Value))
// 	fmt.Printf("%s//%s//%v\n", vars[v].Key, vars[v].VariableType, vars[v].Raw)
// }

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

// func (v Variables) DeleteVarsOutOfScope(scope string) (res Variables) {
// 	for i := range v {
// 		if IsVarPartOfScope(v[i].EnvironmentScope, scope) {
// 			res = append(res, v[i])
// 		}
// 	}
// 	return res
// }

// func IsVarPartOfScope(environment string, varScope string) bool {
// 	return MatchPattern(varScope, environment)
// }

// MatchPattern checks whether a string matches a pattern
// func MatchPattern(pattern, s string) bool {
// 	if pattern == "*" {
// 		return true
// 	}
// 	matched, _ := filepath.Match(pattern, s)
// 	return matched
// }

// toto is a function that retruns a boolean
// and is checking if the environeemnt is part of the scope
// for example if the environment is "staging" and the scope is "staging" it returns true
// if the environment is "staging" and the scope is "staging*" it returns true
// if the environment is "staging/mtrg" and the scope is "staging*" it returns true
// if the environment is "staging/mtrg" and the scope is "production" it returns false
// if the environment is "staging/mtrg" and the scope is "production*" it returns false
// if the environment is "staging/mtrg" and the scope is "staging/mtrg" it returns true
// if the environment is "staging/mtrg" and the scope is "staging/mtrg*" it returns true
// if the environment is "staging/mtrg" and the scope is "staging/mtrg/*" it returns true
// if the environment is "staging/mtrg" and the scope is "staging/mtrg/*/*" it returns true
// if the environment is "staging/mtrg" and the scope is "staging/mtrg/*/*/*" it returns true
// if the environment is "staging/mtrg" and the scope is "staging/mtrg/*/*/*/*" it returns true
// write the code copilot
func IsVarPartOfScope(environment string, varScope string) bool {
	if varScope == "*" {
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
