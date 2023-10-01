$path = Get-Location

Set-Location $PSScriptRoot
    if (!(Test-Path "./ui/dist")) {
        Write-Output "please build web code first"
        return
    }

    if (Test-Path "./build") {
        Remove-Item "./build/*" -Recurse
    } else {
        mkdir "./build"
    }

    # binary executable file
    go build -o "listen_bilibili.exe"

    Move-Item "listen_bilibili.exe" -Destination "./build/listen_bilibili.exe"
    Copy-Item "listen_bilibili.yaml" -Destination "./build/listen_bilibili.yaml"

    # user manual
    Copy-Item "README.md" -Destination "./build/README.md"

    # web code
    mkdir "./build/ui"
    Copy-Item "./ui/dist/" -Destination "./build/ui/" -Recurse

Set-Location $path

# windows not allow run ps script:
# (admin start)Set-ExecutionPolicy RemoteSigned
