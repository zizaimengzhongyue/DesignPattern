package main

import "./administration"

func main() {
	assistant := administration.Assistant{IT: administration.IT{}, HR: administration.HR{}}
	assistant.ITSupport()
	assistant.HRSupport()
}
