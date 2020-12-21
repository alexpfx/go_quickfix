package action

import "regexp"

type Item struct {
	Name       string
	MatchRegex string
	Replace    func(string) string
	MinSize int
	MaxSize int
}

type List interface {
	Query(text string) []Item
}

type list struct {
	actions []Item
}

func (a list) Query(text string) []Item {
	rItems := make([]Item, 0)
	if len(a.actions) == 0 {
		return rItems
	}

	for _, action := range a.actions {
		textLen := len(text)
		if action.MinSize > 0 && textLen < action.MinSize{
			continue
		}
		if action.MaxSize > 0 && textLen > action.MaxSize{
			continue
		}

		re := regexp.MustCompile(action.MatchRegex)
		if re.MatchString(text) {
			rItems = append(rItems, action)
		}

	}
	return rItems

}
