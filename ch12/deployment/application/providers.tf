terraform {
  required_version = "~> 1.2.0"

  backend local {
    path = "./application.tfstate"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.29.0"
    }

    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.20.0"
    }

    random = {
      source  = "hashicorp/random"
      version = "~> 3.1.0"
    }
  }
}

data terraform_remote_state infra {
  backend = "local"
  config = {
    path = "${path.module}/../infrastructure/infrastructure.tfstate"
  }
}

// https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs
provider docker {
  registry_auth {
    address  = local.aws_ecr_url
    username = data.aws_ecr_authorization_token.token.user_name
    password = data.aws_ecr_authorization_token.token.password
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs
provider kubernetes {
  host                   = data.terraform_remote_state.infra.outputs.eks_endpoint
  cluster_ca_certificate = base64decode(data.terraform_remote_state.infra.outputs.eks_certificate_authority_data)
  exec {
    api_version = "client.authentication.k8s.io/v1beta1"
    command     = "aws"
    # This requires the awscli to be installed locally where Terraform is executed
    args = ["eks", "get-token", "--region", local.region, "--cluster-name", local.eks_cluster_id]
  }
}