package env2toml

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
)

func Parse(prefix string) (string, error) {
	var varList []struct {
		Section *string
		Key     string
		Value   string
	}
	var sections []string

	// parse Environ
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		k, v := pair[0], pair[1]
		if strings.HasPrefix(k, prefix) {
			s := strings.ToLower(strings.TrimPrefix(k, prefix))
			keys := strings.Split(s, "__")
			var section string

			for i := 0; i < len(keys)-1; i++ {
				if section != "" {
					section = section + "." + keys[i]
				} else {
					section = keys[i]
				}

				if !contains(sections, section) {
					sections = append(sections, section)
				}
			}

			newKey := keys[len(keys)-1]
			value := v

			// is a toml value?
			if !isValidTomlValue(`a=` + v) {
				value = fmt.Sprintf(`"%s"`, strings.ReplaceAll(v, `\`, `\\`))
			}

			var secPtr *string
			if section != "" {
				secPtr = &section
			}
			varList = append(varList, struct {
				Section *string
				Key     string
				Value   string
			}{secPtr, newKey, value})
		}
	}

	var result strings.Builder

	// is a section
	for _, item := range varList {
		if item.Section == nil {
			result.WriteString(fmt.Sprintf("%s=%s\n", item.Key, item.Value))
		}
	}

	// gen toml text
	for _, sec := range sections {
		result.WriteString(fmt.Sprintf("\n[%s]\n", sec))
		for _, item := range varList {
			if item.Section != nil && *item.Section == sec {
				result.WriteString(fmt.Sprintf("%s=%s\n", item.Key, item.Value))
			}
		}
	}

	return result.String(), nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func isValidTomlValue(value string) bool {
	// Github.com/BurntSushi/toml test toml value.
	var result map[string]interface{}
	_, err := toml.Decode(value, &result)
	return err == nil
}
