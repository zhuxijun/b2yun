//go:generate goversioninfo
package main

func main() {

	app := App{}
	app.Initialize()

	if err := app.Run(); err != nil {
		app.log.Logger.Error(err)
	}

}
