variable "TAG" {
    default = "latest"
}

variable "REPO" {
    default = "larssonoliver/lnkshrt"
}

variable "GUIREPO" {
    default = "larssonoliver/lnkshrt-gui"
}

group "default" {
    targets = ["lnkshrt", "lnkshrt-gui"]
}

target "lnkshrt" {
    context = "lnkshrt"
    dockerfile = "Dockerfile.buildx"
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

target "lnkshrt-gui" {
    context = "gui"
    dockerfile = "Dockerfile.buildx"
    platforms = [
        "linux/amd64", 
        "linux/arm/v7", 
        "linux/arm64",
    ]
    tags = [
        "${GUIREPO}:latest", 
        "${GUIREPO}:${TAG}",
    ]
}