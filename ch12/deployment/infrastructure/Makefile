CONN = $(shell terraform output -raw db_conn)

# Use this command to easily run all of the infrastructure deployment steps
# There are costs associated with this demo; Remember to tear everything down!
deploy: init validate plan apply setup-db

init:
	terraform init -upgrade

validate:
	terraform validate

plan:
	terraform plan -out infra.tfplan

apply:
	terraform apply infra.tfplan
	@echo !!!
	@echo !!!
	@echo !!! There are cost associated with running this demo! Remember to tear everything down!
	@echo !!!
	@echo !!! Use 'make destroy' OR use 'terraform destroy' to tear everything down.
	@echo !!!
	@echo !!!

destroy:
	terraform destroy

# Use this command to populate ~/.kube/config so that you can use k9s to browse the cluster
kubeconfig:
	aws eks update-kubeconfig --region $(shell terraform output region) --name $(shell terraform output project)


setup-db:
	@echo "Initializing the Mallbots service DBs"
	@psql --file="./sql/create_commondb.sql" ${CONN}/postgres
	@psql --file="./sql/init_commondb.sql" ${CONN}/commondb
	@psql --file="./sql/create_services.sql" ${CONN}/postgres

	@psql --file="./sql/init_baskets.sql" ${CONN}/baskets
	@psql --file="./sql/init_cosec.sql" ${CONN}/cosec
	@psql --file="./sql/init_customers.sql" ${CONN}/customers
	@psql --file="./sql/init_depot.sql" ${CONN}/depot
	@psql --file="./sql/init_notifications.sql" ${CONN}/notifications
	@psql --file="./sql/init_ordering.sql" ${CONN}/ordering
	@psql --file="./sql/init_payments.sql" ${CONN}/payments
	@psql --file="./sql/init_search.sql" ${CONN}/search
	@psql --file="./sql/init_stores.sql" ${CONN}/stores
