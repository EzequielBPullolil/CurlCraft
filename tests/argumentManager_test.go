package test

import (
	"fmt"
	"io"
	"testing"

	argumentManager "github.com/EzequielK-source/CurlCraft/internal/argumentManager"
)

func TestXmlFormat(t *testing.T) {
	data := "{\"nombre\":\"ezequiel\", \"password\":\"123123asd\"}"
	esperado := "<data><nombre>ezequiel</nombre><password>123123asd</password></data>"
	resultado, _ := io.ReadAll(argumentManager.ContentTypeEncoder("XML", data))
	if string(resultado) != esperado {
		fmt.Println(esperado)
		t.Errorf("Resultado no esperado %s", resultado)
	}
}

func TestSimpleKeyValue(t *testing.T) {
	data := "{\"nombre\":\"ezequiel\", \"password\":\"123123asd\"}"
	esperado := "nombre=ezequiel&password=123123asd"
	for _, v := range []string{"FORM", "PLAIN", "HTML"} {
		resultado, _ := io.ReadAll(argumentManager.ContentTypeEncoder(v, data))
		if string(resultado) != esperado {
			fmt.Println(esperado)
			t.Errorf("Resultado no esperado %s", resultado)
		}
	}
}
