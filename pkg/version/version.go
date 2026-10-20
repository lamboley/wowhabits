package version

var gitVersion = "v0.0.0-dev"

// Get returns the version of the running binary.
func Get() string {
	return gitVersion
}
