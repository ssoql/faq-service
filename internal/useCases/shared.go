package useCases

type testShutdownSig struct{}

func (testShutdownSig) String() string {
	return "shutdown signal"
}

func (testShutdownSig) Signal() {}
