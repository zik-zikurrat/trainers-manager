terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 4.2.0"
    }
  }
}

provider "docker" {}


resource "docker_network" "training_net" {
  name = "training_net"
}

resource "docker_volume" "db_data" {
  name = "db_training_vol"
}

resource "docker_image" "postgres" {
  name = "postgres:14-alpine"
}

resource "docker_container" "postgres" {
  name  = "db_training"
  image = docker_image.postgres.image_id
  env = [
    "POSTGRES_USER=postgres",
    "POSTGRES_PASSWORD=postgres",
    "POSTGRES_DB=training",
  ]
  networks_advanced {
    name = docker_network.training_net.name
  }
  volumes {
    volume_name    = docker_volume.db_data.name
    container_path = "/var/lib/postgresql/data"
  }

  restart = "unless-stopped"

  healthcheck {
    test     = ["CMD-SHELL", "pg_isready -U postgres -d training"]
    interval = "5s"
    timeout  = "3s"
    retries  = 5
  }
}

resource "docker_image" "training" {
  name = "trainers-manager-training:latest"
  build {
    context    = "${path.module}/trainers-manager"
    dockerfile = "Dockerfile"
  }
}
resource "docker_container" "training" {
  name  = "training"
  image = docker_image.training.image_id

  env = [
    "CONFIG_PATH=config/docker.yaml",
    "LLM_API_KEY=${var.llm_api_key}",
  ]
  networks_advanced {
    name = docker_network.training_net.name
  }
  ports {
    internal = 9045
    external = 9045
  }
  restart = "unless-stopped"

  depends_on = [docker_container.postgres]
}

resource "docker_image" "frontend" {
  name = "trainers-frontend:latest"
  build {
    context    = "${path.module}/trainers-frontend"
    dockerfile = "Dockerfile"
  }
}

resource "docker_container" "frontend" {
  name  = "frontend"
  image = docker_image.frontend.image_id
  networks_advanced {
    name = docker_network.training_net.name
  }

  ports {
    internal = 80
    external = 80
  }

  restart    = "unless-stopped"
  depends_on = [docker_container.training]
}

variable "llm_api_key" {
  type      = string
  sensitive = true
}
