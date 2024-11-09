package http_swipe

import (
	"github.com/gin-gonic/gin"
	http_controllers "github.com/kbiits/dealls-take-home-test/adapters/http"
	"github.com/kbiits/dealls-take-home-test/domain/entity"
	swipe_usecase "github.com/kbiits/dealls-take-home-test/usecases/swipe"
	"github.com/rs/zerolog/log"
)

var (
	logger = log.With().Str("module", "http_swipe_controller").Caller().Logger()
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
		logger.Error().Err(err).Msg("failed to get next profile to swipe")
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
		logger.Error().Err(err).Msg("failed to clear blooms by date")
		http_controllers.HandleError(c, err)
		return
	}

	c.JSON(200, http_controllers.NewSuccessResponse("blooms cleared"))
}

func (p *swipeController) SwipeProfile(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(SwipeProfileReq)
	if err := c.ShouldBindJSON(req); err != nil {
		http_controllers.HandleError(c, err)
		return
	}

	swipeDirection := entity.SwipeDirectionLeft
	if req.IsLiked {
		swipeDirection = entity.SwipeDirectionRight
	}

	err := p.swipeUsecase.SwipeProfile(ctx, swipe_usecase.SwipeProfileSpec{
		ProfileID: req.ProfileID,
		Direction: swipeDirection,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to swipe profile")
		http_controllers.HandleError(c, err)
		return
	}

	c.JSON(200, http_controllers.NewSuccessResponse("profile swiped"))
}
