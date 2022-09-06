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

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/ingress_v1
resource kubernetes_ingress_v1 swagger {
  metadata {
    name = "swagger-ingress"
    namespace = local.project
    annotations = {
      "nginx.ingress.kubernetes.io/whitelist-source-range" = local.allowed_cidr_block
    }
  }

  spec {
    rule {
      http {
        path {
          path = "/"
          path_type = "Prefix"
          backend {
            service {
              name = "baskets" # pick a service; any service
              port {
                number = 80
              }
            }
          }
        }
      }
    }
    ingress_class_name = "nginx"
  }
}
