GO_LIBRARY()

LICENSE(Apache-2.0)

SRCS(
    result.go
    scan_raw.go
    scanner.go
    stats.go
)

GO_TEST_SRCS(
    perfomance_test.go
    result_test.go
    scanner_data_test.go
    scanner_test.go
)

END()

RECURSE(
    gotest
)
