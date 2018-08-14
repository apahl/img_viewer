#!/bin/bash
# 45575/170207-CellPainting-176cpds-10-20_A01_s5_w5B56A82A6-670D-4E30-8602-CB322F996ADB.tif
set -e

SRC_DIR=${1%/}
DEST_DIR=${2%/}

if [ -z "$SRC_DIR" ] || [ ! -d "$SRC_DIR" ]; then
  echo Source Dir $SRC_DIR does not exist
  exit 1
fi


if [ -z $DEST_DIR ] || [ ! -d $DEST_DIR ]; then
  echo Destination Dir $DEST_DIR does not exist
  exit 1
fi

mkdir -p $DEST_DIR/images
echo "Sampling images..."

for row in {A..P}; do
  for col in {01..24}; do
    # sample all images, incl. all controls
    # if [ $col -ne 11 -a $col -ne 12 ] || [ $row$col == H11 ]; then
    for w in {1..5}; do
      for f in "$SRC_DIR"/*/*/*$row${col}_s5_w$w*.tif; do
        if [ -e "$f" ] && [[ "$f" != *_thumb* ]]; then
          echo converting "$SRC_DIR" : $row${col} channel $w ...
          set +e
          convert "$f" -resize 750x750 \
                  -brightness-contrast 60x70 \
                  $DEST_DIR/images/$row${col}_w$w.jpg > /dev/null 2>&1
          set -e
        fi
      done
    done
    # fi
  done
done

