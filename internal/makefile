.PHONY: all air go_watch tailwind_watch run

all: templ_watch tailwind_watch air 

templ_watch:
	@templ generate --watch --proxy="http://localhost:8080" --cmd="go run ." &

tailwind_watch:
	@echo "Starting Tailwind CSS watch..."
	@npm run watch &

air:
	@echo "Starting air..."
	@air &

