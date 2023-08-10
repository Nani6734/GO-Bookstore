 package contaollers
 import(

"fmt"
"encoding/json"
"github.com/gorilla/mux"
"net/http"
"Strconv"
"github.com/vinith/go-bookstore/pkg/utils"
"github.com/vinith/go-bookstore/pkg/models"

)

var NetBook models.Book
 
func GetBook(w http.ResponseWriter, r *http.Request){

	newBooks:=models.GetAllBooks()
	res,_ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglicaton/json")
	w.WriteHeader(http.StatusOK)
    w.Write(res)



func GetBookById(w http.ResponseWriter,r *http.Request){
	vars :=mux.Vars(r)
bookId := vars("bookId")
ID, err := strconv.ParseInt(bookId,0,0)
if err !=nil{
	fmt.Println("error whle parsing")
}
bookDetails,_ :=models.GetBookById(ID)
res,_ :=json.Marshal(bookDetails)
w.Header().Set("Content-Type","pkglication/json")
w.WriterHeader(http.StatusOk)
w.Write(res)
}


func CreateBook(w http.ResponseWriter,r *http.Request){
CreateBook := &models.Book{}
utils.ParsBody(r,CreateBook)
b:= CreateBook.CreateBook()
res,_ :=json.Marshal(b)
w.WriteHeader(http.StatusOK)
w.Writer(res)
} 

func DeleteBook(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	bookId :=vars["bookId"]
	Id,err :=strconv.ParseInt(bookId,0,0)
	if err !=nil{
		fmt.println("error while parsing")
	}

	book :=models.DeleteBook(ID)
	res,_ :=json.Marshal(book)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter,r *http.Request){
	var UpdateBook =&models.Book{}
	utils.parseBody(r, UpdateBook)
	vars :=mux.Vars(r)
	bookId := vars["bookId"]
	Id,err :=strconv.parseInt(bookId,0,0)
	if err !=nil{
		fmt.println("error while parsing")
	}
	
	bookDetails,db:=models.GetBodyById(ID)
	if updatebook. name !=""{
		bookDetails.name =updateBook.Name
	}

	if updateBook.Author !=""{
      bookDetails.Auther =updateBook.Author

	}
	if updateBook.publication !=""{
    bookDetails.Publication=UpdateBook.publication
	}
	DB.SAVE(&bookDetails)
	res,_ :=json.Marshal(bookDetails)
    w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
   w.writer(res)
}