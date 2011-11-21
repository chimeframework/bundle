package controllers

import (
    "fmt"
    "strings"
    kernel "chime/components/httpkernel"
)

type ControllerNameParser struct {
    kernel kernel.Kerneler
}

func NewControllerNameParser(kernel kernel.Kerneler) *ControllerNameParser {
    this := &ControllerNameParser{}
    this.kernel = kernel
    return this
}


func (this *ControllerNameParser) Parse(controller string) *kernel.Callable{
    parts := strings.Split(controller, kernel.SEPARATOR)
    logs := make([]string, 0)
    names := make([]string, 0)

    if len(parts) != 3{
        panic(fmt.Sprintf("The %v is not a valid a:b:c controller string.", controller))
    }

    bundle, controller, action := parts[0], parts[1], parts[2]

    var callable *kernel.Callable = nil
    for _, b := range this.kernel.GetBundle(bundle){
        callable = b.GetCallable(controller, action)
        if callable != nil{
            logs = append(logs, fmt.Sprintf("Unable to find controller %s:%s in bundle: %s", controller,action, bundle))
        } 
        names= append(names, b.GetName())
    }

    if callable == nil{
        handleControllerNotFound(bundle, controller, logs, names)
    }
    return callable
}

func handleControllerNotFound(bundle string, controller string, logs []string, names []string) {
    if len(logs) == 1{
        panic(logs[0])
    }
    msg := fmt.Sprintf("Unable to find controller %s:%s in bundles %s", bundle, controller, strings.Join(names, ", "))
    panic(msg)
}
