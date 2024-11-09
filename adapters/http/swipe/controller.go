package http_swipe

import (
	"github.com/gin-gonic/gin"
	http_controllers "github.com/kbiits/dealls-take-home-test/adapters/http"
	swipe_usecase "github.com/kbiits/dealls-take-home-test/usecases/swipe"
	"github.com/rs/zerolog/log"
)

type swipeController struct {
	swipeUsecase swipe_usecase.SwipeUsecase
}

func NewSwipeController(
	swipeUsecase swipe_usecase.SwipeUsecase,
) http_controllers.SwipeController {
	return &swipeController{
		swipeUsecase: swipeUsecase,
	}
}

func (p *swipeController) GetNextProfileToSwipe(c *gin.Context) {
	ctx := c.Request.Context()

	nextProfile, err := p.swipeUsecase.GetNextProfileToSwipe(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get next profile to swipe")
		http_controllers.HandleError(c, err)
		return
	}

	c.JSON(200, http_controllers.NewSuccessResponse(nextProfile))
}

func (p *swipeController) ClearBloomsByDate(c *gin.Context) {
	ctx := c.Request.Context()

	type request struct {
		Date string `json:"date" validate:"required"`
	}

	req := new(request)
	if err := c.ShouldBindJSON(req); err != nil {
		http_controllers.HandleError(c, err)
		return
	}

	err := p.swipeUsecase.ClearBloomsByDate(ctx, req.Date)
	if err != nil {
		log.Error().Err(err).Msg("failed to clear blooms by date")
		http_controllers.HandleError(c, err)
		return
	}

	c.JSON(200, http_controllers.NewSuccessResponse("blooms cleared"))
}
