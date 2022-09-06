// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/secret_v1
resource kubernetes_secret_v1 stores {
  metadata {
    name      = "stores-secrets"
    namespace = local.project
  }

  data = {
    PG_CONN = "host=${local.db_host} port=${local.db_port} dbname=stores user=stores_user password=stores_pass search_path=stores,public"
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 stores {
  metadata {
    name      = "stores"
    namespace = local.project
    labels    = {
      app = "stores"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name" = "stores"
      }
    }
    template {
      metadata {
        name   = "stores"
        labels = {
          "app.kubernetes.io/name" = "stores"
        }
      }
      spec {
        hostname = "stores"
        container {
          name              = "stores"
          image             = "${aws_ecr_repository.services["stores"].repository_url}:latest"
          image_pull_policy = "Always"
          env_from {
            config_map_ref {
              name = "common-config-map"
            }
          }
          env_from {
            secret_ref {
              name = "stores-secrets"
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
resource kubernetes_service_v1 stores {
  metadata {
    name      = "stores"
    namespace = local.project
    labels    = {
      app = "stores"
    }
  }
  spec {
    selector = {
      "app.kubernetes.io/name" = "stores"
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
resource kubernetes_ingress_v1 stores {
  metadata {
    name = "stores-ingress"
    namespace = local.project
    annotations = {
      "nginx.ingress.kubernetes.io/whitelist-source-range" = local.allowed_cidr_block
    }
  }

  spec {
    rule {
      http {
        path {
          path = "/api/stores"
          path_type = "Prefix"
          backend {
            service {
              name = "stores"
              port {
                number = 80
              }
            }
          }
        }
        path {
          path = "/stores-spec/"
          path_type = "Prefix"
          backend {
            service {
              name = "stores"
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
