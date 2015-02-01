package fhp

import (
	"net/url"
	"strconv"
)

func (fhpApi *FhpApi) GetPhoto(photo_id int, values url.Values) (photoResp PhotoResp, err error) {
	response_ch := make(chan response)

	fhpApi.queryQueue <- query{BaseUrl + "/v1/photos/" + strconv.Itoa(photo_id),
		values,
		&photoResp,
		HTTP_GET,
		response_ch}

	return photoResp, (<-response_ch).err

}
