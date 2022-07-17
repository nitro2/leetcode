#!/bin/bash

# (xxx) xxx-xxxx 
# xxx-xxx-xxxx
grep -E -e '^\([0-9]{3}\)\s[0-9]{3}\-[0-9]{4}$' -e '^[0-9]{3}\-[0-9]{3}\-[0-9]{4}$' file.txt
# grep -Ee '\(' file.txt