CONN = $(shell terraform output -raw db_conn)

deploy: init plan apply

init:
	terraform init -upgrade

plan:
	terraform plan -out mallbots.tfplan

apply:
	terraform apply mallbots.tfplan

destroy:
	terraform destroy

setup-db:
	psql --file="./sql/create_commondb.sql" ${CONN}/postgres
	psql --file="./sql/init_commondb.sql" ${CONN}/commondb

	psql --file="./sql/create_baskets.sql" ${CONN}/postgres
	psql --file="./sql/init_baskets.sql" ${CONN}/baskets
