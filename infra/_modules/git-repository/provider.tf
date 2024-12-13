terraform {
  required_providers {
    gitlab = {
      source  = "gitlabhq/gitlab"
      version = "17.2.0"
    }
  }
}

provider "gitlab" {
  token = var.gitlab_access_token
}