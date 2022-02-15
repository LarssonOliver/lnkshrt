package app

func (a *App) initRoutes() {
	a.Router.Use(LoggingMiddleware)
	a.Router.HandleFunc("/", a.CreateLink).Methods("POST")
	a.Router.HandleFunc("/{id}", a.ResolveLink).Methods("GET")
}
