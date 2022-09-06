terraform {
  required_version = "~> 1.2.0"

  backend local {
    path = "./infrastructure.tfstate"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.29.0"
    }

    random = {
      source  = "hashicorp/random"
      version = "~> 3.1.0"
    }
  }
}

// https://registry.terraform.io/providers/hashicorp/aws/latest/docs
provider aws {
  region = var.region

  default_tags {
    tags = {
      Application = "MallBots"
    }
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs
provider kubernetes {
  host                   = module.eks.cluster_endpoint
  cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)
  exec {
    api_version = "client.authentication.k8s.io/v1beta1"
    command     = "aws"
    # This requires the awscli to be installed locally where Terraform is executed
    args = ["eks", "get-token", "--region", var.region, "--cluster-name", module.eks.cluster_id]
  }
}
