#resource random_password services {
#  for_each = toset(local.services)
#
#  length = 16
#}

resource null_resource init_db {
  provisioner "local-exec" {
    command = "psql --file sql/init_db.psql ${local.db_conn}/postgres"
  }
}

#resource null_resource init_service_dbs {
#  for_each = toset(local.services)
#  provisioner "local-exec" {
#    command = "psql --file sql/init_service_db.psql -v db=$DB -v user=$USER -v pass=$PASS ${local.db_conn}/postgres"
#    environment = {
#      DB = each.key
#      USER = "${each.key}_user"
#      PASS = random_password.services[each.key].result
#    }
#  }
#  depends_on = [
#    null_resource.init_db,
#    random_password.services
#  ]
#}
