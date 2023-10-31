$path = Get-Location

Set-Location $PSScriptRoot

    # generate ts files, install：github.com/mats9693/study/go/goc_ts
    try {
        goc_ts # auto empty output dir
    } catch {
        Write-Output "> generate ts file(s) failed, error: "
        Write-Output $_ # error info
    }

    # copy ts files to its place
    $tsPath = "../ui/src/axios/"

    if (Test-Path "$tsPath") {
        Remove-Item "$tsPath*" -Recurse
    } else {
        try {
            mkdir "$tsPath"
        } catch {
            Write-Output "> mkdir '$tsPath' failed, error: "
            Write-Output $_ # error info
        }
    }

    Copy-Item "ts/*" "$tsPath" -Recurse # use default 'goc_ts' output dir

Set-Location $path

# windows not allow run ps script:
# (admin start)Set-ExecutionPolicy RemoteSigned
