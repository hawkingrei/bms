CREATE TABLE IF NOT EXISTS bench_data (
    date DATE,
    commit VARCHAR(255),
    branch VARCHAR(255),
    compiler_version VARCHAR(255),
    arch VARCHAR(255),
    test_name VARCHAR(255),
    ns_per_op FLOAT,
    allocs_per_op BIGINT,
    bytes_per_op BIGINT
);