package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tsyahputra/luhkumbe/helper"
	"github.com/tsyahputra/luhkumbe/model"
)

func FetchProvinces(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://wilayah.id/api/provinces.json")
	if err != nil {
		helper.ResponseMessage(w, http.StatusNotFound, err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		helper.ResponseMessage(w, http.StatusBadRequest, "API returned a non-200 status code")
		return
	}
	var apiResponse model.ProvinceAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		helper.ResponseMessage(w, http.StatusFailedDependency, err.Error())
		return
	}
	for _, province := range apiResponse.Data {
		if err := model.DB.FirstOrCreate(&province, model.Province{
			Code: province.Code,
		}).Error; err != nil {
			helper.ResponseMessage(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	helper.ResponseMessage(w, http.StatusOK, "Data provinsi berhasil disimpan ke database!")
}

func FetchRegenciesByProvinceId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provinceCode := vars["provincecode"]

	url := fmt.Sprintf("https://wilayah.id/api/regencies/%s.json", provinceCode)
	resp, err := http.Get(url)
	if err != nil {
		helper.ResponseMessage(w, http.StatusNotFound, err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		helper.ResponseMessage(w, http.StatusBadRequest, "API returned a non-200 status code")
		return
	}
	var apiResponse model.RegencyAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		helper.ResponseMessage(w, http.StatusFailedDependency, err.Error())
		return
	}

	for _, regency := range apiResponse.Data {
		regency.ProvinceID = provinceCode
		if err := model.DB.FirstOrCreate(&regency, model.Regency{
			Code: regency.Code,
		}).Error; err != nil {
			helper.ResponseMessage(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	helper.ResponseMessage(w, http.StatusOK, "Data kabupaten berhasil disimpan ke database!")
}

func FetchDistrictsByRegencyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	regencyCode := vars["regencycode"]
	url := fmt.Sprintf("https://wilayah.id/api/districts/%s.json", regencyCode)
	resp, err := http.Get(url)
	if err != nil {
		helper.ResponseMessage(w, http.StatusNotFound, err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		helper.ResponseMessage(w, http.StatusBadRequest, "API returned a non-200 status code")
		return
	}
	var apiResponse model.DistrictAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		helper.ResponseMessage(w, http.StatusFailedDependency, err.Error())
		return
	}
	for _, district := range apiResponse.Data {
		district.RegencyID = regencyCode
		if err := model.DB.FirstOrCreate(&district, model.District{
			Code: district.Code,
		}).Error; err != nil {
			helper.ResponseMessage(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	helper.ResponseMessage(w, http.StatusOK, "Data kecamatan berhasil disimpan ke database!")
}

func FetchVillagesByDistrictId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	districtCode := vars["districtcode"]
	url := fmt.Sprintf("https://wilayah.id/api/villages/%s.json", districtCode)
	resp, err := http.Get(url)
	if err != nil {
		helper.ResponseMessage(w, http.StatusNotFound, err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		helper.ResponseMessage(w, http.StatusBadRequest, "API returned a non-200 status code")
		return
	}
	var apiResponse model.VillageAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		helper.ResponseMessage(w, http.StatusFailedDependency, err.Error())
		return
	}
	for _, village := range apiResponse.Data {
		village.DistrictID = districtCode
		if err := model.DB.FirstOrCreate(&village, model.Village{
			Code: village.Code,
		}).Error; err != nil {
			helper.ResponseMessage(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	helper.ResponseMessage(w, http.StatusOK, "Data kelurahan berhasil disimpan ke database!")
}
