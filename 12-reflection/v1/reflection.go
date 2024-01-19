// Pacote refletction com a função percorre que recebe uma struct x e chama fn para todos os campos string encontrados dentro dela.
package reflection

import "reflect"


func percorre(x interface{}, fn func(entrada string)) {
	valor := reflect.ValueOf(x) // Pegando o valor de x
	campo := valor.Field(0) // No campo 0
	fn(campo.String())
}