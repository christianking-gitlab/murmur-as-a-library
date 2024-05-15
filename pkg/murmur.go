package pkg

import (
	"github.com/busser/murmur/pkg/internal/murmur"
	"github.com/busser/murmur/pkg/internal/cmd"
)

func ResolveAll(vars map[string]string) (map[string]string, error) {
	return murmur.ResolveAll(vars)
}

func PublicMain() {
	cmd.Execute()
}
