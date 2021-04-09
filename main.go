package main


import(
    "log"
     "database/sql"
	_ "github.com/go-sql-driver/mysql"
    "flag"
)


var u string
var p string
var h string
var pass string
var dbn string
var db *sql.DB


func main(){
    
    
    flag.StringVar(&u,"user","","User")
    flag.StringVar(&pass,"pass","","Password")
    flag.StringVar(&p,"port","","Port")
    flag.StringVar(&h,"host","","Host")
    flag.StringVar(&dbn,"db","","Database Name")
    
    flag.Parse()
    
    myDsn := Dsn()	
    
    db, errc := sql.Open("mysql", myDsn)

    if(errc!=nil){
        	
        	log.Println(errc)
    		return
            	
    }
    	
    	
    err := db.Ping()
    	
    if err != nil {
    		log.Println(err)
    		return
    }
    
    defer db.Close()
    
    log.Println("Good Conection")


}

func Dsn() string{

    
	
	str := u+":"+pass+"@tcp("+h+":"+p+")/"+dbn
    
    log.Println(str)
    
    return str
    
}