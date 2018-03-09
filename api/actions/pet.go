package actions

import (
	"github.com/PeterHenell/docker-compose-go-elastic-swagger/api/generated/restapi/operations/pet"
	"github.com/go-openapi/runtime/middleware"
	"github.com/PeterHenell/docker-compose-go-elastic-swagger/api/generated/models"
	"github.com/PeterHenell/docker-compose-go-elastic-swagger/api/generated/restapi/operations"
)

var AddPetHandler pet.AddPetHandlerFunc = func(params pet.AddPetParams, principal interface{}) middleware.Responder {
	// oops, the spec did not contain any respones wich were possitive. The only specified responce is 405
	return pet.NewAddPetMethodNotAllowed()
}

var DeletePetHandler pet.DeletePetHandlerFunc = func(params pet.DeletePetParams, principal interface{}) middleware.Responder {
	return operations.
	return middleware.NotImplemented("operation pet.DeletePet has not yet been implemented")
}

var FindPetsByStatusHandler pet.FindPetsByStatusHandlerFunc = func(params pet.FindPetsByStatusParams, principal interface{}) middleware.Responder {
	return middleware.NotImplemented("operation pet.FindPetsByStatus has not yet been implemented")
}

var FindPetsByTagsHandler pet.FindPetsByTagsHandlerFunc = func(params pet.FindPetsByTagsParams, principal interface{}) middleware.Responder {
	return middleware.NotImplemented("operation pet.FindPetsByTags has not yet been implemented")
}

var GetPetByIDHandler pet.GetPetByIDHandlerFunc = func(params pet.GetPetByIDParams, principal interface{}) middleware.Responder {
	return middleware.NotImplemented("operation pet.GetPetByID has not yet been implemented")
}

var UpdatePetHandler pet.UpdatePetHandlerFunc = func(params pet.UpdatePetParams, principal interface{}) middleware.Responder {
	return middleware.NotImplemented("operation pet.UpdatePet has not yet been implemented")
}

var UpdatePetWithFormHandler pet.UpdatePetWithFormHandlerFunc = func(params pet.UpdatePetWithFormParams, principal interface{}) middleware.Responder {
	return middleware.NotImplemented("operation pet.UpdatePetWithForm has not yet been implemented")
}

var UploadFileHandler pet.UploadFileHandlerFunc = func(params pet.UploadFileParams, principal interface{}) middleware.Responder {
	return middleware.NotImplemented("operation pet.UploadFile has not yet been implemented")
}
