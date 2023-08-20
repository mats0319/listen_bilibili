# Quick Start: https://grpc.io/docs/languages/go/quickstart/
# require:
# 1. protoc, download: https://github.com/protocolbuffers/protobuf/releases, remember set PATH
# 2. protoc-gen-go, install script: go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
# 3. protoc-gen-go-grpc, install script: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

$path = Get-Location

Set-Location $PSScriptRoot

    if (Test-Path "./go") {
#        Remove-Item "./go/*"
    } else {
        mkdir "./go"
    }

    # generate go files
#    protoc --go_out=./go --go_opt=paths=source_relative .\*.proto

    # generate ts files, remember install tool
    try {
        protoc_ts
    }
    catch {
        Write-Output "> generate ts file(s) failed, error: "
        Write-Output $_ # error info
    }

Set-Location $path

# windows not allow run ps script:
# (admin start)Set-ExecutionPolicy RemoteSigned
