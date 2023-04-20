package dto

import (
	"github.com/gin-gonic/gin"
)

type GenericResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   gin.H  `json:"data"`
}

func NewGenericResponseBuilder() *GenericResponse {
	return &GenericResponse{}
}

func (r *GenericResponse) SetStatus(Status int) *GenericResponse {
	r.Status = Status
	return r
}

func (r *GenericResponse) SetMessage(message string) *GenericResponse {
	r.Msg = message
	return r
}

func (r *GenericResponse) SetData(Data gin.H) *GenericResponse {
	r.Data = Data
	return r
}

func (r *GenericResponse) MakeResponse(ctx *gin.Context) {
	ctx.JSON(r.Status, r)
}
