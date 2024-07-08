package env2toml

import (
	"fmt"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	os.Setenv("APP_TITLE", "TOML Example")
	os.Setenv("APP_OWNER__NAME", "Tom Preston-Werner")
	os.Setenv("APP_DATABASE__ENABLED", "true")
	os.Setenv("APP_DATABASE__PORTS", "[ 8000, 8001, 8002 ]")
	os.Setenv("APP_SERVERS__ALPHA__IP", "10.0.0.1")
	os.Setenv("APP_SERVERS__ALPHA__ROLE", "frontend")
	os.Setenv("APP_SERVERS__BETA__IP", "10.0.0.2")
	os.Setenv("APP_SERVERS__BETA__ROLE", "backend")

	result, err := Parse("APP_")
	if err != nil {
		//t.Errorln("Error:", err)
		fmt.Println("Err:", err)
		return
	}

	fmt.Println("result:", result)
}
