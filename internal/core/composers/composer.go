package composers

// Composer is the type that represents a message composer
// used for composing message in the producing phase, starting
// from raw bytes.
// Returns the bytes of the composed message, the content type and an error
type Composer = func(rawBytes []byte) ([]byte, string, error)
