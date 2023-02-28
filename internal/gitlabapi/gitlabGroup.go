package gitlabapi

import (
	"encoding/json"
	"fmt"
)

func GetGroup(groupId int) (*GitlabGroup, error) {
	var g GitlabGroup
	uri := fmt.Sprintf("groups/%d", groupId)
	_, body, _ := Request(uri)
	if err := json.Unmarshal(body, &g); err != nil {
		return nil, err
	}
	// fmt.Println(g)
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
	_, body, _ := Request(uri)
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
