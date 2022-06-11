package src

import (
    "net/http"
    "fmt"
    "os"
    "io"
//    "io/ioutil"
//    "github.com/gin-gonic/gin/binding"
    "github.com/gin-gonic/gin"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func IsDirEmpty(name string) (bool, error) {
     f, err := os.Open(name)
     
     if err != nil {
         return false, err
     }     

     defer f.Close()

     // read in ONLY one file
     _, err = f.Readdir(1)

     // and if the file is EOF... well, the dir is empty.
     if err == io.EOF {
        return true, nil
     }

     return false, err
}

func Parse() string {
    // files, err := ioutil.ReadDir("./local/") <- memory addr
    dat, err := os.ReadFile("./local/payload.txt")
    check(err)
    
    return string(dat)
}

// write method for triggering scores
// perspective api res loaded into
// tmp schema.go and saved to db
func Upload(c *gin.Context) {
    ok, err := IsDirEmpty("./local/")
    
    if err != nil {
        fmt.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
    }
    
    if ok == false {
        c.JSON(http.StatusOK, gin.H{"response": string(Parse())})
    } else { 
        c.JSON(http.StatusOK, gin.H{"response": "Directory empty.."})
    }

    // trigger cloud function || trigger controller node
    // we can either invoke cloud function from event change in cloud store
    // or we can invoke parsification directly from api call using controller
    // can we invoke cloud function from api call? - yes but context needed
}
