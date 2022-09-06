// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/secret_v1
resource kubernetes_secret_v1 notifications {
  metadata {
    name      = "notifications-secrets"
    namespace = local.project
  }

  data = {
    PG_CONN = "host=${local.db_host} port=${local.db_port} dbname=notifications user=notifications_user password=notifications_pass search_path=notifications,public"
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 notifications {
  metadata {
    name      = "notifications"
    namespace = local.project
    labels    = {
      app = "notifications"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name" = "notifications"
      }
    }
    template {
      metadata {
        name   = "notifications"
        labels = {
          "app.kubernetes.io/name" = "notifications"
        }
      }
      spec {
        hostname = "notifications"
        container {
          name              = "notifications"
          image             = "${local.aws_ecr_url}/notifications:latest"
          image_pull_policy = "Always"
          env_from {
            config_map_ref {
              name = "common-config-map"
            }
          }
          env_from {
            secret_ref {
              name = "notifications-secrets"
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

  depends_on = [
    kubernetes_namespace_v1.namespace,
    kubernetes_config_map_v1.common,
    kubernetes_secret_v1.notifications
  ]
}

#// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_v1
#resource kubernetes_service_v1 notifications {
#  metadata {
#    name      = "notifications"
#    namespace = local.project
#    labels    = {
#      app = "notifications"
#    }
#  }
#  spec {
#    selector = {
#      "app.kubernetes.io/name" = "notifications"
#    }
#    session_affinity = "ClientIP"
#    port {
#      name        = "http"
#      protocol    = "TCP"
#      port        = 80
#      target_port = 80
#    }
#    port {
#      name        = "grpc"
#      protocol    = "TCP"
#      port        = 9000
#      target_port = 9000
#    }
#    type = "NodePort"
#  }
#}
#
#// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/ingress_v1
#resource kubernetes_ingress_v1 notifications {
#  metadata {
#    name = "notifications-ingress"
#    namespace = local.project
#  }
#
#  spec {
#    rule {
#      http {
#        path {
#          path = "/api/notifications/"
#          path_type = "Prefix"
#          backend {
#            service {
#              name = "notifications"
#              port {
#                number = 80
#              }
#            }
#          }
#        }
#      }
#    }
#    ingress_class_name = "nginx"
#  }
#}
