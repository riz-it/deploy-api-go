# Gitlab Access Token
variable "gitlab_access_token" {
  type        = string
  description = "Gitlab Access Token"
  sensitive   = true
}

# Repository List
variable "repository_list" {
  type = list(object({
    name        = string
    description = string
    visibility  = string
  }))
}
