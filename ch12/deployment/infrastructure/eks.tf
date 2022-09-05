variable eks_node_instance_types {
  description = "EC2 instance types to use for EKS nodes"
  type        = list(string)
  default     = ["t3.small"]
}

// https://registry.terraform.io/modules/terraform-aws-modules/eks/aws/18.29.0
module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "~> 18.29.0"

  manage_aws_auth_configmap = true

  aws_auth_users = [
    {
      userarn  = data.aws_caller_identity.current.arn
      username = "admin"
      groups   = ["system:masters"]
    }
  ]

  aws_auth_accounts = [
    data.aws_caller_identity.current.account_id
  ]

  cluster_name    = var.project
  cluster_version = "1.22"

  cluster_endpoint_private_access = true
  cluster_endpoint_public_access  = true

  cluster_endpoint_public_access_cidrs = [var.allowed_cidr_block]

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets

  enable_irsa = true

  eks_managed_node_group_defaults = {
    ami_type                              = "AL2_x86_64"
    disk_size                             = 20
    instance_types                        = var.eks_node_instance_types
    create_launch_template                = false
    launch_template_name                  = ""
    attach_cluster_primary_security_group = true
    iam_role_attach_cni_policy            = true
  }

  eks_managed_node_groups = {
    nodes = {
      name = "${var.project}-nodes"

      min_size     = 1
      max_size     = 3
      desired_size = 2
    }
  }
}

// https://registry.terraform.io/modules/terraform-aws-modules/iam/aws/5.3.1/submodules/iam-role-for-service-accounts-eks
module "load_balancer_controller_irsa_role" {
  source = "terraform-aws-modules/iam/aws//modules/iam-role-for-service-accounts-eks"
  version = "~> 5.3"

  role_name                              = "load-balancer-controller"
  attach_load_balancer_controller_policy = true

  oidc_providers = {
    main = {
      provider_arn               = module.eks.oidc_provider_arn
      namespace_service_accounts = ["kube-system:aws-load-balancer-controller"]
    }
  }
}

// https://registry.terraform.io/modules/terraform-aws-modules/iam/aws/5.3.1/submodules/iam-role-for-service-accounts-eks
module "vpc_cni_irsa" {
  source  = "terraform-aws-modules/iam/aws//modules/iam-role-for-service-accounts-eks"
  version = "~> 5.3"

  attach_vpc_cni_policy = true
  vpc_cni_enable_ipv4   = true

  oidc_providers = {
    main = {
      provider_arn               = module.eks.oidc_provider_arn
      namespace_service_accounts = ["kube-system:aws-node"]
    }
  }
}

output eks_cluster_id {
  description = "EKS cluster ID"
  value       = module.eks.cluster_id
}

output eks_endpoint {
  description = "Endpoint for EKS control plane"
  value       = module.eks.cluster_endpoint
}

output eks_certificate_authority_data {
  value = module.eks.cluster_certificate_authority_data
}
