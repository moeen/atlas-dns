package router

import "fmt"

// Start will start the http server with given port and the router
func Start(port int) error {
	r := Router()

	return r.Run(fmt.Sprintf(":%d", port))
}
