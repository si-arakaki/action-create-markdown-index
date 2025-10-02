package env

import "os"

func Home() string {
	v := os.Getenv("HOME")
	if v == "" {
		panic("env.Home\n\tno value set")
	}

	return v
}
