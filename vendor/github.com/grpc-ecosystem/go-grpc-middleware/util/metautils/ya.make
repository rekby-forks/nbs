GO_LIBRARY()

LICENSE(Apache-2.0)

SRCS(
    doc.go
    nicemd.go
)

GO_XTEST_SRCS(nicemd_test.go)

END()

RECURSE(
    # gotest
)
