[CmdletBinding()]
param()

Write-Verbose "Install go bundle tool..."
go install golang.org/x/tools/cmd/bundle@latest

Write-Verbose "Script root: $PSScriptRoot"
$ProjectRoot = (Get-Item $PSScriptRoot).Parent.FullName
Write-Verbose "Project root: $ProjectRoot"
Push-Location $ProjectRoot
Write-Verbose "Bundle..."
& ($env:USERPROFILE + "/go/bin/bundle") -o dist/main.go -dst ./main -prefix '""'  github.com/loicpetit/codingame-fall-challenge-2023/main
