package config

// private
var app_name = "Go Basic"
var version = "1.0"

// access private var
func GetName() string {
	return app_name
}

func GetVersion() string {
	return version
}
