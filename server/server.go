package server

func Init() {
	// can be used to set port c := config.GetConfig()

	r := NewRouter()
	r.Run()
}