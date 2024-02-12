package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Rio Grande", "Avenida"},
		{"Rodovia dos Abençoados", "Rodovia"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA DOS COGUMÉLOS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("Caso de teste falhou para '%s'. O tipo recebido foi '%s', mas o esperado era '%s'",
				cenario.enderecoInserido,
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}

		// 	enderecoParaTeste := "Avenida Paulista"
		// 	tipoEnderecoEsperado := "Avenida"
		// 	tipoEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

		// 	if tipoEnderecoRecebido != tipoEnderecoEsperado {
		// 		t.Errorf("O tipo recebido é diferente do esperado! Esperava-se %s e foi recebido %s", tipoEnderecoEsperado, tipoEnderecoRecebido)
		// }
	}
}
