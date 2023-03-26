CC=gcc
CFLAGS=-Wall -fPIC -O3 -I$(INCLUDE_DIR)
LDFLAGS=-shared -L$(LIB_DIR)
LIBS_FOLDER=libs
INCLUDE_DIR=$(LIBS_FOLDER)/include
LIB_DIR=$(LIBS_FOLDER)/lib
OBJ_DIR=obj

SOURCES=sha256.c yespower-opt.c yespower.c
OBJECTS=$(SOURCES:%.c=$(OBJ_DIR)/%.o)
TARGET_LIB=$(LIB_DIR)/libyespower.so

HEADERS_FOLDER := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
HEADERS := $(wildcard $(HEADERS_FOLDER)*.h)
HEADERS_TO_COPY=$(addprefix $(INCLUDE_DIR)/, $(notdir $(HEADERS)))

.PHONY: all clean

all: $(HEADERS_TO_COPY) $(TARGET_LIB)

$(TARGET_LIB): $(OBJECTS) | $(LIB_DIR)
	$(CC) $(LDFLAGS) -o $@ $^

$(OBJ_DIR)/%.o: %.c | $(OBJ_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

$(HEADERS_TO_COPY): | $(INCLUDE_DIR)

$(INCLUDE_DIR)/%.h: $(HEADERS_FOLDER)%.h | $(INCLUDE_DIR)
	cp $< $@

$(OBJ_DIR) $(LIB_DIR) $(INCLUDE_DIR):
	mkdir -p $@

clean:
	rm -rf $(OBJ_DIR) $(LIB_DIR) $(INCLUDE_DIR)

