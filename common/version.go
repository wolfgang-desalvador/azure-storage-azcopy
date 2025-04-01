package common

<<<<<<< HEAD
const AzcopyVersion = "10.28.1~preview"
=======
const AzcopyVersion = "10.29.0-Preview"
>>>>>>> edb71da0a00b05b01321756da17ec827edea1d43
const UserAgent = "AzCopy/" + AzcopyVersion
const S3ImportUserAgent = "S3Import " + UserAgent
const GCPImportUserAgent = "GCPImport " + UserAgent
const BenchmarkUserAgent = "Benchmark " + UserAgent

// AddUserAgentPrefix appends the global user agent prefix, if applicable
func AddUserAgentPrefix(userAgent string) string {
	prefix := GetEnvironmentVariable(EEnvironmentVariable.UserAgentPrefix())
	if len(prefix) > 0 {
		userAgent = prefix + " " + userAgent
	}

	return userAgent
}
