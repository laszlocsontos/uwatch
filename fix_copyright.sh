#!/bin/sh

for file in `find . -name "*.go"`; do

  count=`grep -c "Copyright \(C\)" $file`

  if [ $count -eq 0 ]; then
    tmp=`basename $file`

    cp $file $tmp

    cat COPYRIGHT $tmp > $file

    rm -f $tmp

    echo $file
  fi

done
