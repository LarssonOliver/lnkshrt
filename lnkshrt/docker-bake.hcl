variable "TAG" {
    default = "latest"
}

variable "REPO" {
    default = "larssonoliver/lnkshrt"
}

group "default" {
    targets = ["lnkshrt"]
}

target "lnkshrt" {
    context = "."
    dockerfile = "Dockerfile"
    platforms = [
        "linux/amd64", 
        "linux/arm/v7", 
        "linux/arm64",
    ]
    tags = [
        "${REPO}:latest", 
        "${REPO}:${TAG}",
    ]
}