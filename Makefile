SOURCE := $(shell git rev-parse --show-toplevel)

include $(SOURCE)/script/make/dev.mk
include $(SOURCE)/script/make/build.mk
