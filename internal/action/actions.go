package action

import (
	"strings"
)

func CpfToNum() Item {
	return Item{
		Name:       "CPF to número",
		MatchRegex: `^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`,
		Replace: func(cpf string) string {
			s := strings.ReplaceAll(cpf, ".", "")
			return strings.ReplaceAll(s, "-", "")
		},
		MinSize: 14,
		MaxSize: 14,
	}


}
