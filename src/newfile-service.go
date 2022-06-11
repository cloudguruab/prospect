package src

import (
    "bufio"
    "net/http"
    "fmt"
    "os"
    "io"
    "github.com/cloudguruab/prospect/controllers"
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

    
// write method for triggering scores
// perspective api res loaded into
// tmp schema.go and saved to db
func Upload(c *gin.Context) {
    ok, err := IsDirEmpty("./local/")
    var score_bucket []string    
    
    if err != nil {
        fmt.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
    }
    
    
    if ok == false {
        data, err := os.Open("./local/payload.txt")
        check(err)
        
        defer data.Close()
        
        obj := bufio.NewScanner(data)
        for obj.Scan() {
          score_bucket = append(score_bucket, obj.Text())
        }
        
        check(obj.Err())
    
        c.JSON(http.StatusOK, gin.H{"response": "Parsing..."})
    } else { 
        c.JSON(http.StatusOK, gin.H{"response": "Directory empty.."})
    }

    
    // taking this further we can create struct - return struct over api
    for i := 0; i <= len(score_bucket); i++ {
        fmt.Printf("%s \n", score_bucket[i])
        controllers.AnalyzeIt(score_bucket[i]) 
        fmt.Println("\n")
    }
    
    // to continue...
    // trigger cloud function || trigger controller node
    // we can either invoke cloud function from event change in cloud store
    // or we can invoke parsification directly from api call using controller
}
