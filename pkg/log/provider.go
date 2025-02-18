package log

func ProviderSet() Log {
	return NewNativeLog()
}
