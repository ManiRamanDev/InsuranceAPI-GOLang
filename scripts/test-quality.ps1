$ErrorActionPreference = "Stop"

Write-Host "==> Discovering packages"
$allPackages = @(go list ./...)
$unitPackages = @($allPackages | Where-Object { $_ -notmatch "/integration$" })

if (-not $unitPackages -or $unitPackages.Count -eq 0) {
    Write-Error "No unit packages found"
    exit 1
}

Write-Host "==> Unit-focused coverage (excluding integration)"
go test -coverprofile coverage-unit.out @unitPackages
$unitLine = go tool cover -func coverage-unit.out | Select-String "^total:"
if (-not $unitLine) {
    Write-Error "Unable to read unit coverage summary"
    exit 1
}
$unitLineText = $unitLine.ToString()
Write-Host $unitLineText

$unitPercentMatch = [regex]::Match($unitLineText, "([0-9]+(\.[0-9]+)?)%")
if (-not $unitPercentMatch.Success) {
    Write-Error "Unable to parse unit coverage percent"
    exit 1
}
$unitPercent = [double]$unitPercentMatch.Groups[1].Value
if ($unitPercent -lt 80.0) {
    Write-Error ("Unit coverage gate failed: {0}% < 80%" -f $unitPercent)
    exit 1
}
Write-Host ("Unit coverage gate passed: {0}% >= 80%" -f $unitPercent)

Write-Host "==> Full-workspace coverage (including integration)"
go test -coverprofile coverage-all.out ./...
$allLine = go tool cover -func coverage-all.out | Select-String "^total:"
if (-not $allLine) {
    Write-Error "Unable to read full coverage summary"
    exit 1
}
Write-Host $allLine.ToString()

Write-Host "==> Benchmark run (all packages)"
go test -bench . -benchmem ./...

Write-Host "==> Quality checks completed"
