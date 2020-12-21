package query

import (
	"regexp"
)

type Actions interface {
	Query(text string) []Action
}

func (a actions) Query(text string) []Action {
	rActions := []Action{}
	if len(a.actions) == 0 {
		return rActions
	}

	for _, action := range a.actions {
		re := regexp.MustCompile(action.regex)
		if re.MatchString(text) {
			rActions = append(rActions, action)
		}

	}
	return rActions

}

type actions struct {
	actions []Action
}

type Action struct {
	regex   string
	replace string
}
