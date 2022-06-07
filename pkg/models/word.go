package models

import(
	"github.com/jinzhu/gorm"
	"github.com/rootturk/words-storage/pkg/config"
	"time"
)

var db *gorm.DB

type Words struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Description string `gorm:""json:"description"`
	CreatedOn time.Time `gorm:""json:"createdOn"`
}

type Dictionary struct{
	gorm.Model
	WordId int64 `gorm:""json:"word_id"`
	LanguageId int `gorm:""json:"language_id"`
	TranslateMean string `gorm:""json:"translate_mean"`
	LongMeanDescription string `gorm:""json:"long_mean_description"`
	CreatedOn time.Time `gorm:""json:"createdOn"`
}  


func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Words{})
	db.AutoMigrate(&Dictionary{})
}

//Dictionaries
func(d *Dictionary) CreateDictionary() *Dictionary{
	db.NewRecord(d)
	db.Create(&d)
	return d
}

func GetAllDictionaries() []Dictionary{
	var Dictionary []Dictionary
	db.Find(&Dictionary)
	return Dictionary
}

func GetDictionaryById(Id int64) (*Dictionary, *gorm.DB){
	var getDictionary Dictionary
	db:=db.Where("ID=?",Id).Find(&getDictionary)
	return &getDictionary, db
}

func DeleteDictionary(Id int64) Dictionary{
	var dictionary Dictionary
	db.Where("ID=?",Id).Delete(dictionary)
	return dictionary
}

// WORDS
func(w *Words) CreateWord() *Words{
	db.NewRecord(w)
	db.Create(&w)
	return w
}

func GetAllWords() []Words{
	var Word []Words
	db.Find(&Word)
	return Word
}

func GetWordById(Id int64) (*Words, *gorm.DB){
	var getWord Words
	db:=db.Where("ID=?",Id).Find(&getWord)
	return &getWord, db
}

func DeleteWord(Id int64) Words{
	var word Words
	db.Where("ID=?",Id).Delete(word)
	return word
}