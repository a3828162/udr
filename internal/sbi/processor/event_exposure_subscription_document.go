/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package processor

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/openapi/models"
	udr_context "github.com/free5gc/udr/internal/context"
	"github.com/free5gc/udr/internal/util"
)

func (p *Processor) RemoveeeSubscriptionsProcedure(c *gin.Context, ueId string, subsId string) {
	udrSelf := udr_context.GetSelf()
	value, ok := udrSelf.UESubsCollection.Load(ueId)
	if !ok {
		pd := util.ProblemDetailsNotFound("USER_NOT_FOUND")
		c.JSON(int(pd.Status), pd)
		return
	}

	UESubsData := value.(*udr_context.UESubsData)
	_, ok = UESubsData.EeSubscriptionCollection[subsId]

	if !ok {
		pd := util.ProblemDetailsNotFound("SUBSCRIPTION_NOT_FOUND")
		c.JSON(int(pd.Status), pd)
		return
	}
	delete(UESubsData.EeSubscriptionCollection, subsId)
	c.Status(http.StatusNoContent)
}

func (p *Processor) UpdateEesubscriptionsProcedure(c *gin.Context, ueId string, subsId string,
	EeSubscription models.EeSubscription,
) {
	udrSelf := udr_context.GetSelf()
	value, ok := udrSelf.UESubsCollection.Load(ueId)
	if !ok {
		pd := util.ProblemDetailsNotFound("USER_NOT_FOUND")
		c.JSON(int(pd.Status), pd)
		return
	}

	UESubsData := value.(*udr_context.UESubsData)
	_, ok = UESubsData.EeSubscriptionCollection[subsId]

	if !ok {
		pd := util.ProblemDetailsNotFound("SUBSCRIPTION_NOT_FOUND")
		c.JSON(int(pd.Status), pd)
		return
	}
	UESubsData.EeSubscriptionCollection[subsId].EeSubscriptions = &EeSubscription

	c.Status(http.StatusNoContent)
}