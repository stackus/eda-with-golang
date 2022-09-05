variable db_instance_type {
  description = "RDS instance type"
  type        = string
  default     = "db.t4g.micro"
}

variable db_username {
  description = "User name for the RDS PostgreSQL database"
  type        = string
}

// https://registry.terraform.io/modules/terraform-aws-modules/rds/aws/5.0.3
module "db" {
  source  = "terraform-aws-modules/rds/aws"
  version = "~> 5.0.0"

  identifier = "${var.project}-db"

  instance_class = var.db_instance_type
  engine         = "postgres"
  engine_version = "14.4"
  family         = "postgres14"

  allocated_storage = 5

  username = var.db_username
  port     = 5432

  multi_az               = true
  db_subnet_group_name   = module.vpc.database_subnet_group
  vpc_security_group_ids = [module.security_group.security_group_id]
  publicly_accessible    = true
}

output db_endpoint {
  value = module.db.db_instance_endpoint
}

output db_host {
  value = module.db.db_instance_address
}

output db_port {
  value = module.db.db_instance_port
}

output db_username {
  value = module.db.db_instance_username
  sensitive = true
}

output db_password {
  value     = module.db.db_instance_password
  sensitive = true
}

output db_conn {
  value = "postgres://${module.db.db_instance_username}:${module.db.db_instance_password}@${module.db.db_instance_endpoint}"
  sensitive = true
}
