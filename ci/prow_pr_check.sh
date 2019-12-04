#!/bin/bash

set -o pipefail

"$(dirname "$0")/prow_setup.sh"

{
    make check
} 2>&1 | tee -a $REPORT_DIR/test_output.log
