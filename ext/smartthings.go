package ext

import (
	"fmt"
	"regexp"
)

// Shard is a multi-part subvalue of Frigga's Stack name. We avoid using the
// entirety of the Frigga stack in favor of a superset that is more terse.
type Shard struct {
	Name        string
	Environment string
	Region      string
}

const shardPattern = "^([a-z]{2}[0-9]{2}|global)([a-z]{0,1})_([a-z0-9]+)$"

var (
	shardMatcher = regexp.MustCompile(shardPattern)
	envMap       = map[string]string{
		"":  "production",
		"s": "staging",
		"q": "qa",
		"t": "loadtest",
		"d": "dev",
	}
)

// ParseStack will attempt to extract values from a Frigga Stack value and
// output a Shard struct.
func ParseStack(input string) (*Shard, error) {
	if !shardMatcher.MatchString(input) {
		return nil, fmt.Errorf("Stack '%s' is not a valid shard format", input)
	}
	m := shardMatcher.FindStringSubmatch(input)

	shard := &Shard{
		Name:   fmt.Sprintf("%s%s", m[1], m[2]),
		Region: m[3],
	}

	for k, env := range envMap {
		if k == m[2] {
			shard.Environment = env
			break
		}
	}

	if shard.Environment == "" {
		return nil, fmt.Errorf("Could not determine environment from stack %s", input)
	}
	return shard, nil
}
