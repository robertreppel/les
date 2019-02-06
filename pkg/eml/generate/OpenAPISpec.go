package generate

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/robertreppel/les/pkg/eml"
	"github.com/robertreppel/les/pkg/eml/generate/openapi"
)

// OpenAPISpec - Event Modeling Language Language to Open API 3.0 (a.k.a. Swagger)
func OpenAPISpec(system eml.Solution) string {
	swaggerYaml := openapi.SwaggerYML(system)
	return swaggerYaml
}

func writeOpenAPISpecHeader(renderingDirectory string, code string, swaggerJSONFileName string) {
	ensureDirectoryExists(renderingDirectory)

	err := ioutil.WriteFile(filepath.Join(renderingDirectory, swaggerJSONFileName), []byte(code), 0644)
	if err != nil {
		log.Fatalf("writeOpenAPISpecHeader: %v", err)
	}

}
