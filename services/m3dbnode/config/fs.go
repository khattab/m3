// DefaultMmapConfiguration is the default mmap configuration.
func DefaultMmapConfiguration() MmapConfiguration {
	return MmapConfiguration{
		HugePages: MmapHugePagesConfiguration{
			Enabled:   true,    // Enable when on a platform that supports
			Threshold: 2 << 14, // 32kb and above mappings use huge pages
		},
	}
}


	// Mmap is the mmap options which features are primarily platform dependent
	Mmap *MmapConfiguration `yaml:"mmap"`
}

// MmapConfiguration is the mmap configuration.
type MmapConfiguration struct {
	// HugePages is the huge pages configuration which will only take affect
	// on platforms that support it, currently just linux
	HugePages MmapHugePagesConfiguration `yaml:"hugePages"`
}

// MmapHugePagesConfiguration is the mmap huge pages configuration.
type MmapHugePagesConfiguration struct {
	// Enabled if true or disabled if false
	Enabled bool `yaml:"enabled"`

	// Threshold is the threshold on which to use huge pages if enabled
	Threshold int64 `yaml:"threshold"`

// MmapConfiguration returns the effective mmap configuration.
func (p FilesystemConfiguration) MmapConfiguration() MmapConfiguration {
	if p.Mmap == nil {
		return DefaultMmapConfiguration()
	}
	return *p.Mmap
}