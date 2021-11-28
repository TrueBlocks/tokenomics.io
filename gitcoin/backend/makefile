#-------------------------------------------------
exec=pouch

#-------------------------------------------------
src= \
./pouch.cpp \
./options.cpp \
./handle_summarize.cpp \
./handle_audit.cpp \
./handle_json2csv.cpp \
./handle_csv2json.cpp \
./classes/logentry_min.cpp \
./classes/donation.cpp \
./classes/balance.cpp \
./classes/record.cpp \
./classes/payout.cpp

headers= \
./options.h \
./classes/balance.h \
./classes/record.h \
./classes/logentry_min.h \
./classes/payout.h \
./classes/donation.h

#-------------------------------------------------
# probably don't have to change below this line

#-------------------------------------------------
incs=/usr/local/include/tb
cflags=-std=c++11 -Werror -Wall -O2 -I. -I$(incs)/cachelib -I$(incs)/pinlib -I$(incs)/acctlib -I$(incs)/etherlib -I$(incs)/utillib

#-------------------------------------------------
libs=/usr/local/lib/tb/*.a
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
    libs += -lcurl
endif
ifeq ($(UNAME_S),Darwin)
    libs += /Library/Developer/CommandLineTools/SDKs/MacOSX11.1.sdk/usr/lib/libcurl.4.tbd
endif

#-------------------------------------------------
product=objs/$(exec)

#-------------------------------------------------
all:
	@make app

app:
	@make $(product)
	mv $(product) ./bin/

objects = $(patsubst %.cpp,objs/%.o,$(src))

$(product): $(objects) $(libs)
	@echo $(objects)
	g++ -o $(product) $(objects) $(libs)
	@strip $(product)

$(objects): | objs

objs:
	@mkdir -p ./bin ./objs ./objs/classes

objs/%.o : %.cpp
	$(CXX) $(CXXFLAGS) $(cflags) -c $< -o $@

clean:
	-@$(RM) -fR ./bin objs 2> /dev/null
