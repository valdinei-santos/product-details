package localerror

import "errors"

var ErrProductNotFound = errors.New("produto não encontrado")
var ErrProductNotFoundMany = errors.New("nenhum produto encontrado")
var ErrProductNotNil = errors.New("produto não pode ser nil")
var ErrProductIDInvalid = errors.New("ID inválido")
var ErrProductUUIDInvalid = errors.New("UUID inválido")
var ErrProductValidation = errors.New("erro de validação")
var ErrProductNoneID = errors.New("nenhum ID informado")
var ErrProductNameInvalid = errors.New("o nome deve ter entre 3 e 50 caracteres")
var ErrProductUrlImageInvalid = errors.New("a url da imagem deve ter entre 10 e 300 caracteres")
var ErrProductPriceNegative = errors.New("o valor do preço não pode ser negativo")
var ErrProductClassificationInvalid = errors.New("a classificacao deve ter entre 3 e 50 caracteres")
var ErrProductDescriptionInvalid = errors.New("a descricao deve ter entre 5 e 100 caracteres")
var ErrProductSpecificationInvalid = errors.New("a especificacao deve ter entre 10 e 200 caracteres")
var ErrProductSaveInDatabase = errors.New("erro ao salvar no banco de dados")
var ErrProductConnectionInDatabase = errors.New("erro de conexão com o banco de dados")
var ErrProductInternal = errors.New("erro interno no servidor")
