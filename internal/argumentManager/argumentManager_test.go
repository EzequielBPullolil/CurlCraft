package argumentmanager

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestXmlFormat(t *testing.T) {
	data := "{\"nombre\":\"ezequiel\", \"password\":\"123123asd\"}"
	esperado := "<data><nombre>ezequiel</nombre><password>123123asd</password></data>"
	resultado, _ := ioutil.ReadAll(xml(data))
	if string(resultado) != esperado {
		fmt.Println(esperado)
		t.Errorf("Resultado no esperado %s", resultado)
	}
}

func TestSimpleKeyValue(t *testing.T) {
	data := "{\"nombre\":\"ezequiel\", \"password\":\"123123asd\"}"
	esperado := "nombre=ezequiel&password=123123asd"
	resultado, _ := ioutil.ReadAll(simpleKeyValue(data))
	if string(resultado) != esperado {
		fmt.Println(esperado)
		t.Errorf("Resultado no esperado %s", resultado)
	}
}
