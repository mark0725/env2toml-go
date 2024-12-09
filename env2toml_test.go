package env2toml

import (
	"fmt"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
)

func TestEnv(t *testing.T) {
	os.Setenv("APP_TITLE", "TOML Example")
	os.Setenv("APP_OWNER__NAME", "Tom Preston-Werner")
	os.Setenv("APP_DATABASE__ENABLED", "true")
	os.Setenv("APP_DATABASE__PORTS", "[ 8000, 8001, 8002 ]")
	os.Setenv("APP_SERVERS__ALPHA__IP", "10.0.0.1")
	os.Setenv("APP_SERVERS__ALPHA__ROLE", "frontend")
	os.Setenv("APP_SERVERS__BETA__IP", "10.0.0.2")
	os.Setenv("APP_SERVERS__BETA__ROLE", "backend")
	os.Setenv("APP_USERS__0__NAME", "USER0")
	os.Setenv("APP_USERS__0__PASSWORD", "u0")
	os.Setenv("APP_USERS__1__NAME", "USER1")
	os.Setenv("APP_USERS__1__PASSWORD", "u1")

	result, err := Parse("APP_")
	if err != nil {
		t.Fatal("Error:", err)
		return
	}

	fmt.Printf("result:\n%s\n", result)
	var tomlVars map[string]interface{}
	if _, err := toml.Decode(result, &tomlVars); err != nil {
		t.Fatal("Error:", err)
		return
	}

	fmt.Printf("result:%v\n", tomlVars)
}
