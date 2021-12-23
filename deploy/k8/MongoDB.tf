resource "kubernetes_namespace" "comments" {
  metadata {
    name = "mongodbcom"
  }
}


resource "kubernetes_deployment" "MongoDB" {
  metadata {
    name = "mongodb"
    namespace = kubernetes_namespace.comments.metadata[0].name
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "mongodb"
      }
    }
    template {
      metadata {
        labels = {
          app = "mongodb"
        }
      }
      spec {
        container {
          image = "mongo:4.0"
          name = "mongodb"
          port {
            container_port = 27017
          }

          // Args for the img

//          args = ["--dbpath", "/data/db"]

          command = [<<EOF
            mongo -u $(MONGO_INITDB_ROOT_USERNAME) \
            -p $(MONGO_INITDB_ROOT_PASSWORD \
            --authenticationDatabase admin $(MONGO_NAME) \
            db.createUser({
    user: "$(MONGO_USER)",
    pwd: "$(MONGO_PASS)",
    roles: [{
        role: "readWrite",
        db: "$(MONGO_NAME)"
    }],
    mechanisms: ["$(MONGO_AUTH)"],
}) \
use $(MONGO_NAME)
db.createCollection("Comments")
db.comments.createIndex({"uuid":1},{unique:true,name:"UQ_uuid"})
EOF
          , "mongo", "--dbpath", "/data/db"]


          env {
            name = "MONGO_INITDB_ROOT_USERNAME"
            value = "root"
          }
          env {
            name = "MONGO_INITDB_ROOT_PASSWORD"
            value = "rootpass"
          }
          env {
            name = "MONGO_NAME"
            value = "blog"
          }
          env {
            name = "MONGO_USER"
            value = "Muser"
          }
          env {
            name = "MONGO_PASS"
            value = "1234567"
          }
          env {
            name = "MONGO_AUTH"
            value = "SCRAM-SHA-256"
          }
          volume_mount {
            mount_path = "/data/db"
            name = kubernetes_persistent_volume.dbstor.metadata[0].name
          }
        }
        volume {
          name = "Mongo-data-dir"
          persistent_volume_claim {
            claim_name = kubernetes_persistent_volume_claim.dbstoreclaim.metadata[0].name
          }
        }
      }
    }
  }
}


resource "kubernetes_service" "MongoDB" {
  metadata {
    name = "mongodb"
    namespace = kubernetes_namespace.comments.metadata[0].name
  }
  spec {
    selector = {
      app = kubernetes_deployment.MongoDB.spec[0].template[0].metadata[0].labels.app
    }
    //type = "NodePort"
    port {
      protocol = "TCP"
      port = 27017
      target_port = 27017
    }
  }

}