package testmodedb

type DB struct {
	Open     bool
	NotFatal bool
}

type TestMode struct {
	ForceMashalling bool
	DB              DB
}

var testModeCfg TestMode

func SetTestMode(testMode *TestMode) {
	if testMode != nil {
		testModeCfg = *testMode
		return
	}

	testModeCfg = TestMode{
		ForceMashalling: true,
		DB:              DB{Open: true, NotFatal: true},
	}
}

func GetTestMode() TestMode {
	return testModeCfg
}
