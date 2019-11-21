BIN = test_basic #test_batch_write ingest_data
CXXFLAGS = -I ../rocksdb/include/ -L ../rocksdb -lrocksdb -lpthread -lz -lsnappy -lboost_program_options -lzstd -llz4 -lbz2 -std=c++11

ifeq ($(MAKECMDGOALS), debug)
    CXXFLAGS += -g -ggdb
else
    CXXFLAGS += -O3
endif

all: release

debug: $(BIN)

release: $(BIN)

test_basic: test_basic.cpp common.h
	g++ -o $@ $< $(CXXFLAGS)

#test_batch_write: test_batch_write.cpp common.h
#	g++ -o $@ $< $(CXXFLAGS)

#ingest_data: ingest_data.cpp common.h
#	g++ -o $@ $< $(CXXFLAGS)

clean:
	rm -f $(BIN)
