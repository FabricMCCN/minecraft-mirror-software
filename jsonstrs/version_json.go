package jsonstrs

type VersionDefine struct {
	Arguments              Arguments              `json:"arguments"`
	AssetIndex             AssetIndex             `json:"assetIndex"` //here
	Assets                 string                 `json:"assets"`
	ComplianceLevel        int64                  `json:"complianceLevel"`
	Downloads              VersionDefindDownloads `json:"downloads"`
	ID                     string                 `json:"id"`
	JavaVersion            JavaVersion            `json:"javaVersion"`
	Libraries              []Library              `json:"libraries"`
	Logging                Logging                `json:"logging"`
	MainClass              string                 `json:"mainClass"`
	MinimumLauncherVersion int64                  `json:"minimumLauncherVersion"`
	ReleaseTime            string                 `json:"releaseTime"`
	Time                   string                 `json:"time"`
	Type                   string                 `json:"type"`
}

type Arguments struct {
	Game []GameElement `json:"game"`
	JVM  []JVMElement  `json:"jvm"`
}

type GameClass struct {
	Rules []GameRule `json:"rules"`
	Value *Value     `json:"value"`
}

type GameRule struct {
	Action   Action   `json:"action"`
	Features Features `json:"features"`
}

type Features struct {
	IsDemoUser          *bool `json:"is_demo_user,omitempty"`
	HasCustomResolution *bool `json:"has_custom_resolution,omitempty"`
}

type JVMClass struct {
	Rules []JVMRule `json:"rules"`
	Value *Value    `json:"value"`
}

type JVMRule struct {
	Action Action   `json:"action"`
	OS     PurpleOS `json:"os"`
}

type PurpleOS struct {
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
	Arch    *string `json:"arch,omitempty"`
}

type AssetIndex struct {
	ID        string `json:"id"`
	Sha1      string `json:"sha1"`
	Size      int64  `json:"size"`
	TotalSize *int64 `json:"totalSize,omitempty"`
	URL       string `json:"url"`
}

type VersionDefindDownloads struct {
	Client         ClientMappingsClass `json:"client"`
	ClientMappings ClientMappingsClass `json:"client_mappings"`
	Server         ClientMappingsClass `json:"server"`
	ServerMappings ClientMappingsClass `json:"server_mappings"`
}

type ClientMappingsClass struct {
	Sha1 string  `json:"sha1"`
	Size int64   `json:"size"`
	URL  string  `json:"url"`
	Path *string `json:"path,omitempty"` //use this!
}

type JavaVersion struct {
	Component    string `json:"component"`
	MajorVersion int64  `json:"majorVersion"`
}

type Library struct {
	Downloads LibraryDownloads `json:"downloads"`
	Name      string           `json:"name"`
	Rules     []LibraryRule    `json:"rules,omitempty"`
	Natives   *Natives         `json:"natives,omitempty"`
	Extract   *Extract         `json:"extract,omitempty"`
}

type LibraryDownloads struct {
	Artifact    ClientMappingsClass `json:"artifact"`
	Classifiers *Classifiers        `json:"classifiers,omitempty"`
}

type Classifiers struct {
	Javadoc        *ClientMappingsClass `json:"javadoc,omitempty"`
	NativesLinux   *ClientMappingsClass `json:"natives-linux,omitempty"`
	NativesMacos   *ClientMappingsClass `json:"natives-macos,omitempty"`
	NativesWindows *ClientMappingsClass `json:"natives-windows,omitempty"`
	Sources        *ClientMappingsClass `json:"sources,omitempty"`
	NativesOsx     *ClientMappingsClass `json:"natives-osx,omitempty"`
}

type Extract struct {
	Exclude []string `json:"exclude"`
}

type Natives struct {
	Osx     *string `json:"osx,omitempty"`
	Linux   *string `json:"linux,omitempty"`
	Windows *string `json:"windows,omitempty"`
}

type LibraryRule struct {
	Action Action    `json:"action"`
	OS     *FluffyOS `json:"os,omitempty"`
}

type FluffyOS struct {
	Name Name `json:"name"`
}

type Logging struct {
	Client LoggingClient `json:"client"`
}

type LoggingClient struct {
	Argument string     `json:"argument"`
	File     AssetIndex `json:"file"` //here
	Type     string     `json:"type"`
}

type Action string

const (
	Allow    Action = "allow"
	Disallow Action = "disallow"
)

type Name string

const (
	Osx Name = "osx"
)

type GameElement struct {
	GameClass *GameClass
	String    *string
}

type Value struct {
	String      *string
	StringArray []string
}

type JVMElement struct {
	JVMClass *JVMClass
	String   *string
}
