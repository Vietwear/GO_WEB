function Install-Software {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)]
        [string] $version
    )

    Write-Host "I installed software version $version. Yippee!"
}
Install-Software -version 2000GetFind-Command

Find-DSCResource