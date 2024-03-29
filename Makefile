SRC_DIR=./messages
DST_DIR=./pkg
DST_DIR2=./python/protorepo

cleanup:
	rm -rf $(DST_DIR)/*
	rm -rf $(DST_DIR2)/*

format:
	go fmt $(DST_DIR)/**

build:
	make cleanup
	protoc -I=$(SRC_DIR) --go_out=$(DST_DIR) --python_out=$(DST_DIR2) $(SRC_DIR)/messages.proto $(SRC_DIR)/hashfiles.proto --mypy_out=$(DST_DIR2)
	touch $(DST_DIR2)/__init__.py
	make format