package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/domain/localerror"
	"github.com/valdinei-santos/product-details/modules/product/dto"
	"github.com/valdinei-santos/product-details/modules/product/usecases/compare"
	"github.com/valdinei-santos/product-details/modules/product/usecases/create"
	"github.com/valdinei-santos/product-details/modules/product/usecases/delete"
	"github.com/valdinei-santos/product-details/modules/product/usecases/get"
	getall "github.com/valdinei-santos/product-details/modules/product/usecases/get-all"
	"github.com/valdinei-santos/product-details/modules/product/usecases/update"
)

// Create - Controlador para criar um produto
func Create(log logger.ILogger, ctx *gin.Context, useCase create.IUsecase) {
	log.Debug("Entrou controller.Get")
	var input *dto.Request
	err := json.NewDecoder(ctx.Request.Body).Decode(&input)
	if err != nil {
		outputError(log, ctx, err, "Create/json.NewDecoder")
		return
	}
	resp, err := useCase.Execute(input)
	if err != nil {
		if err == localerror.ErrProductInternal {
			outputErrorInternal(log, ctx, "Create/useCase.Execute")
			return
		}
		outputError(log, ctx, err, "Create/useCase.Execute")
		return
	}

	/* // Retorna a resposta padrão
	result := &dto.OutputDefault{
		StatusCode: 1,
		Message:    "Produto inserido com sucesso",
	} */

	ctx.JSON(http.StatusCreated, resp)
	log.Info("### Finished OK", "status_code", http.StatusCreated)
}

// Delete - Controlador para deletar um produto
func Delete(log logger.ILogger, ctx *gin.Context, useCase delete.IUsecase) {
	log.Debug("Entrou controller.Delete")
	id, err := getIdParam(log, ctx)
	if err != nil {
		return
	}
	log.Debug("ID: " + id)
	err = useCase.Execute(id)
	if err != nil {
		if err == localerror.ErrProductInternal {
			outputErrorInternal(log, ctx, "Delete/useCase.Execute")
			return
		}
		outputError(log, ctx, err, "Delete/useCase.Execute")
		return
	}

	// Retorna a resposta padrão
	result := &dto.OutputDefault{
		StatusCode: 1,
		Message:    "Produto deletado com sucesso",
	}

	ctx.JSON(http.StatusOK, result)
	log.Info("### Finished OK", "status_code", http.StatusOK)
}

// Get - Controlador para obter um produto por ID
func Get(log logger.ILogger, ctx *gin.Context, useCase get.IUsecase) {
	log.Debug("Entrou controller.Get")
	id, err := getIdParam(log, ctx)
	if err != nil {
		outputError(log, ctx, err, "Get/getIdParam")
		return
	}
	log.Debug("ID: " + id)
	resp, err := useCase.Execute(id)
	if err != nil {
		if err == localerror.ErrProductInternal {
			outputErrorInternal(log, ctx, "Get/useCase.Execute")
			return
		}
		outputError(log, ctx, err, "Get/useCase.Execute")
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
}

// GetAll - Controlador para obter todos os produtos
func GetAll(log logger.ILogger, ctx *gin.Context, useCase getall.IUsecase) {
	log.Debug("Entrou controller.GetAll")

	// Pega os parâmetros de paginação (page) da query string
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		outputError(log, ctx, err, "GetAll/parse_page")
		return
	}

	// Pega os parâmetros de paginação (size) da query string
	size, err := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil || size < 1 {
		outputError(log, ctx, err, "GetAll/parse_size")
		return
	}

	resp, err := useCase.Execute(page, size)
	if err != nil {
		if err == localerror.ErrProductInternal {
			outputErrorInternal(log, ctx, "GetAll/useCase.Execute")
			return
		}
		outputError(log, ctx, err, "GetAll/useCase.Execute")
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
}

// Update - Controlador para alterar um produto pelo ID
func Update(log logger.ILogger, ctx *gin.Context, useCase update.IUsecase) {
	log.Debug("Entrou controller.Update")
	id, err := getIdParam(log, ctx)
	if err != nil {
		return
	}
	log.Debug("ID: " + id)
	var input *dto.Request
	err = json.NewDecoder(ctx.Request.Body).Decode(&input)
	if err != nil {
		outputError(log, ctx, err, "Update/json.NewDecoder")
		return
	}
	resp, err := useCase.Execute(id, input)
	if err != nil {
		if err == localerror.ErrProductInternal {
			outputErrorInternal(log, ctx, "Update/useCase.Execute")
			return
		}
		outputError(log, ctx, err, "Update/useCase.Execute")
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
}

// Compare - Controlador para obter a lista de produtos a comparar por ID
func Compare(log logger.ILogger, ctx *gin.Context, useCase compare.IUsecase) {
	log.Debug("Entrou controller.Compare")

	ids, err := getIdsQueryParams(ctx)
	if err != nil {
		outputError(log, ctx, err, "Compare/getIdsQueryParams")
		return
	}

	resp, err := useCase.Execute(ids)
	if err != nil {
		if err == localerror.ErrProductInternal {
			outputErrorInternal(log, ctx, "Compare/useCase.Execute")
			return
		}
		outputError(log, ctx, err, "Compare/useCase.Execute")
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
}

func getIdParam(log logger.ILogger, ctx *gin.Context) (string, error) {
	idParam := ctx.Param("id")
	if idParam == "" {
		log.Error(localerror.ErrProductNoneID.Error(), "mtd", "getIdParam")
		return "", localerror.ErrProductNoneID
	}
	return idParam, nil
}

func getIdsQueryParams(ctx *gin.Context) ([]string, error) {
	// Pega lista de IDs da query string: /api/products/compare?ids=1,2,3
	idsQueryParam := ctx.Query("ids")
	if idsQueryParam == "" {
		return nil, localerror.ErrProductNoneID
	}
	idStrs := strings.Split(idsQueryParam, ",")
	var ids []string
	ids = append(ids, idStrs...)
	return ids, nil
}

func outputError(log logger.ILogger, ctx *gin.Context, err error, method string) {
	log.Error(err.Error(), "mtd", method)
	dataJErro := dto.OutputDefault{
		StatusCode: -1,
		Message:    err.Error(),
	}
	ctx.JSON(http.StatusBadRequest, dataJErro)
	log.Info("### Finished ERROR", "status_code", http.StatusBadRequest)
}

func outputErrorInternal(log logger.ILogger, ctx *gin.Context, method string) {
	log.Error(localerror.ErrProductInternal.Error(), "mtd", method)
	dataJErro := dto.OutputDefault{
		StatusCode: -1,
		Message:    localerror.ErrProductInternal.Error(),
	}
	ctx.JSON(http.StatusInternalServerError, dataJErro)
	log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
}
