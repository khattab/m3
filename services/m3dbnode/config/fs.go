		HugeTLB: MmapHugeTLBConfiguration{
	// HugeTLB is the huge pages configuration which will only take affect
	HugeTLB MmapHugeTLBConfiguration `yaml:"hugeTLB"`
// MmapHugeTLBConfiguration is the mmap huge TLB configuration.
type MmapHugeTLBConfiguration struct {
	// Threshold is the threshold on which to use the huge TLB flag if enabled