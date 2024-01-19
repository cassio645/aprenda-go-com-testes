// Pacote refletction com a função percorre que recebe uma struct x e chama fn para todos os campos string encontrados dentro dela.
package reflection

import "reflect"

func percorre(x interface{}, fn func(entrada string)) {
	valor := reflect.ValueOf(x) // Pegando valor da variavel x

	// Adicionando cada um deles na função fn
	for i := 0; i < valor.NumField(); i++ {
		campo := valor.Field(i)

		// Verifica se o valor é do tipo string, antes de passar para a função fn
		// Se forem campos aninhados ele faz uma chamada recursiva
		switch campo.Kind() {
		case reflect.String:
			fn(campo.String())
		case reflect.Struct:
			percorre(campo.Interface(), fn)

		}

	}
}
