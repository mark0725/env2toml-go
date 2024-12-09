package env2toml

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

type VarItem struct {
	Section    *string
	Array      bool
	ArrayIndex int
	Key        string
	Value      string
}

func Parse(prefix string) (string, error) {
	var varList []VarItem
	var sections []string

	// parse Environ
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		k, v := pair[0], pair[1]
		if strings.HasPrefix(k, prefix) {
			s := strings.ToLower(strings.TrimPrefix(k, prefix))
			keys := strings.Split(s, "__")
			var section string
			isArray := false
			arrayIndex := -1

			for i := 0; i < len(keys)-1; i++ {
				if index, err := strconv.Atoi(keys[i]); err == nil {
					isArray = true
					arrayIndex = index
					continue
				}

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
			varList = append(varList, VarItem{
				Section:    secPtr,
				Array:      isArray,
				ArrayIndex: arrayIndex,
				Key:        newKey,
				Value:      value,
			})
		}
	}

	var result strings.Builder

	// is a section
	for _, item := range varList {
		if item.Section == nil {
			result.WriteString(fmt.Sprintf("%s=%s\n", item.Key, item.Value))
		}
	}

	// Generate TOML text for sections
	for _, sec := range sections {
		arrayItems := make(map[int][]VarItem)
		isArray := false

		for _, item := range varList {
			if item.Section != nil && *item.Section == sec {
				if item.Array {
					isArray = true
					arrayItems[item.ArrayIndex] = append(arrayItems[item.ArrayIndex], item)
				} else {
					arrayItems[0] = append(arrayItems[0], item)
				}
			}
		}

		if !isArray {
			result.WriteString(fmt.Sprintf("\n[%s]\n", sec))
			for _, item := range arrayItems[0] {
				result.WriteString(fmt.Sprintf("%s=%s\n", item.Key, item.Value))
			}

			continue
		}

		keys := make([]int, 0, len(arrayItems))
		for key := range arrayItems {
			keys = append(keys, key)
		}
		sort.Ints(keys)
		for _, k := range keys {
			items := arrayItems[k]
			result.WriteString(fmt.Sprintf("\n[[%s]]\n", sec))
			for _, item := range items {
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
