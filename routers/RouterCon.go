package routers

import "github.com/gorilla/mux"

func RouterCon() *mux.Router  {
	return mux.NewRouter()
}
