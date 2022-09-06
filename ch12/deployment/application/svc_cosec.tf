// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/secret_v1
resource kubernetes_secret_v1 cosec {
  metadata {
    name      = "cosec-secrets"
    namespace = local.project
  }

  data = {
    PG_CONN = "host=${local.db_host} port=${local.db_port} dbname=cosec user=cosec_user password=cosec_pass search_path=cosec,public"
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 cosec {
  metadata {
    name      = "cosec"
    namespace = local.project
    labels    = {
      app = "cosec"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name" = "cosec"
      }
    }
    template {
      metadata {
        name   = "cosec"
        labels = {
          "app.kubernetes.io/name" = "cosec"
        }
      }
      spec {
        hostname = "cosec"
        container {
          name              = "cosec"
          image             = "${aws_ecr_repository.services["cosec"].repository_url}:latest"
          image_pull_policy = "Always"
          env_from {
            config_map_ref {
              name = "common-config-map"
            }
          }
          env_from {
            secret_ref {
              name = "cosec-secrets"
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

#// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_v1
#resource kubernetes_service_v1 cosec {
#  metadata {
#    name      = "cosec"
#    namespace = local.project
#    labels    = {
#      app = "cosec"
#    }
#  }
#  spec {
#    selector = {
#      "app.kubernetes.io/name" = "cosec"
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
#resource kubernetes_ingress_v1 cosec {
#  metadata {
#    name = "cosec-ingress"
#    namespace = local.project
#  }
#
#  spec {
#    rule {
#      http {
#        path {
#          path = "/api/cosec/"
#          path_type = "Prefix"
#          backend {
#            service {
#              name = "cosec"
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
