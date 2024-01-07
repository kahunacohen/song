SHELL := /bin/bash
css: ./css/input.css
	./taillwindcss -i ./css/input.css -o static/output.css
	
