variable services {
  description = "List of MallBots microservices"
  type        = list(string)
  default     = ["baskets", "cosec", "customers", "depot", "ordering", "notifications", "payments", "search", "stores"]
}

// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity
data aws_caller_identity current {}

locals {
  region         = data.terraform_remote_state.infra.outputs.region
  aws_ecr_url    = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${data.terraform_remote_state.infra.outputs.region}.amazonaws.com"
  eks_cluster_id = data.terraform_remote_state.infra.outputs.eks_cluster_id
  project        = data.terraform_remote_state.infra.outputs.project
  db_endpoint    = data.terraform_remote_state.infra.outputs.db_endpoint
  db_host        = data.terraform_remote_state.infra.outputs.db_host
  db_port        = data.terraform_remote_state.infra.outputs.db_port
}
