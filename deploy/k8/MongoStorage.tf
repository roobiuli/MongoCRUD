resource "kubernetes_persistent_volume" "dbstor" {
  metadata {
    name = "dbstorage"
  }
  spec {
    access_modes = ["ReadWriteOnce"]
    capacity = {
      storage = "500M"
    }
    persistent_volume_source {
      host_path {
        path = "/data/mongo"
      }
    }
  }
}

resource "kubernetes_persistent_volume_claim" "dbstoreclaim" {
  metadata {
    name = "dbmongopvc"
  }
  spec {
    storage_class_name = ""
    access_modes = ["ReadWriteOnce"]
    resources {
      requests = {
        storage = "500M"
      }
    }
  }
}