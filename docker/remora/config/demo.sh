#!/bin/sh

metric_name="demo.test_metric.number"
# metric number to post
metric=6
date=`date +%s`

echo "${metric_name}\t${metric}\t${date}"
