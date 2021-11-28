#-------------------------------------------------
libraries=$(HOME)/Development/trueblocks-core/src/libs/
cflags=-std=c++11 -Werror -Wall -O2 -I. -I$(libraries)/utillib -I$(libraries)/etherlib -I$(libraries)/acctlib
bin=bin

#-------------------------------------------------
#curl=/Library/Developer/CommandLineTools/SDKs/MacOSX11.1.sdk/usr/lib/libcurl.4.tbd
#libs=$(libraries)/*.a $(curl)
libs=$(libraries)/*.a -lcurl

#-------------------------------------------------
exec=pouch
product=objs/$(exec)
dest=$(bin)/$(exec)

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
all:
	@mkdir -p objs/classes
	@make app

every:
	@cd ../../../build ; make ; cd -
	@make app

app:
	@mkdir -p bin
	@make $(product)
	mv $(product) $(dest)

test:
	@make
	@echo "No tests"

#-------------------------------------------------
# probably don't have to change below this line

objects = $(patsubst %.cpp,objs/%.o,$(src))

$(product): $(objects) $(libs)
	g++ -o $(product) $(objects) $(libs)
	@strip $(product)

$(objects): | objs

objs:
	@mkdir -p $@

objs/%.o : %.cpp $(headers)
	$(CXX) $(CXXFLAGS) $(cflags) -c $< -o $@

cleanall:
	@cd $(libraries) ; make clean ; cd -
	@make clean

clean:
	-@$(RM) -f $(dest) $(product) objs/* 2> /dev/null
