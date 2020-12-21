package action

import (
	"github.com/Nhanderu/brdoc"
	"strings"
)

func All() List {
	return list{
		actions: []Item{
			CpfToNum(),
		},
	}
}

func CpfToNum() Item {
	return Item{
		Name: "CPF to número",
		MatchRegex: `^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`,
		Validate: func(s string) bool {
			return true
		},
		Replace: func(cpf string) string {
			s := strings.ReplaceAll(cpf, ".", "")
			return strings.ReplaceAll(s, "-", "")
		},
		MinSize: 14,
		MaxSize: 14,
	}

}

func NumToCpf() Item {
	return Item{
		Name: "Número to CPF",
		MaxSize: 11,
		MinSize: 9,

		Validate: func(s string) bool {
			return brdoc.IsCPF(s)
		},
		Replace: func(cpf string) string {
			return cpf
		},
	}

}
