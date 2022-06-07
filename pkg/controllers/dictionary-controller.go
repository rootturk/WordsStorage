package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/rootturk/words-storage/pkg/utils"
	"github.com/rootturk/words-storage/pkg/models"
)


var NewDictionary models.Dictionary

func GetDictionaries(w http.ResponseWriter, r *http.Request){
	newDictionaries := models.GetAllDictionaries()
	
	res, _ := json.Marshal(newDictionaries)
	
	w.Header().Set("Content-Type","pkglcation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDictionaryById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	dictionaryId := vars["dictionaryId"]
	ID, err:= strconv.ParseInt(dictionaryId, 0, 0)

	if err != nil{
		fmt.Println("error while parsing")
	}

	dictionaryDetails, _ := models.GetDictionaryById(ID)
	res, _ := json.Marshal(dictionaryDetails)
	
	w.Header().Set("Content-Type", "pkglcation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateDictionary(w http.ResponseWriter, r *http.Request){
	createDictionary := &models.Dictionary{}
	utils.ParseBody(r, createDictionary)
	d :=  createDictionary.CreateDictionary()
	res, _ := json.Marshal(d)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteDictionary(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	dictionaryId := vars["dictionaryId"]
	ID, err := strconv.ParseInt(dictionaryId, 0,0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	dictionary := models.DeleteDictionary(ID)
	res, _ := json.Marshal(dictionary)
	w.Header().Set("Content-Type", "pkglcation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateDictionary(w http.ResponseWriter, r *http.Request){
	var UpdateDictionary = &models.Dictionary{}
	utils.ParseBody(r, UpdateDictionary)

	vars := mux.Vars(r)
	dictionaryId := vars["dictionaryId"]
	ID, err := strconv.ParseInt(dictionaryId, 0,0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	dictionaryDetails, db:= models.GetDictionaryById(ID)

	if UpdateDictionary.LanguageId > 0{
		dictionaryDetails.LanguageId = UpdateDictionary.LanguageId
	}

	if UpdateDictionary.TranslateMean != ""{
		dictionaryDetails.TranslateMean = UpdateDictionary.TranslateMean
	}


	if UpdateDictionary.LongMeanDescription != ""{
		dictionaryDetails.LongMeanDescription = UpdateDictionary.LongMeanDescription
	}

	db.Save(&dictionaryDetails)
	res, _ := json.Marshal(dictionaryDetails)
	w.Header().Set("Content-Type", "pkglcation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
