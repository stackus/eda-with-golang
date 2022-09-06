// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/secret_v1
resource kubernetes_secret_v1 ordering {
  metadata {
    name      = "ordering-secrets"
    namespace = local.project
  }

  data = {
    PG_CONN = "host=${local.db_host} port=${local.db_port} dbname=ordering user=ordering_user password=ordering_pass search_path=ordering,public"
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 ordering {
  metadata {
    name      = "ordering"
    namespace = local.project
    labels    = {
      app = "ordering"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name" = "ordering"
      }
    }
    template {
      metadata {
        name   = "ordering"
        labels = {
          "app.kubernetes.io/name" = "ordering"
        }
      }
      spec {
        hostname = "ordering"
        container {
          name              = "ordering"
          image             = "${aws_ecr_repository.services["ordering"].repository_url}:latest"
          image_pull_policy = "Always"
          env_from {
            config_map_ref {
              name = "common-config-map"
            }
          }
          env_from {
            secret_ref {
              name = "ordering-secrets"
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
resource kubernetes_service_v1 ordering {
  metadata {
    name      = "ordering"
    namespace = local.project
    labels    = {
      app = "ordering"
    }
  }
  spec {
    selector = {
      "app.kubernetes.io/name" = "ordering"
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
resource kubernetes_ingress_v1 ordering {
  metadata {
    name = "ordering-ingress"
    namespace = local.project
    annotations = {
      "nginx.ingress.kubernetes.io/whitelist-source-range" = local.allowed_cidr_block
    }
  }

  spec {
    rule {
      http {
        path {
          path = "/api/ordering"
          path_type = "Prefix"
          backend {
            service {
              name = "ordering"
              port {
                number = 80
              }
            }
          }
        }
        path {
          path = "/ordering-spec/"
          path_type = "Prefix"
          backend {
            service {
              name = "ordering"
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
