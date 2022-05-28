package controllers

import (
    "fmt"
)

// gateway for services
type key struct {
    Service string
    Data map[string]string `omitempty`
}
