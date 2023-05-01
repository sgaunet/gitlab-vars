package gitlabapi

import (
	"encoding/json"
	"fmt"
)

func GetGroup(groupId int) (*GitlabGroup, error) {
	var g GitlabGroup
	uri := fmt.Sprintf("groups/%d", groupId)
	gc := NewGitlapApiClient()
	body, err := gc.Request(uri)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &g); err != nil {
		return nil, err
	}
	return &g, nil
}

func (g *GitlabGroup) GetParentId() int {
	return g.ParentId
}

func (g *GitlabGroup) IsRootGroup() bool {
	return g.ParentId == 0
}

func (g *GitlabGroup) GetAllGroupParentId() ([]int, error) {
	var gids []int
	if g.IsRootGroup() {
		return gids, nil
	}
	g2, err := GetGroup(g.GetParentId())
	if err != nil {
		return nil, err
	}
	gids = append(gids, g.Id)
	parentGids, err := g2.GetAllGroupParentId()
	if err != nil {
		return nil, err
	}
	gids = append(gids, parentGids...)
	return gids, nil
}

func (g *GitlabGroup) GetVarsOfGroup(scope string) (Variables, error) {
	var v, vResult Variables
	uri := fmt.Sprintf("groups/%d/variables", g.Id)
	gc := NewGitlapApiClient()
	body, err := gc.Request(uri)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &v); err != nil {
		return nil, err
	}
	for _, varRange := range v {
		if IsVarPartOfScope(scope, varRange.EnvironmentScope) {
			vResult = append(vResult, varRange)
		}
	}
	return vResult, nil
}

func (g *GitlabGroup) GetAllVars(scope string) ([]Variable, error) {
	var v []Variable
	parents, err := GetAllGroupParentId(g.Id)
	if err != nil {
		return nil, err
	}
	for p := range parents {
		gTmp, err := GetGroup(parents[p])
		if err != nil {
			return nil, err
		}
		vTmp, err := gTmp.GetVarsOfGroup(scope)
		if err != nil {
			return nil, err
		}
		v = MergeVars(v, vTmp)
	}
	return v, nil
}
