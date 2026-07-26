function Get-EnvironmentValue {
    param(
        [string]$Name
    )

    $processValue = [Environment]::GetEnvironmentVariable($Name, "Process")
    if (-not [string]::IsNullOrWhiteSpace($processValue)) {
        return $processValue
    }

    $userValue = [Environment]::GetEnvironmentVariable($Name, "User")
    if (-not [string]::IsNullOrWhiteSpace($userValue)) {
        Set-Item -Path "Env:$Name" -Value $userValue
        return $userValue
    }

    return $null
}

function Test-ConfigurationSection {
    param(
        [string]$SectionName,
        [string[]]$VariableNames
    )

    $missingVariables = @()

    foreach ($variableName in $VariableNames) {
        if ([string]::IsNullOrWhiteSpace((Get-EnvironmentValue -Name $variableName))) {
            $missingVariables += $variableName
        }
    }

    if ($missingVariables.Count -gt 0) {
        Write-Error "$SectionName environment variables are missing: $($missingVariables -join ', ')"
        return $false
    }

    Write-Host "******** $SectionName environment variables validated ********"
    return $true
}

$databaseVariables = @(
    "INSURANCE_DB_HOST",
    "INSURANCE_DB_PORT",
    "INSURANCE_DB_NAME",
    "INSURANCE_DB_USER",
    "INSURANCE_DB_SECRET"
)

$apiVariables = @(
    "INSURANCE_HTTP_ADDRESS",
    "INSURANCE_HTTPS_ENABLED",
    "INSURANCE_MAX_REQUEST_BODY_BYTES"
)

$loggingVariables = @(
    "INSURANCE_LOG_DIRECTORY",
    "INSURANCE_LOG_FILE",
    "INSURANCE_LOG_LEVEL"
)

if (-not (Test-ConfigurationSection -SectionName "Database" -VariableNames $databaseVariables)) {
    exit 1
}

if (-not (Test-ConfigurationSection -SectionName "API" -VariableNames $apiVariables)) {
    exit 1
}

if (-not (Test-ConfigurationSection -SectionName "Logging" -VariableNames $loggingVariables)) {
    exit 1
}

if ((Get-EnvironmentValue -Name "INSURANCE_HTTPS_ENABLED") -eq "true") {
    $tlsVariables = @(
        "INSURANCE_TLS_CERT_FILE",
        "INSURANCE_TLS_KEY_FILE"
    )

    if (-not (Test-ConfigurationSection -SectionName "HTTPS" -VariableNames $tlsVariables)) {
        exit 1
    }
}
Write-Host " "
go run ./cmd/server/main.go
Write-Host " "
Write-Host "******** End of GO Project execution ********"
