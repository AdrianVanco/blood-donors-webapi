param (
    $command
)

if (-not $command) {
    $command = "start"
}

$ProjectRoot = "${PSScriptRoot}/.."

$env:BLOOD_DONORS_API_ENVIRONMENT = "Development"
$env:BLOOD_DONORS_API_PORT = "8080"
$env:BLOOD_DONORS_API_MONGODB_USERNAME = "root"
$env:BLOOD_DONORS_API_MONGODB_PASSWORD = "neUhaDnes"

function mongo {
    docker compose --file ${ProjectRoot}/deployments/docker-compose/compose.yaml $args
}

switch ($command) {
    "start" {
        try {
            mongo up --detach
            go run ${ProjectRoot}/cmd/blood-donors-api-service
        }
        finally {
            mongo down
        }
    }
    "test" {
        go test -v ./...
    }
    "mongo" {
        mongo up
    }
    "docker" {
         docker build -t xvancoa/cv2xvancoa-blood-donors-webapi:local-build -f ${ProjectRoot}/build/docker/Dockerfile .
   }
    "openapi" {
        docker run --rm -ti -v ${ProjectRoot}:/local openapitools/openapi-generator-cli generate -c /local/scripts/generator-cfg.yaml
    }
    default {
        throw "Unknown command: $command"
    }
}