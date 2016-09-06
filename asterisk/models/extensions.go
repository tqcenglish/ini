package models

// WebRTCCallCenter WebRTC 模版
type WebRTCCallCenter struct {
	CidNumber     string   `ini:"cid_number"`
	Language      string   `ini:"language"`
	Srtpapable    bool     `ini:"srtpapable"`
	Encryption    bool     `ini:"encryption"`
	Avpf          bool     `ini:"avpf"`
	ForceAvp      bool     `ini:"force_avp"`
	Nat           []string `ini:"nat"`
	Hints         bool     `ini:"hints"`
	Hassip        bool     `ini:"hassip"`
	Icesupport    bool     `ini:"icesupport"`
	Dtlsenable    bool     `ini:"dtlsenable"`
	Dtlsverify    bool     `ini:"dtlsverify"`
	Dtlscertfile  string   `ini:"dtlscertfile"`
	Dtlscafile    string   `ini:"dtlscafile"`
	Dtlssetup     string   `ini:"dtlssetup"`
	Transport     string   `ini:"transport"`
	Videosupport  bool     `ini:"videosupport"`
	Transfer      bool     `ini:"transfer"`
	ExtensionType string   `ini:"type"`
	Qualifyfreq   int      `ini:"qualifyfreq"`
	Hasdirectory  bool     `ini:"hasdirectory"`
	Host          string   `ini:"host"`
	Disallow      string   `ini:"disallow"`
	Allow         []string `ini:"allow"`
}

// SipCallCenter WebRTC 模版
type SipCallCenter struct {
	CidNumber     string   `ini:"cid_number"`
	Language      string   `ini:"language"`
	Srtpapable    bool     `ini:"srtpapable"`
	Encryption    bool     `ini:"encryption"`
	Avpf          bool     `ini:"avpf"`
	ForceAvp      bool     `ini:"force_avp"`
	Nat           []string `ini:"nat"`
	Hints         bool     `ini:"hints"`
	Hassip        bool     `ini:"hassip"`
	Icesupport    bool     `ini:"icesupport"`
	Dtlsenable    bool     `ini:"dtlsenable"`
	Dtlsverify    bool     `ini:"dtlsverify"`
	Dtlscertfile  string   `ini:"dtlscertfile"`
	Dtlscafile    string   `ini:"dtlscafile"`
	Dtlssetup     string   `ini:"dtlssetup"`
	Transport     string   `ini:"transport"`
	Videosupport  bool     `ini:"videosupport"`
	Transfer      bool     `ini:"transfer"`
	ExtensionType string   `ini:"type"`
	Qualifyfreq   int      `ini:"qualifyfreq"`
	Hasdirectory  bool     `ini:"hasdirectory"`
	Host          string   `ini:"host"`
	Disallow      string   `ini:"disallow"`
	Allow         []string `ini:"allow"`
}

// Extension 分机结构体
type Extension struct {
	CidNumber string `ini:"cid_number"`
	Fullname  string `ini:"fullname"`
	Context   string `ini:"context"`
	Secret    string `ini:"secret"`
}
