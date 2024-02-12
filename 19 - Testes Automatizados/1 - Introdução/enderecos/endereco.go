package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoLetrasMinusculas := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoLetrasMinusculas, " ")[0]

	enderecoTemTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo Inválido"

}
