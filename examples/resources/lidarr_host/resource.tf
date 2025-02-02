resource "lidarr_host" "test" {
  launch_browser  = true
  port            = 8686
  url_base        = ""
  bind_address    = "*"
  application_url = ""
  instance_name   = "Lidarr"
  proxy = {
    enabled = false
  }
  ssl = {
    enabled                = false
    certificate_validation = "enabled"
  }
  logging = {
    log_level      = "info"
    log_size_limit = 1
  }
  backup = {
    folder    = "/backup"
    interval  = 5
    retention = 10
  }
  authentication = {
    method = "none"
  }
  update = {
    mechanism = "docker"
    branch    = "develop"
  }
}