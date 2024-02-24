#!/bin/sh

./migrate -source file://migrations -database postgresql://postgres:password@postgres:5432/songs?sslmode=disable up