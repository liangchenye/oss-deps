package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"gopkg.in/macaron.v1"

	"github.com/liangchenye/oss-deps/service"
)

type httpListRet struct {
	Message string
	Content interface{}
}

//TODO: better http return result
func httpRet(head string, content interface{}, err error) (int, []byte) {
	var ret httpListRet
	var code int

	if err != nil {
		ret.Message = head + " fail"
		ret.Content = err.Error()
		code = http.StatusBadRequest
	} else {
		ret.Message = head
		ret.Content = content
		code = http.StatusOK
	}

	result, _ := json.Marshal(ret)
	return code, result
}

// TrainListHandler lists  all the trains
func TrainListHandler(ctx *macaron.Context) (int, []byte) {
	h, _ := service.GetDefault()
	trains, err := h.GetTrains()

	return httpRet("Train List Trains", trains, err)
}

// TrainPackageHandler gets the list packages of a train
func TrainPackageHandler(ctx *macaron.Context) (int, []byte) {
	name := ctx.Params(":name")
	version := ctx.Params(":version")
	h, _ := service.GetDefault()
	trains, _ := h.GetTrains()
	for _, t := range trains {
		if t.Name == name && t.Version == version {
			pkgs, err := h.GetPackagesFromTrain(t)
			return httpRet("Train Get Packages", pkgs, err)
		}
	}
	//TODO: 404
	return httpRet("Train Get Meta", nil, errors.New("Cannot find the train"))
}

// TrainCreateHandler gets the meta data of all the namespace/repository
func TrainCreateHandler(ctx *macaron.Context) (int, []byte) {
	//	name := ctx.Params(":name")
	return http.StatusOK, nil

	//return httpRet("Train Create a train", nil, err)
}
