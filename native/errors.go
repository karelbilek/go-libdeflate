package native

import "errors"

var (
	errorOutOfMemory       = errors.New("libdeflate: native: out of memory")
	errorInvalidLevel      = errors.New("libdeflate: native: illegal compression level")
	errorShortBuffer       = errors.New("libdeflate: native: short buffer")
	errorNoInput           = errors.New("libdeflate: native: empty input")
	errorBadData           = errors.New("libdeflate: native: bad data: data was corrupted, invalid or unsupported")
	errorUnknown           = errors.New("libdeflate: native: unknown error code from c library")
	errorShortOutput       = errors.New("libdeflate: native: buffer too long: decompressed to fewer bytes than expected, indicating possible error in decompression. Make sure your out buffer has the exact length of the decompressed data or pass nil for out")
	errorInsufficientSpace = errors.New("libdeflate: native: buffer too short. Retry with larger buffer")
)
