PATH := ./:$(PATH)

css: ./css/input.css
	./taillwindcss -i ./css/input.css -o static/output.css
