package function

import (
	"io/ioutil"
	"log"
	"os"
)

// Handle a serverless request
func Handle(req []byte) string {
	name := getName()

	key := os.Getenv("debug")
	if key != "" {
		secretFile, secretValue := getSecret(key)

		log.Printf("secretValue: %s [%s]\n%s\n", key, secretFile, secretValue)
	}

	return name
}

func getSecret(key string) (string, string) {
	file := "/var/openfaas/secrets/" + key
	apiKey, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("error reading secret api-key %s\n", err)
	}
	return file, string(apiKey)
}

func getName() string {
	name, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return name
}
