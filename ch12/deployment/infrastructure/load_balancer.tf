// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/namespace_v1
resource kubernetes_namespace_v1 lb {
  metadata {
    name   = "ingress-nginx"
    labels = {
      "app.kubernetes.io/name"     = "ingress-nginx"
      "app.kubernetes.io/instance" = "ingress-nginx"
    }
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_account_v1
resource kubernetes_service_account_v1 lb {
  metadata {
    name      = "ingress-nginx"
    namespace = "ingress-nginx"
    labels    = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "controller"
    }
  }
  automount_service_account_token = true
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_account_v1
resource kubernetes_service_account_v1 lb_admission {
  metadata {
    name      = "ingress-nginx-admission"
    namespace = "ingress-nginx"
    labels    = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "admission-webhook"
    }
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/role_v1
resource kubernetes_role_v1 lb {
  metadata {
    name      = "ingress-nginx"
    namespace = "ingress-nginx"
    labels    = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "controller"
    }
  }

  rule {
    api_groups = [""]
    resources  = ["namespaces"]
    verbs      = ["get"]
  }

  rule {
    api_groups = [""]
    resources  = ["configmaps", "pods", "secrets", "endpoints"]
    verbs      = ["get", "list", "watch"]
  }

  rule {
    api_groups = [""]
    resources  = ["services"]
    verbs      = ["get", "list", "watch"]
  }

  rule {
    api_groups = ["networking.k8s.io"]
    resources  = ["ingresses"]
    verbs      = ["get", "list", "watch"]
  }

  rule {
    api_groups = ["networking.k8s.io"]
    resources  = ["ingresses/status"]
    verbs      = ["update"]
  }

  rule {
    api_groups = ["networking.k8s.io"]
    resources  = ["ingressclasses"]
    verbs      = ["get", "list", "watch"]
  }

  rule {
    api_groups     = [""]
    resource_names = ["ingress-controller-leader"]
    resources      = ["configmaps"]
    verbs          = ["get", "update"]
  }

  rule {
    api_groups = [""]
    resources  = ["configmaps"]
    verbs      = ["create"]
  }

  rule {
    api_groups     = ["coordination.k8s.io"]
    resource_names = ["ingress-controller-leader"]
    resources      = ["leases"]
    verbs          = ["get", "update"]
  }

  rule {
    api_groups = ["coordination.k8s.io"]
    resources  = ["leases"]
    verbs      = ["create"]
  }

  rule {
    api_groups = [""]
    resources  = ["events"]
    verbs      = ["create", "patch"]
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/role_v1
resource kubernetes_role_v1 lb_admission {
  metadata {
    name      = "ingress-nginx-admission"
    namespace = "ingress-nginx"
    labels    = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "admission-webhook"
    }
  }

  rule {
    api_groups = [""]
    resources  = ["secrets"]
    verbs      = ["get", "create"]
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/cluster_role_v1
resource kubernetes_cluster_role_v1 lb {
  metadata {
    name   = "ingress-nginx"
    labels = {
      "app.kubernetes.io/name"     = "ingress-nginx"
      "app.kubernetes.io/instance" = "ingress-nginx"
    }
  }

  rule {
    api_groups = [""]
    resources  = ["configmaps", "endpoints", "nodes", "pods", "secrets", "namespaces"]
    verbs      = ["list", "watch"]
  }

  rule {
    api_groups = ["coordination.k8s.io"]
    resources  = ["leases"]
    verbs      = ["list", "watch"]
  }

  rule {
    api_groups = [""]
    resources  = ["nodes"]
    verbs      = ["get"]
  }

  rule {
    api_groups = [""]
    resources  = ["services"]
    verbs      = ["get", "list", "watch"]
  }

  rule {
    api_groups = ["networking.k8s.io"]
    resources  = ["ingresses"]
    verbs      = ["get", "list", "watch"]
  }

  rule {
    api_groups = [""]
    resources  = ["events"]
    verbs      = ["create", "patch"]
  }

  rule {
    api_groups = ["networking.k8s.io"]
    resources  = ["ingresses/status"]
    verbs      = ["update"]
  }

  rule {
    api_groups = ["networking.k8s.io"]
    resources  = ["ingressclasses"]
    verbs      = ["get", "list", "watch"]
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/cluster_role_v1
resource kubernetes_cluster_role_v1 lb_admission {
  metadata {
    name   = "ingress-nginx-admission"
    labels = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "admission-webhook"
    }
  }

  rule {
    api_groups = ["admissionregistration.k8s.io"]
    resources  = ["validatingwebhookconfigurations"]
    verbs      = ["get", "update"]
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/role_binding_v1
resource kubernetes_role_binding_v1 lb {
  metadata {
    name      = "ingress-nginx"
    namespace = "ingress-nginx"
    labels    = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "controller"
    }
  }

  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "Role"
    name      = "ingress-nginx"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "ingress-nginx"
    namespace = "ingress-nginx"
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/role_binding_v1
resource kubernetes_role_binding_v1 lb_admission {
  metadata {
    name      = "ingress-nginx-admission"
    namespace = "ingress-nginx"
    labels    = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "admission-webhook"
    }
  }

  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "Role"
    name      = "ingress-nginx-admission"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "ingress-nginx-admission"
    namespace = "ingress-nginx"
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/cluster_role_binding_v1
resource kubernetes_cluster_role_binding_v1 lb {
  metadata {
    name   = "ingress-nginx"
    labels = {
      "app.kubernetes.io/name"     = "ingress-nginx"
      "app.kubernetes.io/instance" = "ingress-nginx"
    }
  }

  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "ingress-nginx"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "ingress-nginx"
    namespace = "ingress-nginx"
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/cluster_role_binding_v1
resource kubernetes_cluster_role_binding_v1 lb_admission {
  metadata {
    name   = "ingress-nginx-admission"
    labels = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "admission-webhook"
    }
  }

  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "ingress-nginx-admission"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "ingress-nginx-admission"
    namespace = "ingress-nginx"
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/config_map_v1
resource kubernetes_config_map_v1 lb {
  metadata {
    name      = "ingress-nginx-controller"
    namespace = "ingress-nginx"
    labels    = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "controller"
    }
  }
  data = {
    "allow-snippet-annotations" : "true"
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_v1
resource kubernetes_service_v1 lb_controller {
  metadata {
    name      = "ingress-nginx-controller"
    namespace = "ingress-nginx"
    labels    = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "controller"
    }
    annotations = {
      "service.beta.kubernetes.io/aws-load-balancer-backend-protocol"                  = "tcp"
      "service.beta.kubernetes.io/aws-load-balancer-cross-zone-load-balancing-enabled" = "true"
      "service.beta.kubernetes.io/aws-load-balancer-type"                              = "nlb"
    }
  }

  spec {
    selector = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "controller"
    }
    type                    = "LoadBalancer"
    external_traffic_policy = "Local"
    port {
      name        = "http"
      port        = 80
      target_port = "http"
      protocol    = "TCP"
    }
    port {
      name        = "https"
      port        = 443
      target_port = "https"
      protocol    = "TCP"
    }
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_v1
resource kubernetes_service_v1 lb_controller_admission {
  metadata {
    name      = "ingress-nginx-controller-admission"
    namespace = "ingress-nginx"
    labels    = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "controller"
    }
  }

  spec {
    selector = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "controller"
    }
    type = "ClusterIP"
    port {
      name        = "https-webhook"
      port        = 443
      target_port = "webhook"
    }
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 lb_controller {
  metadata {
    name      = "ingress-nginx-controller"
    namespace = "ingress-nginx"
    labels    = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "controller"
    }
  }
  spec {
    selector {
      match_labels = {
        "app.kubernetes.io/name"      = "ingress-nginx"
        "app.kubernetes.io/instance"  = "ingress-nginx"
        "app.kubernetes.io/component" = "controller"
      }
    }
    revision_history_limit = 10
    min_ready_seconds      = 0
    template {
      metadata {
        labels = {
          "app.kubernetes.io/name"      = "ingress-nginx"
          "app.kubernetes.io/instance"  = "ingress-nginx"
          "app.kubernetes.io/component" = "controller"
        }
      }
      spec {
        container {
          name              = "controller"
          image             = "registry.k8s.io/ingress-nginx/controller:v1.3.0@sha256:d1707ca76d3b044ab8a28277a2466a02100ee9f58a86af1535a3edf9323ea1b5"
          image_pull_policy = "IfNotPresent"
          lifecycle {
            pre_stop {
              exec {
                command = ["/wait-shutdown"]
              }
            }
          }
          args = [
            "/nginx-ingress-controller",
            "--publish-service=$(POD_NAMESPACE)/ingress-nginx-controller",
            "--election-id=ingress-controller-leader",
            "--controller-class=k8s.io/ingress-nginx",
            "--ingress-class=nginx",
            "--configmap=$(POD_NAMESPACE)/ingress-nginx-controller",
            "--validating-webhook=:8443",
            "--validating-webhook-certificate=/usr/local/certificates/cert",
            "--validating-webhook-key=/usr/local/certificates/key"
          ]
          security_context {
            capabilities {
              drop = ["ALL"]
              add  = ["NET_BIND_SERVICE"]
            }
            run_as_user                = 101
            allow_privilege_escalation = true
          }
          env {
            name = "POD_NAME"
            value_from {
              field_ref {
                field_path = "metadata.name"
              }
            }
          }
          env {
            name = "POD_NAMESPACE"
            value_from {
              field_ref {
                field_path = "metadata.namespace"
              }
            }
          }
          env {
            name  = "LD_PRELOAD"
            value = "/usr/local/lib/libmimalloc.so"
          }
          liveness_probe {
            http_get {
              path   = "/healthz"
              port   = 10254
              scheme = "HTTP"
            }
            initial_delay_seconds = 10
            period_seconds        = 10
            timeout_seconds       = 1
            success_threshold     = 1
            failure_threshold     = 3
          }
          readiness_probe {
            http_get {
              path   = "/healthz"
              port   = 10254
              scheme = "HTTP"
            }
            initial_delay_seconds = 10
            period_seconds        = 10
            timeout_seconds       = 1
            success_threshold     = 1
            failure_threshold     = 3
          }
          port {
            name           = "http"
            container_port = 80
            protocol       = "TCP"
          }
          port {
            name           = "https"
            container_port = 443
            protocol       = "TCP"
          }
          port {
            name           = "webhook"
            container_port = 8443
            protocol       = "TCP"
          }
          volume_mount {
            name       = "webhook-cert"
            mount_path = "/usr/local/certificates/"
            read_only  = true
          }
          resources {
            requests = {
              cpu    = "100m"
              memory = "90Mi"
            }
          }
        }
        dns_policy                       = "ClusterFirst"
        service_account_name             = "ingress-nginx"
        termination_grace_period_seconds = 300
        node_selector                    = {
          "kubernetes.io/os" = "linux"
        }
        volume {
          name = "webhook-cert"
          secret {
            secret_name = "ingress-nginx-admission"
          }
        }
      }
    }
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/job_v1
resource kubernetes_job_v1 lb_admission_create {
  metadata {
    name      = "ingress-nginx-admission-create"
    namespace = "ingress-nginx"
    labels    = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "admission-webhook"
    }
  }

  spec {
    template {
      metadata {
        name   = "ingress-nginx-admission-create"
        labels = {
          "app.kubernetes.io/name"      = "ingress-nginx"
          "app.kubernetes.io/instance"  = "ingress-nginx"
          "app.kubernetes.io/component" = "admission-webhook"
        }
      }
      spec {
        container {
          name              = "create"
          image             = "registry.k8s.io/ingress-nginx/kube-webhook-certgen:v1.1.1@sha256:64d8c73dca984af206adf9d6d7e46aa550362b1d7a01f3a0a91b20cc67868660"
          image_pull_policy = "IfNotPresent"
          args              = [
            "create",
            "--host=ingress-nginx-controller-admission,ingress-nginx-controller-admission.$(POD_NAMESPACE).svc",
            "--namespace=$(POD_NAMESPACE)",
            "--secret-name=ingress-nginx-admission"
          ]
          env {
            name = "POD_NAMESPACE"
            value_from {
              field_ref {
                field_path = "metadata.namespace"
              }
            }
          }
        }
        node_selector = {
          "kubernetes.io/os" = "linux"
        }
        restart_policy       = "OnFailure"
        service_account_name = "ingress-nginx-admission"
        security_context {
          run_as_non_root = true
          run_as_user     = 2000
          fs_group        = 2000
        }
      }
    }
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/job_v1
resource kubernetes_job_v1 lb_admission_patch {
  metadata {
    name      = "ingress-nginx-admission-patch"
    namespace = "ingress-nginx"
    labels    = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "admission-webhook"
    }
  }

  spec {
    template {
      metadata {
        name   = "ingress-nginx-admission-patch"
        labels = {
          "app.kubernetes.io/name"      = "ingress-nginx"
          "app.kubernetes.io/instance"  = "ingress-nginx"
          "app.kubernetes.io/component" = "admission-webhook"
        }
      }
      spec {
        container {
          name              = "patch"
          image             = "registry.k8s.io/ingress-nginx/kube-webhook-certgen:v1.1.1@sha256:64d8c73dca984af206adf9d6d7e46aa550362b1d7a01f3a0a91b20cc67868660"
          image_pull_policy = "IfNotPresent"
          args              = [
            "patch",
            "--webhook-name=ingress-nginx-admission",
            "--namespace=$(POD_NAMESPACE)",
            "--patch-mutating=false",
            "--secret-name=ingress-nginx-admission",
            "--patch-failure-policy=Fail"
          ]
          env {
            name = "POD_NAMESPACE"
            value_from {
              field_ref {
                field_path = "metadata.namespace"
              }
            }
          }
          security_context {
            allow_privilege_escalation = false
          }
        }
        node_selector = {
          "kubernetes.io/os" = "linux"
        }
        restart_policy       = "OnFailure"
        service_account_name = "ingress-nginx-admission"
        security_context {
          run_as_non_root = true
          run_as_user     = 2000
          fs_group        = 2000
        }
      }
    }
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/ingress_class_v1
resource kubernetes_ingress_class_v1 lb_ingress {
  metadata {
    name   = "nginx"
    labels = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "controller"
    }
  }
  spec {
    controller = "k8s.io/ingress-nginx"
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/validating_webhook_configuration_v1
resource kubernetes_validating_webhook_configuration_v1 lb_admission {
  metadata {
    name   = "ingress-nginx-admission"
    labels = {
      "app.kubernetes.io/name"      = "ingress-nginx"
      "app.kubernetes.io/instance"  = "ingress-nginx"
      "app.kubernetes.io/component" = "admission-webhook"
    }
  }
  webhook {
    name = "validate.nginx.ingress.kubernetes.io"
    rule {
      api_groups   = ["networking.k8s.io"]
      api_versions = ["v1"]
      operations   = ["CREATE", "UPDATE"]
      resources    = ["ingresses"]
    }
    side_effects              = "None"
    admission_review_versions = ["v1"]
    failure_policy            = "Ignore"
    client_config {
      service {
        name      = "ingress-nginx-controller-admission"
        namespace = "ingress-nginx"
        path      = "/extensions/v1/ingresses"
      }
      ca_bundle = <<EOF
-----BEGIN CERTIFICATE-----
MIIBdDCCARugAwIBAgIQZAo33gBADQ0uQqCR4vFUljAKBggqhkjOPQQDAjAPMQ0w
CwYDVQQKEwRuaWwxMCAXDTIyMDkwNTA2Mzg1MVoYDzIxMjIwODEyMDYzODUxWjAP
MQ0wCwYDVQQKEwRuaWwxMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE565nmZLZ
CO8yZD1fxBZrR2gDj3XLK7az3L+g8/MUe+8U+8a+VgrCB1NgjjrPY23RGPQaFXsS
hekeMWWS2495cKNXMFUwDgYDVR0PAQH/BAQDAgIEMBMGA1UdJQQMMAoGCCsGAQUF
BwMBMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFMxaSreFGm62s3S8ukuatyu/
d/U1MAoGCCqGSM49BAMCA0cAMEQCICQXxNlEUzGJ8i8Qrzlh1cd3ZtQxY5X5k+bN
or6i3bBeAiAWFdCGdwKSn0UFcgeBDz0kT//cpWmGPF++r3b6j2scIg==
-----END CERTIFICATE-----
EOF
    }
  }
}
