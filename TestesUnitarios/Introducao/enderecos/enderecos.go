package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	primeiraPalavra := strings.ToLower(strings.Split(endereco, " ")[0])

	for _, value := range tiposValidos {
		if value == primeiraPalavra {
			return "Válido"
		}
	}

	return "Inválido"
}
