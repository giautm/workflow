package simple

import (
	"log"

	wf "github.com/giautm/workflow"
)

var create = wf.ActivityFunc(func(wf.Context) error {
	log.Println("Propose the request to the first manager to approve")
	return nil
})

var complete = wf.ActivityFunc(func(wf.Context) error {
	log.Println("The process is completed")
	return nil
})

var mail = wf.ActivityFunc(func(context wf.Context) error {
	if msg, ok := context.Param("msg"); ok {
		log.Printf("Send mail: %s \n", msg)
	}
	return nil
})
