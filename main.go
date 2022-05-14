package main

import (
	"github.com/gjlanc65/gintest/docs"
	_ "github.com/gjlanc65/gintest/docs"
	"github.com/gjlanc65/gintest/routes"
)

// @title Go + Gin Trip Segment API
// @version 1.0
// @description This is a sample server. You can visit the GitHub repository at https://github.com/GJLANC65/gintest

// @contact.name API Support
// @contact.url https://github.com/GJLANC65/gintest
// @contact.email [protected]

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

func main() {
	docs.SwaggerInfo.BasePath = "/"

	router := routes.SetupRouter()
	// router.SetTrustedProxies([]string{"192.168.1.2"})
	router.SetTrustedProxies(nil)
	router.Run("localhost:8080")

	/* TODO : Later - 12 Factor App
	// See also https://liamclarke.dev/post/building-a-graceful-http-server-in-go/
	//          https://github.com/clarke94/serverfx

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")

	*/
}
