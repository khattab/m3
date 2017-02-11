	"github.com/m3db/m3db/persist/encoding/msgpack"
	"github.com/m3db/m3db/storage/block"
	"github.com/m3db/m3db/x/io"
	"github.com/m3db/m3x/checked"
	"github.com/m3db/m3x/instrument"
	"github.com/m3db/m3x/pool"
	Write(id ts.ID, data checked.Bytes, checksum uint32) error
	WriteAll(id ts.ID, data []checked.Bytes, checksum uint32) error
	// Read returns the next id, data, checksum tuple or error, will return io.EOF at end of volume.
	// Use either Read or ReadMetadata to progress through a volume, but not both.
	Read() (id ts.ID, data checked.Bytes, checksum uint32, err error)

	// ReadMetadata returns the next id and metadata or error, will return io.EOF at end of volume.
	// Use either Read or ReadMetadata to progress through a volume, but not both.
	ReadMetadata() (id ts.ID, length int, checksum uint32, err error)
// FileSetSeeker provides an out of order reader for a TSDB file set
type FileSetSeeker interface {
	io.Closer

	// Open opens the files for the given shard and version for reading
	Open(namespace ts.ID, shard uint32, start time.Time) error

	// Seek returns the data for specified ID provided the index was loaded upon open. An
	// error will be returned if the index was not loaded or ID cannot be found.
	Seek(id ts.ID) (data checked.Bytes, err error)

	// SeekOffset returns the offset for specified ID provided the index was loaded upon open. If
	// the index was not loaded or ID cannot be found the value -1 will be returned.
	// This can be helpful ahead of issuing a number of seek requests so that the seek
	// requests can be made in order.
	SeekOffset(id ts.ID) int

	// Range returns the time range associated with data in the volume
	Range() xtime.Range

	// Entries returns the count of entries in the volume
	Entries() int

	// IDs retrieves all the identifiers present in the file set
	IDs() []ts.ID
}

// FileSetSeekerManager provides management of seekers for a TSDB namespace.
type FileSetSeekerManager interface {
	io.Closer

	// Open opens the seekers for a given namespace.
	Open(namespace ts.ID) error

	// CacheShardIndices will pre-parse the indexes for given shards
	// to improve times when seeking to a block.
	CacheShardIndices(shards []uint32) error

	// Seeker returns an open seeker for a given shard and block start time.
	Seeker(shard uint32, start time.Time) (FileSetSeeker, error)
}

// BlockRetriever provides a block retriever for TSDB file sets
type BlockRetriever interface {
	io.Closer
	block.DatabaseBlockRetriever

	// Open the block retriever to retrieve from a namespace
	Open(namespace ts.ID) error
}

// RetrievableBlockSegmentReader is a retrievable block reader
type RetrievableBlockSegmentReader interface {
	xio.SegmentReader
}

	// SetDecodingOptions sets the decoding options
	SetDecodingOptions(value msgpack.DecodingOptions) Options

	// DecodingOptions returns the decoding options
	DecodingOptions() msgpack.DecodingOptions


// BlockRetrieverOptions represents the options for block retrieval
type BlockRetrieverOptions interface {
	// SetRequestPoolOptions sets the request pool options
	SetRequestPoolOptions(value pool.ObjectPoolOptions) BlockRetrieverOptions

	// RequestPoolOptions returns the request pool options
	RequestPoolOptions() pool.ObjectPoolOptions

	// SetBytesPool sets the bytes pool
	SetBytesPool(value pool.CheckedBytesPool) BlockRetrieverOptions

	// BytesPool returns the bytes pool
	BytesPool() pool.CheckedBytesPool

	// SetSegmentReaderPool sets the segment reader pool
	SetSegmentReaderPool(value xio.SegmentReaderPool) BlockRetrieverOptions

	// SegmentReaderPool returns the segment reader pool
	SegmentReaderPool() xio.SegmentReaderPool

	// SetFetchConcurrency sets the fetch concurrency
	SetFetchConcurrency(value int) BlockRetrieverOptions

	// FetchConcurrency returns the fetch concurrency
	FetchConcurrency() int
}