package version

var (
	Version     string
	FullCommit  string
	ReleaseDate string
	BuildDate   string
)

type PackageVersion struct {
	Version     string `json:"version"`
	FullCommit  string `json:"full_commit"`
	ReleaseDate string `json:"release_date"`
	BuildDate   string `json:"build_date"`
}
