SMALL := Anthracite_small_Yellow.layout AnthraciteC_small.layout AnthraciteC_small_Yellow.layout

all: $(SMALL)

$(SMALL): Anthracite_small.layout
	cp $< $@