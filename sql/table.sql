CREATE TABLE IF NOT EXISTS BenchData (
    Date DATE,
    Commit VARCHAR(255),
    Branch VARCHAR(255),
    CompilerVersion VARCHAR(255),
    Arch VARCHAR(255),
    TestName VARCHAR(255),
    NsPerOp BIGINT,
    AllocsPerOp BIGINT,
    BytesPerOp BIGINT,
    PRIMARY KEY (Commit, TestName)
);