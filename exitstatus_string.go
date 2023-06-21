// Code generated by "stringer -type=ExitStatus"; DO NOT EDIT.

package arigo

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Success-0]
	_ = x[UnknownError-1]
	_ = x[Timeout-2]
	_ = x[ResourceNotFound-3]
	_ = x[ResourceNotFoundReached-4]
	_ = x[DownloadSpeedTooSlow-5]
	_ = x[NetworkError-6]
	_ = x[UnfinishedDownloads-7]
	_ = x[RemoteNoResume-8]
	_ = x[NotEnoughDiskSpace-9]
	_ = x[PieceLengthMismatch-10]
	_ = x[SameFileBeingDownloaded-11]
	_ = x[SameInfoHashBeingDownloaded-12]
	_ = x[FileAlreadyExists-13]
	_ = x[RenamingFailed-14]
	_ = x[CouldNotOpenExistingFile-15]
	_ = x[CouldNotCreateNewFile-16]
	_ = x[FileIOError-17]
	_ = x[CouldNotCreateDirectory-18]
	_ = x[NameResolutionFailed-19]
	_ = x[MetalinkParsingFailed-20]
	_ = x[FTPCommandFailed-21]
	_ = x[HTTPResponseHeaderBad-22]
	_ = x[TooManyRedirects-23]
	_ = x[HTTPAuthorizationFailed-24]
	_ = x[BencodedFileParseError-25]
	_ = x[TorrentFileCorrupt-26]
	_ = x[MagnetURIBad-27]
	_ = x[RemoteServerHandleRequestError-28]
	_ = x[JSONRPCParseError-29]
	_ = x[Reserved-30]
	_ = x[ChecksumValidationFailed-31]
}

const _ExitStatus_name = "SuccessUnknownErrorTimeoutResourceNotFoundResourceNotFoundReachedDownloadSpeedTooSlowNetworkErrorUnfinishedDownloadsRemoteNoResumeNotEnoughDiskSpacePieceLengthMismatchSameFileBeingDownloadedSameInfoHashBeingDownloadedFileAlreadyExistsRenamingFailedCouldNotOpenExistingFileCouldNotCreateNewFileFileIOErrorCouldNotCreateDirectoryNameResolutionFailedMetalinkParsingFailedFTPCommandFailedHTTPResponseHeaderBadTooManyRedirectsHTTPAuthorizationFailedBencodedFileParseErrorTorrentFileCorruptMagnetURIBadRemoteServerHandleRequestErrorJSONRPCParseErrorReservedChecksumValidationFailed"

var _ExitStatus_index = [...]uint16{0, 7, 19, 26, 42, 65, 85, 97, 116, 130, 148, 167, 190, 217, 234, 248, 272, 293, 304, 327, 347, 368, 384, 405, 421, 444, 466, 484, 496, 526, 543, 551, 575}

func (i ExitStatus) String() string {
	if i >= ExitStatus(len(_ExitStatus_index)-1) {
		return "ExitStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ExitStatus_name[_ExitStatus_index[i]:_ExitStatus_index[i+1]]
}
