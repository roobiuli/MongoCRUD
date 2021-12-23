

resource "kubernetes_deployment" "commentsservicedeployment" {
  metadata {
    name = "commentsservice"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "comservice"
      }
    }
    template {
      metadata {
        labels = {
          app = "comservice"
        }
      }
      spec {
        container {
          image = "localhost:5000/comments:latest"
          name = "comments"
        }
      }
    }
  }
}


resource "kubernetes_service" "commentserviceserv" {
  metadata {
    name = "commentservice"
    namespace = kubernetes_namespace.comments.metadata[0].name
  }
  spec {
    selector = {
      app = kubernetes_deployment.MongoDB.spec[0].template[0].metadata[0].name
    }
    type = "NodePort"
    port {
      port = 8081
      node_port = 8081
    }
  }
}