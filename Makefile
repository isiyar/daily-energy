DC = sudo docker compose

dev: start_all_dev logs_all_dev

start_all_dev:
	$(DC) --profile dev up --build

drop_all_dev:
	$(DC) --profile dev down -v

logs_all_dev:
	$(DC) --profile dev logs -f

prod: start_all_prod logs_all_prod

start_all_prod:
	$(DC) --profile prod up --build

drop_all_prod:
	$(DC) --profile prod down -v

logs_all_prod:
	$(DC) --profile prod logs -f
