# Database configuration
setx INSURANCE_DB_HOST "localhost"
setx INSURANCE_DB_PORT "5432"
setx INSURANCE_DB_NAME "insurance_api"
setx INSURANCE_DB_USER "postgres"
setx INSURANCE_DB_SECRET "dmlqYXkxMjM="

# API configuration
setx INSURANCE_HTTP_ADDRESS ":8443"
setx INSURANCE_MAX_REQUEST_BODY_BYTES "512000"

# Logging configuration
setx INSURANCE_LOG_DIRECTORY "logs"
setx INSURANCE_LOG_FILE "insurance-api.log"
setx INSURANCE_LOG_LEVEL "INFO"

# Enable these only after certificate files are available.
setx INSURANCE_HTTPS_ENABLED "true"
setx INSURANCE_TLS_CERT_FILE "C:/certificates/rootSSL.crt"
setx INSURANCE_TLS_KEY_FILE "C:/certificates/rootSSL.key"
