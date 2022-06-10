package fibn

// ----

func fibnRun() {
	firstpoint := new()
	latest := firstpoint

	for i := 0; i < 50; i++ {
		latest = latest.createNext()
	}

	firstpoint.printFromHere()
}
