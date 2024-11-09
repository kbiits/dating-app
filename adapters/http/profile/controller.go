package http_profile

import (
	"github.com/gin-gonic/gin"
	http_controllers "github.com/kbiits/dealls-take-home-test/adapters/http"
	profile_usecase "github.com/kbiits/dealls-take-home-test/usecases/profile"
	validator_util "github.com/kbiits/dealls-take-home-test/utils/validator"
	"github.com/rs/zerolog/log"
	"github.com/samber/mo"
)

type profileController struct {
	profileUsecase profile_usecase.ProfileUsecase
}

func NewProfileController(
	profileUsecase profile_usecase.ProfileUsecase,
) http_controllers.ProfileController {
	return &profileController{
		profileUsecase: profileUsecase,
	}
}

func (p *profileController) GetLoggedInProfile(c *gin.Context) {
	ctx := c.Request.Context()

	profile, err := p.profileUsecase.GetLoggedInProfile(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get logged in profile")
		http_controllers.HandleError(c, err)
		return
	}

	c.JSON(200, http_controllers.NewSuccessResponse(profile))
}

func (p *profileController) CompleteProfile(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(CompleteProfileReq)
	if err := c.BindJSON(req); err != nil {
		log.Error().Err(err).Msg("failed to bind request")
		http_controllers.HandleError(c, err)
		return
	}

	if err := validator_util.GetValidator().Struct(req); err != nil {
		log.Error().Err(err).Msg("validation error")
		http_controllers.HandleError(c, err)
		return
	}

	result, err := p.profileUsecase.CompleteProfile(ctx, profile_usecase.CompleteProfileSpec{
		DisplayName:   req.DisplayName,
		Bio:           mo.PointerToOption(req.Bio),
		Gender:        mo.PointerToOption(req.Gender),
		Dob:           mo.PointerToOption(req.Dob),
		ProfilePicURL: mo.PointerToOption(req.ProfilePicURL),
		DistrictID:    mo.PointerToOption(req.DistrictID),
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to complete profile")
		http_controllers.HandleError(c, err)
		return
	}

	c.JSON(200, http_controllers.NewSuccessResponse(result))
}
