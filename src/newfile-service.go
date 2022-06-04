package src

import (
    "net/http"
    "fmt"

    "github.com/gin-gonic/gin/binding"
    "github.com/gin-gonic/gin"
)

type Args struct {
    file bool `binding:"required"`
}

type function_triggers struct {
    isNew bool
    isNot bool
}

func (s *function_triggers) exec() int {
    exe := true 

    if s.isNew == exe {
        return 0
    } else if s.isNot != exe {
        return 1
    }
    
    return -1
}

// write method for triggering scores
// perspective api res loaded into
// tmp schema.go and saved to db

func Upload(c *gin.Context) {
    var params Args

    if err := c.ShouldBindWith(&params, binding.Query); err == nil {
        // trigger cloud function || trigger controller node
        // we can either invoke cloud function from event change in cloud store
        // or we can invoke parsification directly from api call using controller
        // can we invoke cloud function from api call? - yes but context needed.
        c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("file specification valid %s", params)})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}
