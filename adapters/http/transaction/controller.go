package http_transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
	http_controllers "github.com/kbiits/dealls-take-home-test/adapters/http"
	transaction_usecase "github.com/kbiits/dealls-take-home-test/usecases/transaction"
	validator_util "github.com/kbiits/dealls-take-home-test/utils/validator"
)

type transactionController struct {
	transactionUsecase transaction_usecase.TransactionUsecase
}

func NewTransactionController(
	transactionUsecase transaction_usecase.TransactionUsecase,
) http_controllers.TransactionController {
	return &transactionController{
		transactionUsecase: transactionUsecase,
	}
}

func (t *transactionController) Buy(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(BuyRequest)

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator_util.GetValidator().Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := t.transactionUsecase.Buy(ctx, req.PackageID)
	if err != nil {
		http_controllers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, http_controllers.NewSuccessResponse("Transaction success"))
}
