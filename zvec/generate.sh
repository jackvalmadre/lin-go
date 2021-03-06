#!/bin/bash

dir=../vec
gofile=.gofiles

ls $dir | grep '\.go$' | grep -v '_real[_.]' >$gofile
for f in `cat $gofile`
do
	echo $dir/$f -\> ./$f
	cp $dir/$f ./
	go fmt $f >/dev/null
	sed -i 's/float64/complex128/g' $f
	sed -i 's/^package vec$/package zvec/g' $f
	sed -i 's/"math"/"math\/cmplx"/g' $f
	sed -i 's/math\./cmplx\./g' $f
	go fmt $f >/dev/null
	chmod a-w $f
done
