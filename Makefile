.PHONY: build push

IMAGE = windsource/newsapp
VERSION = 1.2.0

build:
	docker build -t $(IMAGE):$(VERSION) -t $(IMAGE):latest .

push: build
	docker push $(IMAGE):$(VERSION)
	docker push $(IMAGE):latest