[CmdletBinding()]
param()

if ("SilentlyContinue" -eq $VerbosePreference) {
    go test ./main
} else {
    go test -v ./main
}
