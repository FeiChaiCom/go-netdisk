#!/bin/bash

set -x

chmod a+rx bin/server

bin/server "$@"
