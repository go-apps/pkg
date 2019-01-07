package core

var gAPP *coreApplication

// CoreApplication singleton
func CoreApplication() *coreApplication {
	if gAPP == nil {
		gAPP = new(coreApplication)
	}

	return gAPP
}

func init() {
	CoreApplication()
}
