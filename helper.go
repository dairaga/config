package config

import (
	"flag"
	"os"
	"strings"

	"github.com/spf13/cast"
)

func SearchMap(obj map[string]any, path []string) any {
	if len(path) == 0 {
		return obj
	}

	next, ok := obj[path[0]]
	if ok {
		if len(path) == 1 {
			return next
		}

		switch v := next.(type) {
		case map[string]any:
			return SearchMap(v, path[1:])
		case map[any]any:
			return SearchMap(cast.ToStringMap(v), path[1:])
		default:
			return nil
		}

	}
	return nil
}

// -----------------------------------------------------------------------------

func Get(obj map[string]any, path string, delim string) any {
	path = strings.ToLower(path)
	return SearchMap(obj, strings.Split(path, delim))
}

// -----------------------------------------------------------------------------

func SetMap(obj map[string]any, path []string, value any) {
	if len(path) == 0 {
		return
	}

	if len(path) == 1 {
		obj[path[0]] = value
		return
	}

	next, ok := obj[path[0]]
	if !ok {
		next = make(map[string]any)
		obj[path[0]] = next
	}

	switch v := next.(type) {
	case map[string]any:
		SetMap(v, path[1:], value)
	}
}

// -----------------------------------------------------------------------------

func OpenCandidate(candidates []string) ([]byte, error) {
	for i := range candidates {
		fi, err := os.Stat(candidates[i])
		if err != nil {
			continue
		}

		if fi.IsDir() {
			continue
		}

		return os.ReadFile(candidates[i])
	}

	return nil, ErrNoCandidates
}

// -----------------------------------------------------------------------------

func BindEnv(obj map[string]any) {
	envs := os.Environ()

	for i := range envs {
		if pair := strings.SplitN(envs[i], "=", 2); len(pair) == 2 {
			key := strings.ToLower(pair[0])
			SetMap(obj, strings.Split(key, ENV_DELIM), pair[1])
		}
	}
}

// -----------------------------------------------------------------------------

func BindFlag(obj map[string]any) {
	if !flag.Parsed() {
		flag.Parse()
	}

	flag.VisitAll(func(f *flag.Flag) {
		key := strings.ToLower(f.Name)
		SetMap(obj, strings.Split(key, FLAG_DELIM), f.Value.String())
	})
}
