variable "tag" {
    default = "0.1"
}

variable "repo" {
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
        "${repo}:latest", 
        "${repo}:${tag}",
    ]
}