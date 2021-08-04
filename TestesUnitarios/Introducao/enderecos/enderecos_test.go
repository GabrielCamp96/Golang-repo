package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Válido"},
		{"Viela X", "Inválido"},
		{"Estrada y", "Válido"},
		{"Aevnida s", "Inválido"},
		{"Avenida Paulista", "Válido"},
		{"Rodovia Imigrantes", "Válido"},
		{"Avenidatudojunto", "Inválido"},
	}

	for _, cenario := range cenariosDeTeste {

		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		tipoDeEnderecoEsperado := cenario.retornoEsperado

		if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s para %s e recebeu %s",
				tipoDeEnderecoEsperado,
				cenario.enderecoInserido,
				tipoDeEnderecoRecebido)
		}
	}

}
