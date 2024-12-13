resource "gitlab_project" "repository" {
  for_each         = { for idx, repo in var.repository_list : repo.name => repo }
  name             = each.value.name
  description      = each.value.description
  visibility_level = each.value.visibility
}
