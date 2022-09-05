// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/namespace
resource kubernetes_namespace mallbots {
  metadata {
    name = local.project
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/config_map_v1
resource kubernetes_config_map_v1 common {
  metadata {
    name      = "common-config-map"
    namespace = local.project
  }

  data = {
    ENVIRONMENT  = "production"
    WEB_PORT     = ":80"
    RPC_PORT     = ":9000"
    NATS_URL     = "nats:4222"
    RPC_SERVICES = "STORES=stores:9000,CUSTOMERS=customers:9000"
  }
}

