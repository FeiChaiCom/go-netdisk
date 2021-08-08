package version

const (
	PlatformName = "feichai-disk"
)

var (
	// Edition will be injected during build time.
	Edition = InnerEdition
)

const (
	InnerEdition = "inner_edition"
)

func GetEdition() string {
	return Edition
}
