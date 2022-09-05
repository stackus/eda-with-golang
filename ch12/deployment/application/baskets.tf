#// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/config_map_v1
#resource kubernetes_config_map_v1 baskets {
#  metadata {
#    name      = "baskets-config-map"
#    namespace = local.project
#  }
#
#  data = {
#    WEB_PORT = ":80"
#    RPC_PORT = ":9000"
#  }
#}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/secret_v1
resource kubernetes_secret_v1 baskets {
  metadata {
    name      = "baskets-secrets"
    namespace = local.project
  }

  data = {
    PG_CONN = "host=${local.db_host} port=${local.db_port} dbname=baskets user=baskets_user password=baskets_pass search_path=baskets,public"
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 baskets {
  metadata {
    name      = "baskets"
    namespace = local.project
    labels    = {
      app = "baskets"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name" = "baskets"
      }
    }
    template {
      metadata {
        name   = "baskets"
        labels = {
          "app.kubernetes.io/name" = "baskets"
        }
      }
      spec {
        hostname = "baskets"
        container {
          name              = "baskets"
          image             = "${aws_ecr_repository.services["baskets"].repository_url}:latest"
          image_pull_policy = "Always"
          env_from {
            config_map_ref {
              name = "common-config-map"
            }
          }
#          env_from {
#            config_map_ref {
#              name = "baskets-config-map"
#            }
#          }
          env_from {
            secret_ref {
              name = "baskets-secrets"
            }
          }
          port {
            protocol       = "TCP"
            container_port = 80
          }
          port {
            protocol       = "TCP"
            container_port = 9000
          }
          liveness_probe {
            http_get {
              path = "/liveness"
              port = 80
            }
            initial_delay_seconds = 3
            period_seconds        = 5
          }
        }
      }
    }
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_v1
resource kubernetes_service_v1 baskets {
  metadata {
    name      = "baskets"
    namespace = local.project
    labels    = {
      app = "baskets"
    }
  }
  spec {
    selector = {
      "app.kubernetes.io/name" = "baskets"
    }
    session_affinity = "ClientIP"
    port {
      name        = "http"
      protocol    = "TCP"
      port        = 80
      target_port = 80
    }
    port {
      name        = "grpc"
      protocol    = "TCP"
      port        = 9000
      target_port = 9000
    }
    type = "NodePort"
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/ingress_v1
resource kubernetes_ingress_v1 baskets {
  metadata {
    name = "baskets-ingress"
    namespace = local.project
  }

  spec {
    rule {
      http {
        path {
          path = "/api/baskets/"
          path_type = "Prefix"
          backend {
            service {
              name = "baskets"
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
