$path = Get-Location

Set-Location $PSScriptRoot
    # make sure web code exist
    if (!(Test-Path "./ui/dist")) {
        Write-Output "please build web code first"
        return
    }

    # mkdir output dir
    if (Test-Path "./listen") {
        Remove-Item "./listen/*" -Recurse
    } else {
        mkdir "./listen"
    }

    # binary executable file
    go build -o "listen_bilibili.exe"

    Move-Item "listen_bilibili.exe" -Destination "./listen/listen_bilibili.exe"
    Copy-Item "list.yaml" -Destination "./listen/list.yaml"

    # user manual
    Copy-Item "README.md" -Destination "./listen/manual.md"

    # web code
    mkdir "./listen/ui"
    Copy-Item "./ui/dist/" -Destination "./listen/ui/" -Recurse

Set-Location $path

# windows not allow run ps script:
# (admin start)Set-ExecutionPolicy RemoteSigned
