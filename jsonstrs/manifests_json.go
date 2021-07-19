package jsonstrs

type Manifest struct {
	Latest   Latest    `json:"latest"`
	Versions []Version `json:"versions"`
}

type Latest struct {
	Release  string `json:"release"`
	Snapshot string `json:"snapshot"`
}

type Version struct {
	ID              string `json:"id"`
	Type            Type   `json:"type"`
	URL             string `json:"url"`
	Time            string `json:"time"`
	ReleaseTime     string `json:"releaseTime"`
	Sha1            string `json:"sha1"`
	ComplianceLevel int64  `json:"complianceLevel"`
}

type Type string

const (
	OldAlpha Type = "old_alpha"
	OldBeta  Type = "old_beta"
	Release  Type = "release"
	Snapshot Type = "snapshot"
)
